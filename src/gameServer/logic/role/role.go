package role

import (
	"go-snake/gameServer/entityBase"
	"go-snake/gameServer/logic/role/role_model"
)

func init() {
	entityBase.RegisterModel(entityBase.M_ROLE, func(entity entityBase.IEntityUser) interface{} { return newRoleCache(entity) })
}

type RoleCache struct {
	entityBase.IEntityUser

	role *role_model.Role
}

func newRoleCache(entity entityBase.IEntityUser) *RoleCache {
	return &RoleCache{
		IEntityUser: entity,
	}
}
