package rpcBase

import (
	"reflect"
	"go-snake/akmessage"
	"go-snake/common/myetcd"
	"go-snake/common/rpcBase"
	"context"

	"github.com/Peakchen/xgameCommon/akLog"

)

type BattleRpc struct {
	*myetcd.RpcMessageNode

}

func newBattleRpc()*BattleRpc{
	return &BattleRpc{
		&myetcd.RpcMessageNode{
			MsgNodes: map[akmessage.RPCMSG]*rpcBase.RpcMessage{
			},
			NodeName: rpcBase.RPC_GAME,
		},
	}
}

func (this *BattleRpc) Register(id akmessage.RPCMSG, pb interface{}, fn rpcBase.RpcMessageFunc) {

	this.RpcMessageNode.MsgNodes[id] = &rpcBase.RpcMessage{
		RefPb:   reflect.TypeOf(pb),
		RefFn: 	 reflect.ValueOf(fn),
	}

}

func (this *BattleRpc) Send(ctx context.Context, in *akmessage.RpcRequest)(*akmessage.RpcResponse, error){
	
	akLog.FmtPrintln("rpc call: ", this.NodeName)

	return this.RpcMessageNode.Call(ctx, in)
}
