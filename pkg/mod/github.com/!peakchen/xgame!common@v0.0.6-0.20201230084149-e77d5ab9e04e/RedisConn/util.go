package RedisConn

import (
	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/globalsign/mgo/bson"
)

type RedisErr int

const (
	RedisErr_Invalid       RedisErr = 1 //非法数据
	RedisErr_Empty         RedisErr = 2 //空值
	RedisErr_Success       RedisErr = 3 //成功
	RedisErr_MarshalFail   RedisErr = 4 //编码失败
	RedisErr_UNMarshalFail RedisErr = 5 //解码失败
	RedisErr_DoFail        RedisErr = 6 //redis 操作异常
)

func (this *TAokoRedis) SetEncodeCache(identify string, src interface{}) (ret RedisErr) {
	data, err := bson.Marshal(src)
	if err != nil {
		akLog.Error("bson marshal fail, identify: %v, src: %v.", identify, src)
		return RedisErr_MarshalFail
	}
	ret = this.Set(identify, data)
	return
}

func (this *TAokoRedis) GetDecodeCache(identify string, out interface{}) (ret RedisErr) {
	data, gret := this.Get(identify)
	if gret != RedisErr_Success {
		return gret
	}
	if data == nil {
		return RedisErr_Empty
	}
	err := bson.Unmarshal(data.([]byte), out)
	if err != nil {
		akLog.Error("bson unmarshal fail, identify: %v.", identify)
		return RedisErr_UNMarshalFail
	}
	return RedisErr_Success
}

func (this *TAokoRedis) Get(key string) (data interface{}, ret RedisErr) {
	var err error
	data, err = this.RedPool.Get().Do("GET", key)
	if err != nil {
		akLog.Error("[Get] key: %v, err: %v.\n", key, err)
		return nil, RedisErr_DoFail
	}
	return data, RedisErr_Success
}

func (this *TAokoRedis) Set(key string, val interface{}) (ret RedisErr) {
	_, err := this.RedPool.Get().Do("SET", key, val)
	if err != nil {
		akLog.Error("[Get] key: %v, err: %v.\n", key, err)
		return RedisErr_DoFail
	}
	return RedisErr_Success
}
