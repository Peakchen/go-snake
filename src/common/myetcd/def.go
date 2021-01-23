package myetcd

type NodeInfo struct {
	value string
	session net.Conn
}

type EtcdNode struct{
	name string
	nodes map[string]*NodeInfo
	mx sync.Mutex
}

type EtcdCallService interface{
	Name() string
}

//dail node rpc call back.
type EtcdClient interface{
	Name()string
	Update(k, v string)
	Delete(k string)
	Connect(addr string)
}
