package mail

import (
	"go-snake/core/user"
)

func init() {
	user.RegisterModel(user.M_Mail, func(entity user.IEntityUser) interface{} { return newMail(entity) })
}

func newMail(entity user.IEntityUser)*Mail{
	return &Mail{
		IEntityUser: entity,
	}
}

type Mail struct {

	user.IEntityUser

}



