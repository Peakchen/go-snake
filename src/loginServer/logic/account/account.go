package account

import (
	"go-snake/core/user"
	"go-snake/dbmodel/acc_model"
)

func init() {
	user.RegisterModel(user.M_ACCOUNT, func(entity user.IEntityUser) interface{} { return newAcc(entity) })
}

type Acc struct {
	user.IEntityUser

	user *accdb.Acc
}

func newAcc(entity user.IEntityUser) *Acc {
	return &Acc{
		IEntityUser: entity,
		user:        nil,
	}
}

func (this *Acc) GetAcc(user, pwd string) *accdb.Acc {
	return this.user
}

func (this *Acc) LoadAcc(acc *accdb.Acc) {
	this.user = acc
}
