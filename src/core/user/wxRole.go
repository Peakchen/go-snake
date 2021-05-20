package user

import (
	"go-snake/dbmodel/wechat_model"
)

type IWxRole interface {
	LoadWxRole(role *wechat_model.WxRole)
}
