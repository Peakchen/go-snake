package rpcBase

import (
	"reflect"
	"go-snake/akmessage"
	"go-snake/common/myetcd"
	"go-snake/common/logicBase"
	"context"

	"github.com/Peakchen/xgameCommon/akLog"

)

type LoginRpc struct {
	*myetcd.RpcMessageNode

}

func newLoginRpc()*LoginRpc{
	return &LoginRpc{
		&myetcd.RpcMessageNode{
			MsgNodes: map[akmessage.RPCMSG]*logicBase.RpcMessage{
			},
			NodeName: logicBase.RPC_LOGIN,
		},
	}
}

func (this *LoginRpc) Register(id akmessage.RPCMSG, pb interface{}, fn logicBase.RpcMessageFunc) {

	this.RpcMessageNode.MsgNodes[id] = &logicBase.RpcMessage{
		RefPb:   reflect.TypeOf(pb),
		RefFn: 	 reflect.ValueOf(fn),
	}

}

func (this *LoginRpc) Send(ctx context.Context, in *akmessage.RpcRequest)(*akmessage.RpcResponse, error){
	akLog.FmtPrintln("rpc call: ", this.RpcMessageNode.NodeName)
	return this.RpcMessageNode.Call(ctx, in)
}
