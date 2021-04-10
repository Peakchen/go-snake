package rpcBase

import (
	"reflect"
	"google.golang.org/protobuf/proto"
	"go-snake/akmessage"
	"errors"
	"fmt"
)

const (
	RPC_LOGIN 		= "login"
	RPC_GAME 		= "game"
	RPC_GATE 		= "gate"
	RPC_CHAT 		= "chat"
	RPC_EMAIL   	= "email"
	RPC_SIMULATION 	= "simulation"
)

var (
	allnodeMap = []string{RPC_LOGIN,RPC_GAME,RPC_GATE,RPC_CHAT,RPC_EMAIL,RPC_SIMULATION}
)

func GetAllNode()[]string{
	return allnodeMap
}

type (
	RpcMessageFunc func(proto.Message) (*akmessage.RpcResponse, error)
	
	RpcMessage struct {
		RefFn reflect.Value //	RpcMessageFunc
		RefPb reflect.Type 	// 	proto.Message
	}

	
)

func MakeRpcResponse(msgRef reflect.Value)(*akmessage.RpcResponse, error){

	data, err := proto.Marshal(msgRef.Elem().Interface().(proto.Message))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("proto marshal fail, msg: %v.", msgRef.Elem().String()))
	}

	return &akmessage.RpcResponse{RespData: data}, nil

}