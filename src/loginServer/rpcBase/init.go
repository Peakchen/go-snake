package rpcBase

func RunRpClient(host string){
	myetcd.NewEtcdClient(host, 5, logicBase.RPC_LOGIN)
}

func RunRpcServer(host string){
	myetcd.NewEtcdServer(host, newLoginRpc())
}

func RunRpc(host string){
	RunRpClient(host)
	RunRpcServer(host)
}