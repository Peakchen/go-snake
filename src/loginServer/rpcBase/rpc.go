package rpcBase

import (
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

func (this *LoginRpc) CallBackxxxx(ctx context.Context,msg interface{})(*akmessage.RpcResponse, error){
	akLog.FmtPrintln("rpc call: ", this.NodeName)
	return &akmessage.RpcResponse{}, nil
}
