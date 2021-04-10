package service

import (
	"context"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/Peakchen/xgameCommon/MgoConn"
	"github.com/Peakchen/xgameCommon/RedisConn"
	"github.com/Peakchen/xgameCommon/ado"
	"github.com/Peakchen/xgameCommon/akLog"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gomodule/redigo/redis"
)

/*
	db module: a lot of redis sessions
	purpose: 建立指定数量的redis 链接，不同玩家唯一认证与之关联，定时快速写入mgo，保证数据文档安全.
*/

type TClusterDBProvider struct {
	//redConn []*RedisConn.TAokoRedis
	redConn     *RedisConn.TAokoRedis
	mgoConn     *MgoConn.TAokoMgo
	mgoSessions []*mgo.Session
	Server      string
	ctx         context.Context
	cancle      context.CancelFunc
	wg          sync.WaitGroup
}

func (this *TClusterDBProvider) init(Server string) {
	this.Server = Server
	this.mgoSessions = make([]*mgo.Session, ado.EMgo_Thread_Cnt)
}

func (this *TClusterDBProvider) Start(Server string, rediscfg *ado.TRedisConfig, mgocfg *ado.TMgoConfig) {
	this.init(Server)
	this.runDBloop(Server, rediscfg, mgocfg)
}

func (this *TClusterDBProvider) GetRedisConn() redis.Conn {
	return this.redConn.RedPool.Get()
}

func (this *TClusterDBProvider) Exit() {
	if this.mgoConn != nil {
		this.mgoConn.Exit()
	}

	if this.redConn != nil {
		this.redConn.Exit()
	}
}

func (this *TClusterDBProvider) runDBloop(Server string, rediscfg *ado.TRedisConfig, mgocfg *ado.TMgoConfig) {
	this.redConn = RedisConn.NewRedisConn(rediscfg.Connaddr, rediscfg.DBIndex, rediscfg.Passwd, nil)
	this.mgoConn = MgoConn.NewMgoConn(Server, mgocfg.Username, mgocfg.Passwd, mgocfg.Host)

	this.ctx, this.cancle = context.WithCancel(context.Background())
	session, err := this.mgoConn.GetMgoSession()
	if err != nil {
		akLog.Error(err)
		return
	}

	for midx := int32(0); midx < ado.EMgo_Thread_Cnt; midx++ {
		this.mgoSessions[midx] = session.Copy()
	}

	this.wg.Add(2)
	go this.LoopDBUpdate(&this.wg)
	defer func() {
		http.ListenAndServe(rediscfg.Pprofaddr, nil)
	}()
	this.wg.Wait()
}

func (this *TClusterDBProvider) LoopDBUpdate(wg *sync.WaitGroup) {
	defer func() {
		this.Exit()
		wg.Done()
	}()

	ticker := time.NewTicker(time.Duration(ado.EDB_DATA_SAVE_INTERVAL) * time.Second)
	for {
		select {
		case <-this.ctx.Done():
			ticker.Stop()
			return
		case <-ticker.C:
			// do something...
			this.flushdb()
		}
	}
}

func (this *TClusterDBProvider) flushdb() {

	c := this.redConn.RedPool.Get()
	if c == nil {
		akLog.Error("redis conn invalid or disconntion.")
		return
	}

	var (
		ridx int32
	)
	for {
		if ridx == ado.EMgo_Thread_Cnt {
			break
		}

		this.dbupdate(ridx, c)
		ridx++
	}
}

func (this *TClusterDBProvider) dbupdate(ridx int32, c redis.Conn) {
	//akLog.FmtPrintln("db update idx: ", ridx)
	// TODO: Presist redis...
	if this.redConn == nil {
		akLog.Error("redis conn invalid or conn number invalid, info: ", this.redConn, ridx)
		return
	}

	updateidx := strconv.Itoa(int(ridx))
	onekey := RedisConn.ERedScript_Update + updateidx
	members, err := c.Do("HKEYS", onekey)
	if err != nil || members == nil {
		akLog.Error("ClusterDBProvider get redis,err: ", err)
		return
	}

	// TODO: Presist mgo...
	mgosession := this.mgoSessions[ridx]
	if mgosession == nil {
		akLog.Error("mgoConn invalid or disconntion.")
		return
	}

	for _, item := range members.([]interface{}) {
		dstkey := string(item.([]byte))
		dstval, err := c.Do("GET", dstkey)
		if err != nil {
			akLog.Error("get fail, err: ", err)
			continue
		}

		bsdata := bson.Raw{Kind: byte(0), Data: dstval.([]byte)}
		err = MgoConn.Save(mgosession, this.Server, dstkey, bsdata)
		if err != nil {
			akLog.Error("mgo update err: ", err)
		}
	}
}
