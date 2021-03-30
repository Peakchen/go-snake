package rpcBase

import (
	"go-snake/common/myetcd"
	"go-snake/common/rpcBase"

)

func RunRpClient(etcdhost, nodehost string){
	myetcd.NewEtcdClient(etcdhost, nodehost, 5, rpcBase.RPC_LOGIN)
}

func RunRpcServer(etcdhost, nodehost string){
	myetcd.NewEtcdServer(etcdhost, nodehost, rpcBase.RPC_LOGIN, newLoginRpc())
}

func RunRpc(etcdhost, nodehost string){
	RunRpClient(etcdhost, nodehost)
	RunRpcServer(etcdhost, nodehost)
}