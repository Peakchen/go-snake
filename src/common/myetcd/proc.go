package myetcd

import (
	"reflect"
	"context"
	"go-snake/akmessage"
	"go-snake/common/logicBase"
	"errors"

	"google.golang.org/protobuf/proto"

)

type (
	RpcMessageNode struct {
		akmessage.UnimplementedRpcServer

		NodeName string
		MsgNodes map[akmessage.RPCMSG]*logicBase.RpcMessage
	}
)

func (this *RpcMessageNode) Call(ctx context.Context, in *akmessage.RpcRequest)(*akmessage.RpcResponse, error){
	node, ok := this.MsgNodes[in.MsgId]
	if !ok || node == nil{
		return nil, errors.New("can not find msg node")
	}
	dst := reflect.New(node.RefPb.Elem()).Interface().(proto.Message)
	err := proto.Unmarshal(in.ReqData, dst)
	if err != nil {
		return nil, errors.New("proto unmarshal fail")
	}
	rets := node.RefFn.Call([]reflect.Value{
		reflect.ValueOf(ctx),
		reflect.ValueOf(dst),
	})
	rsp := rets[0].Interface().(*akmessage.RpcResponse)
	err = rets[1].Interface().(error)
	return rsp, err
}