package inner

import (
	"go-snake/loginServer/entityBase"
)

func init() {
	entityBase.RegisterModel(entityBase.M_SERVERINNER, func(entity entityBase.IEntityUser) interface{} { return newInner(entity) })
}

type ServerInner struct {
	entityBase.IEntityUser
}

func newInner(entity entityBase.IEntityUser) *ServerInner {
	return &ServerInner{IEntityUser: entity}
}
