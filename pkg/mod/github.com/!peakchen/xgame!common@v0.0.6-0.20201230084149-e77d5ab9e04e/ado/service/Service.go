package service

import (
	"context"
	"sync"

	"github.com/Peakchen/xgameCommon/MgoConn"
	"github.com/Peakchen/xgameCommon/RedisConn"
	"github.com/Peakchen/xgameCommon/ado"
	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/public"
)

type TDBProvider struct {
	rconn  *RedisConn.TAokoRedis
	mconn  *MgoConn.TAokoMgo
	Server string
	ctx    context.Context
	cancle context.CancelFunc
	wg     sync.WaitGroup
}

func (this *TDBProvider) StartDBService(Server string, upcb public.UpdateDBCacheCallBack, rediscfg *ado.TRedisConfig, mgocfg *ado.TMgoConfig) {
	this.Server = Server
	this.rconn = RedisConn.NewRedisConn(rediscfg.Connaddr, rediscfg.DBIndex, rediscfg.Passwd, upcb)
	this.mconn = MgoConn.NewMgoConn(Server, mgocfg.Username, mgocfg.Passwd, mgocfg.Host)
}

func (this *TDBProvider) GetAkRedis() *RedisConn.TAokoRedis {
	return this.rconn
}

func (this *TDBProvider) RediSave(identify string, rediskey string, data []byte, Oper ado.EDBOperType) (err error) {
	err, _ = this.rconn.SaveEx(identify, rediskey, data, Oper)
	if err != nil {
		akLog.ErrorIDCard(identify, "update redis fail, rediskey: ", rediskey, ", err: ", err)
	}
	return
}

func (this *TDBProvider) GetMogoConn() *MgoConn.TAokoMgo {
	return this.mconn
}
