package logicBase

import (
	"reflect"

)

const (
	RPC_LOGIN 	= "login"
	RPC_GAME 	= "game"
	RPC_GATE 	= "gate"
	RPC_CHAT 	= "chat"
	RPC_EMAIL   = "email"

)

var (
	allnodeMap = []string{RPC_LOGIN,RPC_GAME,RPC_GATE,RPC_CHAT,RPC_EMAIL}
)

func GetAllNode()[]string{
	return allnodeMap
}

type (
	FunCallBack func(arg interface{})
	
	RpcMessage struct {
		RefFn reflect.Value //	FunCallBack
		RefPb reflect.Type 	// 	proto.Message
	}
)