package chat

import (
	"go-snake/core/user"
)

func init() {
	user.RegisterModel(user.M_Chat, func(entity user.IEntityUser) interface{} { return newChat(entity) })
}

func newChat(entity user.IEntityUser)*Chat{
	return &Chat{
		IEntityUser: entity,
	}
}

type Chat struct {

	user.IEntityUser

}



