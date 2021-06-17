package friend

import (
	"go-snake/core/user"
)

func init() {
	user.RegisterModel(user.M_Friend, func(entity user.IEntityUser) interface{} { return newFriend(entity) })
}

func newFriend(entity user.IEntityUser)*Friend{
	return &Friend{
		IEntityUser: entity,
	}
}

type Friend struct {

	user.IEntityUser

}



