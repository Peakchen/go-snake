package allocation

import (
	"errors"
	"fmt"

	"github.com/Peakchen/xgameCommon/RedisConn"
	"github.com/Peakchen/xgameCommon/ado/service"
)

var (
	allocIntMap = map[AllocInteger]string{
		AllocInt_8:  "collectionInt8",
		AllocInt_16: "collectionInt16",
		AllocInt_32: "collectionInt32",
		AllocInt_64: "collectionInt64",
	}
)

func AllocateInteger(dbprovider *service.TDBProvider, at AllocInteger, out interface{}) (err error) {
	sAlt, exist := allocIntMap[at]
	if !exist {
		err = fmt.Errorf("err: %v, at: %v.", errors.New("can not find allocIntMap data."), at)
		return
	}
	dbprovider.GetAkRedis().GetDecodeCache(sAlt, out)
	return nil
}

func AllocateInteger2(conn *RedisConn.TAokoRedis, at AllocInteger, out interface{}) (err error) {
	sAlt, exist := allocIntMap[at]
	if !exist {
		err = fmt.Errorf("err: %v, at: %v.", errors.New("can not find allocIntMap data."), at)
		return
	}
	ret := conn.GetDecodeCache(sAlt, out)
	if ret == RedisErr_Empty {
		out = 1
		conn.SetEncodeCache(sAlt, out)
	}
	return nil
}
