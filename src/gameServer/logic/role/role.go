package role

import (
	"go-snake/gameServer/entityMgr"
	"go-snake/gameServer/logic/role/role_model"
)

func init() {
	entityMgr.RegisterModel(entityMgr.M_ROLE, func(entity entityMgr.IEntityUser) interface{} { return newRoleCache(entity) })
}

type RoleCache struct {
	entityMgr.IEntityUser

	role *role_model.Role
}

func newRoleCache(entity entityMgr.IEntityUser) *RoleCache {
	return &RoleCache{
		IEntityUser: entity,
	}
}
