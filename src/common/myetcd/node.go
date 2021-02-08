package myetcd

import (
	"sync"
	"errors"

	"google.golang.org/grpc"

)

type RpcNode struct {
	nodes map[string]*NodeInfo
	mutx sync.RWMutex
	name string
	nodeAddr string
}

func NewRpcMgr(name, addr string)*RpcNode{
	return &RpcNode{
		nodes: make(map[string]*NodeInfo),
		name: name,
		nodeAddr: addr,
	}
}

func (this *RpcNode) Name()string{
	return this.name
}

func (this *RpcNode) Update(k, v string) error {
	if this.nodeAddr == v {
		return nil
	}
	lis, err := this.Connect(v)
	if lis == nil {
		return err
	}
	this.mutx.Lock()
	defer this.mutx.Unlock()

	this.nodes[k] = &NodeInfo{
		value: v,
		session: lis,
	}
	return nil
}

func (this *RpcNode) Delete(k string) error{
	this.mutx.Lock()
	defer this.mutx.Unlock()
	
	node, ok := this.nodes[k]
	if !ok {
		return errors.New("can not find rpc node.")
	}
	node.session.Close()
	delete(this.nodes, k)
	return nil
}

func (this *RpcNode) Connect(addr string) (*grpc.ClientConn,error){
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (this *RpcNode) GetNodeConn(name string)*grpc.ClientConn{
	this.mutx.RLock()
	defer this.mutx.RUnlock()
	
	node, ok := this.nodes[name]
	if !ok {
		return nil
	}
	return node.session
}