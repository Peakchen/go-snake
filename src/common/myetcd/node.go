package myetcd

import (
	"sync"
	"errors"
	"net"

	"google.golang.org/grpc"

)

type RpcNode struct {
	nodes map[string]*NodeInfo
	mutx sync.RWMutex
	name string
}

func newRpcMgr(name string)*RpcNode{
	return &RpcNode{
		nodes: make(map[string]string),
		name: name,
	}
}

func (this *RpcNode) Name()string{
	return this.name
}

func (this *RpcNode) Update(k, v string) error {
	lis, err := this.Connect(v)
	if lis == nil {
		return err
	}
	this.mutx.Lock()
	defer this.mutx.UnLock()

	this.nodes[k] = &myetcd.NodeInfo{
		value: v,
		session: lis,
	}
	return nil
}

func (this *RpcNode) Delete(k string) error{
	this.mutx.Lock()
	defer this.mutx.UnLock()
	
	node, ok := this.nodes[k]
	if !ok {
		return errors.New("can not find rpc node.")
	}
	node.session.Close()
	delete(this.nodes, k)
}

func (this *RpcNode) Connect(addr string) (net.Conn,error){
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, errors.New("connect node fail.")
	}
	return conn, nil
}

func (this *RpcNode) GetNodeConn(name string)net.Conn{
	this.mutx.RLock()
	defer this.mutx.RUnLock()
	
	node, ok := this.nodes[name]
	if !ok {
		return nil
	}
	return node.session
}