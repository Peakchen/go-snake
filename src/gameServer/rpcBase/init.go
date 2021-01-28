package rpcBase

func RunRpClient(host string){
	myetcd.NewEtcdClient(host, 5, logicBase.RPC_GAME)
}

func RunRpcServer(host string){
	myetcd.NewEtcdServer(host, newGameRpc())
}

func RunRpc(host string){
	RunRpClient(host)
	RunRpcServer(host)
}