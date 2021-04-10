package service

// add by stefan

import (
	"github.com/Peakchen/xgameCommon/ado"
	"github.com/Peakchen/xgameCommon/public"
	"github.com/gomodule/redigo/redis"
)

/*
	@func: Update
	@param1: Identify string     player only one
	@param2: data public.IDBCache  module need save
	@param3: Oper ado.EDBOperType  db operation
	purpose: first db save data then update cache.
*/
func (this *TDBProvider) Update(Identify string, data public.IDBCache, Oper ado.EDBOperType) (err error) {
	var cacheOper bool
	err, cacheOper = this.rconn.Update(Identify, data, Oper)
	if err != nil || cacheOper {
		return
	}

	err = this.mconn.SaveOne(Identify, data)
	if err != nil {
		return
	}

	return
}

/*
	@func: Insert
	@param1: Identify string     player only one
	@param2: data public.IDBCache  module need save
	purpose: first insert data to cache then update db.
*/
func (this *TDBProvider) Insert(Identify string, data public.IDBCache) (err error) {
	err = this.rconn.Insert(Identify, data)
	if err == nil {
		err = this.mconn.InsertOne(Identify, data)
	}
	return
}

/*
	@func: Get
	@param1: Identify string     player only one
	@param2: Output public.IDBCache  module need query
	purpose: first query from cache if not exist, then find from db.
*/
func (this *TDBProvider) Get(Identify string, Output public.IDBCache) (err error, exist bool) {
	err = this.rconn.Query(Identify, Output)
	if err != nil {
		err, exist = this.mconn.QueryOne(Identify, Output)
		//redis not exist, then update
		err, _ = this.rconn.Update(Identify, Output, ado.EDBOper_Update)
	} else {
		exist = true
	}
	return
}

/*
	@func: GetAcc
	@param1: Identify string     player only one
	@param2: Output public.IDBCache  module need query
	purpose: first query use accout from db.
*/
func (this *TDBProvider) GetAcc(usrName string, Output public.IDBCache) (err error, exist bool) {
	err, exist = this.mconn.QueryAcc(usrName, Output)
	return
}

/*
	@func: DBGetSome
	@param1: Identify string     player only one
	@param2: Output public.IDBCache  module need query
	purpose: first query from cache if not then db.
*/
func (this *TDBProvider) DBGetSome(Output public.IDBCache) (err error) {
	// has no func need.
	return nil
}

func (this *TDBProvider) GetRedisConn() redis.Conn {
	return this.rconn.RedPool.Get()
}
