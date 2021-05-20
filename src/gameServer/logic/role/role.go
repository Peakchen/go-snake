package role

import (
	"go-snake/core/user"
	"go-snake/dbmodel/role_model"
)

func init() {
	user.RegisterModel(user.M_ROLE, func(entity user.IEntityUser) interface{} { return newRoleCache(entity) })
}

type RoleCache struct {
	user.IEntityUser

	role *role_model.Role
}

func newRoleCache(entity user.IEntityUser) *RoleCache {
	return &RoleCache{
		IEntityUser: entity,
	}
}
