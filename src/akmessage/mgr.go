package akmessage

var _msgRoute = map[MSG]ServerType{
	MSG_CS_LOGIN:            ServerType_Login,
	MSG_CS_LOGOUT:           ServerType_Login,
	MSG_CS_ENTER_GAME_SCENE: ServerType_Game,
	MSG_CS_HEARTBEAT:        ServerType_Gate,
	MSG_SS_HEARTBEAT_REQ:    ServerType_Gate,
	MSG_SS_REGISTER_REQ:     ServerType_Gate,
}

func GetMsgRoute(msgid MSG) ServerType {
	return _msgRoute[msgid]
}
