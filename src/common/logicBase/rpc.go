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

type (
	FunCallBack func(arg interface{})
	
	RpcMessage struct {
		refFn reflect.Value //	FunCallBack
		refPb reflect.Type 	// 	proto.Message
	}
)