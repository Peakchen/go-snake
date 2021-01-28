package myetcd

import (
	"net"

)

type NodeInfo struct {
	value string
	session net.Conn
}

//rpc call back
type EtcdRpc interface{
	Name() string
}

//dail node
type EtcdNodeIF interface{
	Name()string
	Update(k, v string) error
	Delete(k string) error
	Connect(addr string)(net.Conn, error)
	GetNodeConn(name string) net.Conn
}