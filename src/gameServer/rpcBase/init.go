package rpcBase

import (
	"go-snake/common/myetcd"
	"go-snake/common/logicBase"

)

func RunRpClient(etcdhost, nodehost string){
	myetcd.NewEtcdClient(etcdhost, nodehost, 5, logicBase.RPC_GAME)
}

func RunRpcServer(etcdhost, nodehost string){
	myetcd.NewEtcdServer(etcdhost, nodehost, logicBase.RPC_GAME, newGameRpc())
}

func RunRpc(etcdhost, nodehost string){
	RunRpClient(etcdhost, nodehost)
	RunRpcServer(etcdhost, nodehost)
}