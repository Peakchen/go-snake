package app

import "ak-remote/common/messageBase"

type IApp interface {
	Online(nt messageBase.NetType, sess interface{})
	Offline(nt messageBase.NetType, id string)
}
