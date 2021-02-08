package myetcd

import (
	"google.golang.org/grpc"

)

type NodeInfo struct {
	value   string
	session *grpc.ClientConn
}

//rpc call back
type EtcdRpc interface {
	Name() string
}

//dail node
type EtcdNodeIF interface {
	Name() string
	Update(k, v string) error
	Delete(k string) error
	Connect(addr string) (*grpc.ClientConn, error)
	GetNodeConn(name string) *grpc.ClientConn
}
