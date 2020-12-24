package inner

import (
	"go-snake/gameServer/entityMgr"
)

func init() {
	entityMgr.RegisterModel(entityMgr.M_SERVERINNER, func(entity entityMgr.IEntityUser) interface{} { return newInner(entity) })
}

type ServerInner struct {
	entityMgr.IEntityUser
}

func newInner(entity entityMgr.IEntityUser) *ServerInner {
	return &ServerInner{IEntityUser: entity}
}
