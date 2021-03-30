package rpcBase

import (
	"go-snake/common/myetcd"
	"go-snake/common/rpcBase"

)

func RunRpClient(etcdhost, nodehost string){
	myetcd.NewEtcdClient(etcdhost, nodehost, 5, rpcBase.RPC_GAME)
}

func RunRpcServer(etcdhost, nodehost string){
	myetcd.NewEtcdServer(etcdhost, nodehost, rpcBase.RPC_GAME, gameRpc)
}

func RunRpc(etcdhost, nodehost string){
	RunRpClient(etcdhost, nodehost)
	RunRpcServer(etcdhost, nodehost)
}