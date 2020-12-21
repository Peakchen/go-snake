package innner

import (
	"go-snake/loginServer/entityMgr"
)

func init() {
	entityMgr.RegisterModel(entityMgr.M_SERVERINNER, func(entity entityMgr.IEntityUser) interface{} { return newClientReg(entity) })
}

type ServerInner struct {
	entityMgr.IEntityUser
}

func newClientReg(entity entityMgr.IEntityUser) *ServerInner {
	return &ServerInner{IEntityUser: entity}
}
