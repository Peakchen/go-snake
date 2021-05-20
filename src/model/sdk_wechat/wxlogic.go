package sdk_wechat

import (
	"go-snake/core/user"
	"go-snake/dbmodel/wechat_model"
)

func init() {
	user.RegisterModel(user.M_WXROLE, func(entity user.IEntityUser) interface{} { return newWxRole(entity) })
}

type WxRole struct {
	user.IEntityUser

	role *wechat_model.WxRole
}

func newWxRole(entity user.IEntityUser) *WxRole {
	return &WxRole{
		IEntityUser: entity,
		role:        nil,
	}
}

func (this *WxRole) LoadWxRole(role *wechat_model.WxRole) {
	this.role = role
}
