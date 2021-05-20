package usermgr

import (
	"go-snake/core/user"
	"go-snake/dbmodel/acc_model"
	"go-snake/dbmodel/wechat_model"

	"github.com/Peakchen/xgameCommon/akLog"
)

func (this *EntityManager) LoadAll() {
	this.LoadAllRole()
	this.LoadAllWxRole()
}

func (this *EntityManager) LoadAllRole() {
	accs := (&accdb.Acc{}).Load()
	if len(accs) > 0 {
		for _, model := range accs {
			entity := user.InitEntity(model.DBID)
			if !this.AddEnity(model.DBID, entity) {
				akLog.Fail("exist same entity, dbid: ", model.DBID)
			}
			entity.LoadAcc(model)
		}
	}
}

func (this *EntityManager) LoadAllWxRole() {
	wxroles := (&wechat_model.WxRole{}).Load()
	if len(wxroles) > 0 {
		for _, model := range wxroles {
			entity := user.InitEntity(model.DBID)
			if !this.AddEnity(model.DBID, entity) {
				akLog.Fail("exist same entity, dbid: ", model.DBID)
			}
			entity.LoadWxRole(model)
		}
	}
}
