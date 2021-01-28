// add by stefan

package RedisConn

import (
	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/ado"
	"github.com/Peakchen/xgameCommon/ado/dbStatistics"
	"github.com/Peakchen/xgameCommon/public"
	"fmt"
	"strconv"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/gomodule/redigo/redis"
)

type TAokoRedis struct {
	ConnAddr  string
	DBIndex   int32
	Passwd    string
	RedPool   *redis.Pool
	us        *TRedisScript
	upcachecb public.UpdateDBCacheCallBack
}

func NewRedisConn(ConnAddr string, DBIndex int32, Passwd string, upcb public.UpdateDBCacheCallBack) *TAokoRedis {
	Rs := &TAokoRedis{
		ConnAddr:  ConnAddr,
		DBIndex:   DBIndex,
		Passwd:    Passwd,
		upcachecb: upcb,
	}

	Rs.us = &TRedisScript{
		name:   ERedScript_Update,
		script: updateScript(),
	}
	Rs.NewDial()
	return Rs
}

func (this *TAokoRedis) NewDial() error {
	this.RedPool = &redis.Pool{
		MaxIdle:     IDle_three,
		IdleTimeout: IDleTimeOut_four_min,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp",
				this.ConnAddr,
				redis.DialDatabase(int(this.DBIndex)),
				redis.DialPassword(this.Passwd),
				redis.DialReadTimeout(1*time.Second),
				redis.DialWriteTimeout(1*time.Second))
		},
	}
	this.RedPool.Get().Do("FLUSHDB")
	return nil
}

func (this *TAokoRedis) Exit() {
	this.RedPool.Close()
}

/*
	Redis Oper func: Insert
	SaveType: EDBOper_Insert
	purpose: in order to Insert data type EDBOperType to Redis Cache.
*/
func (this *TAokoRedis) Insert(Identify string, Input public.IDBCache) (err error) {
	RedisKey := MakeRedisModel(Identify, Input.MainModel(), Input.SubModel())
	BMarlData, err := bson.Marshal(Input)
	if err != nil {
		err = fmt.Errorf("bson.Marshal err: %v.\n", err)
		akLog.Error("[Update] err: %v", err)
		return
	}

	err = this.Save(Identify, RedisKey, BMarlData, ado.EDBOper_Insert)
	return //this.Update(Identify, Input, ado.EDBOper_Insert)
}

/*
	Redis Oper func: Update
	SaveType: EDBOper_Update
	purpose: in order to Update data type EDBOperType to Redis Cache.
*/
func (this *TAokoRedis) Update(Identify string, Input public.IDBCache, SaveType ado.EDBOperType) (err error, cacheOper bool) {
	RedisKey := MakeRedisModel(Identify, Input.MainModel(), Input.SubModel())
	BMarlData, err := bson.Marshal(Input)
	if err != nil {
		err = fmt.Errorf("bson.Marshal err: %v.\n", err)
		akLog.Error("%v", err)
		return
	}

	err, cacheOper = this.SaveEx(Identify, RedisKey, BMarlData, SaveType)
	return
}

/*
	Redis Oper func: Query
	purpose: in order to Get data from Redis Cache.
*/
func (this *TAokoRedis) Query(Identify string, Output public.IDBCache) (ret error) {
	ret = nil
	RedisKey := MakeRedisModel(Identify, Output.MainModel(), Output.SubModel())
	data, err := this.RedPool.Get().Do("GET", RedisKey)
	if err != nil {
		ret = fmt.Errorf("Identify: %v, MainModel: %v, SubModel: %v, data: %v.\n", Identify, Output.MainModel(), Output.SubModel(), data)
		akLog.Error("[Query] err: %v.\n", ret)
		return
	}

	if data == nil {
		ret = fmt.Errorf("Identify: %v, MainModel: %v, SubModel: %v, Nil data is invalid.\n", Identify, Output.MainModel(), Output.SubModel())
		akLog.Error("[Query] err: %v.\n", ret)
		return
	}

	BUmalErr := bson.Unmarshal(data.([]byte), Output)
	if BUmalErr != nil {
		ret = fmt.Errorf("Identify: %v, MainModel: %v, SubModel: %v, data: %v.\n", Identify, Output.MainModel(), Output.SubModel(), data)
		akLog.Error("[Query] can not bson Unmarshal get data to Output, err: %v.\n", ret)
		return
	}

	return
}

