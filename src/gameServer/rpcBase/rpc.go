package rpcBase

import (
	"go-snake/akmessage"
	"go-snake/common/myetcd"

	"github.com/Peakchen/xgameCommon/akLog"

)

type GameRpc struct {
	myetcd.RpcMessageNode

}

func newGameRpc()*GameRpc{
	return &GameRpc{
		msgNodes: map[akmessage.RPCMSG]*logicBase.RpcMessage{
			
		},
		name: logicBase.RPC_GAME,
	}
}

func (this *GameRpc) CallBackxxxx(msg interface{})(*akmessage.RpcResponse, error){
	akLog.FmtPrintln("rpc call: ", this.Name())
	return &akmessage.RpcResponse{}, nil
}
