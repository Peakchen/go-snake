package account

import (
	"go-snake/loginServer/entityBase"
	"go-snake/loginServer/logic/account/acc_model"
)

func init() {
	entityBase.RegisterModel(entityBase.M_ACCOUNT, func(entity entityBase.IEntityUser) interface{} { return newAcc(entity) })
}

type Acc struct {
	entityBase.IEntityUser

	user *acc_model.Acc
}

func newAcc(entity entityBase.IEntityUser) *Acc {
	return &Acc{
		IEntityUser: entity,
		user:        nil,
	}
}

func (this *Acc) GetAcc(user, pwd string) *acc_model.Acc {
	return this.user
}

func (this *Acc) LoadAcc(acc *acc_model.Acc) {
	this.user = acc
}
