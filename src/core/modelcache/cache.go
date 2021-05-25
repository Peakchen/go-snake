package modelcache

/*
	model load cache.
*/

import (
	"go-snake/core/user"
	"go-snake/dbmodel/acc_model"
	"go-snake/dbmodel/wechat_model"
	"go-snake/core/usermgr"
	"github.com/Peakchen/xgameCommon/akLog"
)


func LoadAllRole(em *usermgr.EntityManager) {

	accs := (&accdb.Acc{}).Load()

	if len(accs) > 0 {

		for _, model := range accs {

			entity := user.InitEntity(model.DBID)

			if !em.AddEnity(model.DBID, entity) {
				akLog.Error("exist same entity, dbid: ", model.DBID)
			}

			entity.LoadAcc(model)
		}

	}
}

func LoadAllWxRole(em *usermgr.EntityManager) {

	wxroles := (&wechat_model.WxRole{}).Load()

	if len(wxroles) > 0 {

		for _, model := range wxroles {

			entity := user.InitEntity(model.DBID)
			if !em.AddEnity(model.DBID, entity) {
				akLog.Error("exist same entity, dbid: ", model.DBID)
			}

			entity.LoadWxRole(model)

		}
	}
}
