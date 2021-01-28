package myetcd

import (
	"reflect"
	"context"
	"go-snake/akmessage"
	"go-snake/common/logicBase"

	"google.golang.org/protobuf/proto"

)

type (
	RpcMessageNode struct {
		name string
		msgNodes map[akmessage.RPCMSG]*logicBase.RpcMessage
	}
)

func (this *RpcMessageNode) Name() string{
	return this.name
}

func (this *RpcMessageNode) Call(ctx context.Context, in *akmessage.RpcRequest)(*akmessage.RpcResponse, error){
	node, ok := this.msgNodes[in.MsgId]
	if !ok || fn == nil{
		return nil, errors.New("can not find msg node")
	}
	dst := reflect.New(node.refPb.Elem()).Interface().(proto.Message)
	err := proto.Unmarshal(in.ReqData, refPb)
	if err != nil {
		return nil, errors.New("proto unmarshal fail")
	}
	rets := node.refFn.Call([]reflect.Value{
		reflect.ValueOf(refPb),
	})
	rsp := rets.Interface().(*akmessage.RpcResponse)
	err = rets.Interface().(error)
	return rsp, err
}