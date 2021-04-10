package rpcBase

import (
	"go-snake/common/myetcd"
	"go-snake/app/application"
)

func RunRpClient(etcdhost, nodehost string){
	myetcd.NewEtcdClient(etcdhost, nodehost, 5, application.GetAppName())
}

func RunRpcServer(etcdhost, nodehost string){
	myetcd.NewEtcdServer(etcdhost, nodehost, application.GetAppName(), newSimulationRpc())
}

func RunRpc(etcdhost, nodehost string){
	RunRpClient(etcdhost, nodehost)
	RunRpcServer(etcdhost, nodehost)
}