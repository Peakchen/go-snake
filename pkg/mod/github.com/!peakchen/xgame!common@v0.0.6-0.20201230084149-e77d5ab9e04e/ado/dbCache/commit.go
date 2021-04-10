package dbCache

import (
	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/ado"
	"github.com/Peakchen/xgameCommon/ado/service"
	"github.com/Peakchen/xgameCommon/public"
	"github.com/globalsign/mgo/bson"
	"sync"
)

/*
	redis or db cache with protects, prevent breakdown.
*/

type TModelOper struct {
	buff  []byte
	opers int
}

type TDBCache struct {
	users      sync.Map // key: identify, value: map[string]*TModelOper
	dbprovider *service.TDBProvider
}

var (
	_dbCache *TDBCache
)

func InitDBCache(dbProvider *service.TDBProvider) {
	_dbCache = &TDBCache{
		dbprovider: dbProvider,
	}
}

func GetDBCache() *TDBCache {
	return _dbCache
}

func (this *TDBCache) loadOrAddUser(identify string) (modeldata map[string]*TModelOper) {
	modeldata = nil
	value, _ := this.users.LoadOrStore(identify, make(map[string]*TModelOper, 0))
	if value == nil {
		akLog.Error("can not load cache model.")
		return
	}

	if value == nil {
		akLog.Error("cache model invalid.")
		return
	}

	var ok bool
	modeldata, ok = value.(map[string]*TModelOper)
	if !ok {
		akLog.Error("cache model invalid data type.")
		return
	}
	return
}

func (this *TDBCache) push(identify string, model string) {
	modeldata := this.loadOrAddUser(identify)
	if modeldata == nil {
		return
	}

	Moper, ok := modeldata[model]
	if !ok {
		modeldata[model] = &TModelOper{
			opers: 1,
		}
	} else {
		Moper.opers++
	}
}

func (this *TDBCache) hasExist(identify string, model string) (exist bool) {
	modeldata := this.loadOrAddUser(identify)
	if modeldata == nil {
		return
	}

	_, ok := modeldata[model]
	if !ok {
		return
	}

	exist = true
	return
}

func (this *TDBCache) getCache(identify string, model string, Output public.IDBCache) {
	modeldata := this.loadOrAddUser(identify)
	if modeldata == nil {
		return
	}

	m, ok := modeldata[model]
	if !ok {
		return
	}

	err := bson.Unmarshal(m.buff, Output)
	if err != nil {
		return
	}
}

func (this *TDBCache) updateCache(identify string, model string, data []byte) (succ bool) {
	modeldata := this.loadOrAddUser(identify)
	if modeldata == nil {
		return
	}

	m, ok := modeldata[model]
	if !ok {
		return
	}

	m.buff = data
	m.opers++
	succ = true
	return
}

func (this *TDBCache) pop(identify string) {
	modeldata := this.loadOrAddUser(identify)
	if modeldata == nil {
		return
	}

	modeldata = map[string]*TModelOper{}
}

func (this *TDBCache) updateDB(identify string) {
	modeldata := this.loadOrAddUser(identify)
	if modeldata == nil {
		return
	}

	if len(modeldata) == 0 {
		return
	}

	for smodel, Operdata := range modeldata {
		RedisKey := smodel + "." + identify
		err := this.dbprovider.RediSave(identify, RedisKey, Operdata.buff, ado.EDBOper_Update)
		if err != nil {
			akLog.ErrorIDCard(identify, "update redis fail, model: ", smodel, ", err: ", err)
		}
	}
}
