package inner

import (
	"go-snake/core/user"
)

func init() {
	user.RegisterModel(user.M_SERVERINNER, func(entity user.IEntityUser) interface{} { return newInner(entity) })
}

type ServerInner struct {
	user.IEntityUser
}

func newInner(entity user.IEntityUser) *ServerInner {
	return &ServerInner{IEntityUser: entity}
}
