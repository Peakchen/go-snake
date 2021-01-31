package rpcBase

import (
	"go-snake/akmessage"
	"go-snake/common/myetcd"
	"go-snake/common/logicBase"
	"context"

	"github.com/Peakchen/xgameCommon/akLog"

)

type GameRpc struct {
	*myetcd.RpcMessageNode

}

func newGameRpc()*GameRpc{
	return &GameRpc{
		&myetcd.RpcMessageNode{
			MsgNodes: map[akmessage.RPCMSG]*logicBase.RpcMessage{
			},
			NodeName: logicBase.RPC_GAME,
		},
	}
}

func (this *GameRpc) CallBackxxxx(ctx context.Context,msg interface{})(*akmessage.RpcResponse, error){
	akLog.FmtPrintln("rpc call: ", this.NodeName)
	return &akmessage.RpcResponse{}, nil
}
