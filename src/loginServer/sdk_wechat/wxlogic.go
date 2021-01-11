package sdk_wechat

import (
	"go-snake/loginServer/entityBase"
	"go-snake/loginServer/sdk_wechat/wechat_model"
)

func init() {
	entityBase.RegisterModel(entityBase.M_WXROLE, func(entity entityBase.IEntityUser) interface{} { return newWxRole(entity) })
}

type WxRole struct {
	entityBase.IEntityUser

	role *wechat_model.WxRole
}

func newWxRole(entity entityBase.IEntityUser) *WxRole {
	return &WxRole{
		IEntityUser: entity,
		role:        nil,
	}
}

func (this *WxRole) LoadWxRole(role *wechat_model.WxRole) {
	this.role = role
}
