package akmessage

func GetMsgRoute(msgid MSG) ServerType {
	if msgid < MSG_LOGIN_MAX {
		return ServerType_Login
	} else if MSG_LOGIN_MAX < msgid && msgid < MSG_SS_MAX {
		return ServerType_Gate
	} else if MSG_GAME_MAX > msgid && msgid > MSG_SS_MAX {
		return ServerType_Game
	} else {
		return ServerType_No
	}
}
