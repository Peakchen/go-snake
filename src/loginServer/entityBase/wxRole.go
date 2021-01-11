package entityBase

import (
	"go-snake/loginServer/sdk_wechat/wechat_model"
)

type IWxRole interface {
	LoadWxRole(role *wechat_model.WxRole)
}
