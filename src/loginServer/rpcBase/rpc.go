package rpcBase

import (
	"go-snake/akmessage"
	"go-snake/common/myetcd"

	"github.com/Peakchen/xgameCommon/akLog"

)

type LoginRpc struct {
	myetcd.RpcMessageNode

}

func newLoginRpc()*LoginRpc{
	return &LoginRpc{
		msgNodes: map[akmessage.RPCMSG]*logicBase.RpcMessage{

		},
		name: logicBase.RPC_LOGIN,
	}
}

func (this *LoginRpc) CallBackxxxx(msg interface{})(*akmessage.RpcResponse, error){
	akLog.FmtPrintln("rpc call: ", this.Name())
	return &akmessage.RpcResponse{}, nil
}