func (this *TAokoRedis) Save(rolekey, RedisKey string, data interface{}, SaveType ado.EDBOperType) (ret error) {
	ret = nil
	switch SaveType {
	case ado.EDBOper_Insert:
		ExpendCmd := []interface{}{RedisKey, data}
		Ret, err := this.RedPool.Get().Do("SETNX", ExpendCmd...) // set if not exist
		if err != nil {
			akLog.Error("[Save] SETNX data: %v, err: %v.\n", data, err)
			return
		}

		if Ret == 0 {
			// connect key and value.
			if _, err := this.RedPool.Get().Do("SET", ExpendCmd...); err != nil {
				akLog.Error("[Save] Insert SET data: %v, err: %v..\n", data, err)
				return
			}
		}

	case ado.EDBOper_Update:
		// connect key and value.
		var ExpendCmd = []interface{}{RedisKey, data, "EX", REDIS_SET_DEADLINE}
		if _, err := this.RedPool.Get().Do("SETEX", ExpendCmd...); err != nil {
			akLog.Error("[Save] Update Set data: %v, err: %v.\n", data, err)
			return
		}

		CollectKey := ":" + RedisKey + "_Update_Oper"
		// Add to collection.
		if _, err := this.RedPool.Get().Do("SADD", CollectKey, RedisKey); err != nil {
			akLog.Error("[Save] SADD CollectKey: %v, RedisKey: %v, err: %v.", CollectKey, RedisKey, err)
			return
		}

	case ado.EDBOper_Delete:
		// nothing...
	case ado.EDBOper_DB: //it can be presisted to database.
		// for mogo db.
	default:
		// nothing...

	}

	return
}

func (this *TAokoRedis) SaveEx(rolekey, RedisKey string, data interface{}, SaveType ado.EDBOperType) (ret error, cacheOper bool) {
	var (
		extime int32
		bsetEx bool
	)
	if SaveType == ado.EDBOper_EXPIRE {
		extime = ado.EDB_DATA_SAVE_INTERVAL
		bsetEx = true
	}

	// if can cache data, then not save redis.
	if this.upcachecb != nil {
		if this.upcachecb(rolekey, RedisKey, data.([]byte)) {
			cacheOper = true
			return
		}
	}

	dbStatistics.DBOperStatistics(rolekey, RedisKey)
	ret = this.redSetAct(rolekey, RedisKey, data, bsetEx, extime)
	return
}

func (this *TAokoRedis) redSetAct(key string, fieldkey string, data interface{}, bsetEx bool, extime int32) (err error) {
	nhashk := RoleKey2Haskey(key)
	strkey := ERedScript_Update + strconv.Itoa(nhashk)
	akLog.FmtPrintf("redis act, hashKey: %v, fieldkey: %v.", strkey, fieldkey)
	c := this.RedPool.Get()
	if c == nil {
		err = akLog.RetError("red pool get session fail.")
		return
	}

	var (
		exCmd = []interface{}{}
	)

	if bsetEx {
		exCmd = []interface{}{strkey, fieldkey, 1, data, "EX", extime}
	} else {
		exCmd = []interface{}{strkey, fieldkey, 1, data, "", 0}
	}

	_, err = this.us.script.Do(c, exCmd...)
	if err != nil {
		err = akLog.RetError("name: %v, ex cmd %v, err: %v", this.us.name, exCmd, err)
		return
	}

	akLog.FmtPrintf("redis update succ, hashKey: %v.", strkey)
	err = nil
	return
}
