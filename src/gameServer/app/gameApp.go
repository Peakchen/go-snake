package app

import "ak-remote/common/messageBase"

//gate2 <-> game server
type IGameApp interface {
	Online(nt messageBase.NetType, sess interface{})
	Offline(nt messageBase.NetType, id string)
}

var _app IGameApp

func SetApp(g IGameApp) { _app = g }
func GetApp() IGameApp  { return _app }
