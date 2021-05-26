package logic

/*
	load db for cache.
*/

import (
	"go-snake/core/usermgr"
	"go-snake/core/modelcache"
)

func LoadDB(){

	usermgr.DBPreAddAndLoad(
		modelcache.LoadAllRole,
		modelcache.LoadAllWxRole)

}