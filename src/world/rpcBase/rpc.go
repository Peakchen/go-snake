package rpcBase

import (
	"reflect"
	"go-snake/akmessage"
	"go-snake/common/myetcd"
	"go-snake/common/rpcBase"
	"context"

	"github.com/Peakchen/xgameCommon/akLog"

)

type WorldRpc struct {
	*myetcd.RpcMessageNode

}

func newWorldRpc()*WorldRpc{
	return &WorldRpc{
		&myetcd.RpcMessageNode{
			MsgNodes: map[akmessage.RPCMSG]*rpcBase.RpcMessage{
			},
			NodeName: rpcBase.RPC_GAME,
		},
	}
}

func (this *WorldRpc) Register(id akmessage.RPCMSG, pb interface{}, fn rpcBase.RpcMessageFunc) {

	this.RpcMessageNode.MsgNodes[id] = &rpcBase.RpcMessage{
		RefPb:   reflect.TypeOf(pb),
		RefFn: 	 reflect.ValueOf(fn),
	}

}

func (this *WorldRpc) Send(ctx context.Context, in *akmessage.RpcRequest)(*akmessage.RpcResponse, error){
	
	akLog.FmtPrintln("rpc call: ", this.NodeName)

	return this.RpcMessageNode.Call(ctx, in)
}
