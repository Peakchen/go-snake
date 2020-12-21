package account

import (
	"go-snake/loginServer/entityMgr"
	"go-snake/loginServer/logic/account/acc_model"
)

func init() {
	entityMgr.RegisterModel(entityMgr.M_ACCOUNT, func(entity entityMgr.IEntityUser) interface{} { return newAcc(entity) })
}

type Acc struct {
	entityMgr.IEntityUser

	user *acc_model.Acc
}

func newAcc(entity entityMgr.IEntityUser) *Acc {
	return &Acc{
		IEntityUser: entity,
		user:        nil,
	}
}

func GetAcc(user, pwd string) *Acc {
	return nil
}

func (this *Acc) setAcc(acc *acc_model.Acc) {
	this.user = acc
}
