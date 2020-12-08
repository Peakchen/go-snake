package cache

import (
	//"reflect"
	"ak-remote/common/rediscache"
	"fmt"

	"github.com/globalsign/mgo/bson"
)

/*
	--------------------加密数据部分缓存----------------------
*/
func GetDecodeCache(identify string, v interface{}) (err error, succ bool) {
	data := GMemCache.Get(identify)
	if data == nil {
		err, succ = rediscache.GetDecodeCache(identify, v.(rediscache.IDBCache))
		return
	}

	err = bson.Unmarshal(data.([]byte), v)
	succ = true
	return
}

func SetEncodeCache(identify string, v interface{}) (err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		err = fmt.Errorf("bson marshal err, key: ", identify)
		return
	}

	GMemCache.Set(identify, data)
	err = nil
	return
}

/*
	---------------------非数据加密部分-----------------
*/
func GetCache(identify string) (v interface{}, err error) {
	v = GMemCache.Get(identify)
	if v == nil {
		v, err = rediscache.GetCache(identify)
	} else {
		err = nil
	}
	return
}

func SetCache(identify string, v interface{}) {
	GMemCache.Set(identify, v)
}
