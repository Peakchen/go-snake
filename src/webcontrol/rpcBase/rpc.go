package rpcBase

import (
	"go-snake/akmessage"
	"go-snake/common/myetcd"
	"go-snake/common/rpcBase"
	"context"

	"github.com/Peakchen/xgameCommon/akLog"

)

type WebcontrolRpc struct {
	*myetcd.RpcMessageNode

}

func newWebcontrolRpc()*WebcontrolRpc{
	return &WebcontrolRpc{
		&myetcd.RpcMessageNode{
			MsgNodes: map[akmessage.RPCMSG]*rpcBase.RpcMessage{
			},
			NodeName: rpcBase.RPC_SIMULATION,
		},
	}
}

func (this *WebcontrolRpc) Send(ctx context.Context, in *akmessage.RpcRequest)(*akmessage.RpcResponse, error){
	
	akLog.FmtPrintln("rpc call: ", this.NodeName)

	return this.RpcMessageNode.Call(ctx, in)
}