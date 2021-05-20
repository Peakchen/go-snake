package myetcd

import (
	"context"
	"errors"
	"go-snake/akmessage"
	"go-snake/common"
	"go-snake/common/rpcBase"
	"os"
	"reflect"
	//"sync"
	"time"

	"go.etcd.io/etcd/clientv3"
	//"github.com/coreos/etcd/clientv3"
	"google.golang.org/protobuf/proto"
	//"github.com/coreos/etcd/mvcc/mvccpb"
	//"go.etcd.io/etcd/mvcc/mvccpb"
	//"go.etcd.io/etcd/client/v3"
	//"go.etcd.io/etcd/api/v3/mvccpb"
	"github.com/Peakchen/xgameCommon/akLog"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"

)

type etcdService struct {
	addr   string
	cli    *clientv3.Client
	RpcRef reflect.Value
}

func (this *etcdService) GetCli() *clientv3.Client {
	return this.cli
}

func (this *etcdService) GetAddr() string {
	return this.addr
}

type nodeService struct {
	Node    string
	Service *etcdService
}

var (
	_ec    *etcdService
	Nodes map[string]*nodeService
)

func NewEtcdClient(host, nodehost string, timeOuts int, service string) {

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{host},
		DialTimeout: time.Duration(timeOuts) * time.Second,
	})
	
	if err != nil {
		akLog.Fail("can not create etcd client, host: ", host)
		return
	}
	clientv3.SetLogger(grpclog.NewLoggerV2(os.Stderr, os.Stderr, os.Stderr))

	_ec = &etcdService{
		addr:   host,
		cli:    client,
		RpcRef: reflect.ValueOf(NewRpcMgr(service, nodehost)),
	}

	others := rpcBase.GetAllNode()
	Nodes = make(map[string]*nodeService, len(others)-1)
	mynode := _ec.RpcRef.MethodByName("Name").Call([]reflect.Value{})[0].Interface().(string)

	for _, node := range others {

		if mynode == node {
			continue
		}

		Nodes[node] = &nodeService{
			Service: _ec,
			Node:    node,
		}

		common.DosafeRoutine(Nodes[node].ListenUpdateNodes, func() {})
	}
}

func (this *nodeService) ListenUpdateNodes() {
	collectNodes := func() {
		kv := clientv3.NewKV(this.Service.GetCli())
		rsp, err := kv.Get(context.TODO(), this.Node, clientv3.WithPrefix())
		if err != nil {
			akLog.Error("get etcd nodes fail, node: ", this.Node)
			return
		}

		for _, kv := range rsp.Kvs {
			akLog.FmtPrintln("node kvs: ", string(kv.Key), string(kv.Value))
			ret := this.Service.RpcRef.MethodByName("Update").Call([]reflect.Value{
				reflect.ValueOf(string(kv.Key)),
				reflect.ValueOf(string(kv.Value)),
			})
			if ret[0].Interface() == nil {
				continue
			}
			if ret[0].Interface().(error) != nil {
				akLog.Error("service update: ", ret[0].Interface().(error))
			}
		}
	}

	common.Dosafe(func() {
		collectNodes()
	}, nil)

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	watcher := clientv3.NewWatcher(this.Service.GetCli())
	watchCh := watcher.Watch(context.TODO(), this.Node, clientv3.WithPrefix())
	for {
		select {
		case <-ticker.C:
			common.Dosafe(func() {
				collectNodes()
			}, nil)
		case rsp := <-watchCh:
			for _, e := range rsp.Events {
				switch e.Type {
				case mvccpb.PUT:
					ret := this.Service.RpcRef.MethodByName("Update").Call([]reflect.Value{
						reflect.ValueOf(string(e.Kv.Key)),
						reflect.ValueOf(string(e.Kv.Value)),
					})
					if ret[0].Interface() == nil {
						continue
					}
					if ret[0].Interface().(error) != nil {
						akLog.Error("service update: ", ret[0].Interface().(error))
					}
				case mvccpb.DELETE:
					ret := this.Service.RpcRef.MethodByName("Delete").Call([]reflect.Value{
						reflect.ValueOf(string(e.Kv.Key)),
					})

					if ret[0].Interface() == nil {
						continue
					}
					if ret[0].Interface().(error) != nil {
						akLog.Error("service delete: ", ret[0].Interface().(error))
					}
				}
			}
		}
	}
}

func (this *nodeService) Call(name string, msgID akmessage.RPCMSG, arg interface{}) (*akmessage.RpcResponse, error) {

	ret := this.Service.RpcRef.MethodByName("GetNodeConn").Call([]reflect.Value{
		reflect.ValueOf(name),
	})

	if ret[0].Interface() == nil {
		return nil, errors.New("can not get node session or not create node session.")
	}

	conn := ret[0].Interface().(*grpc.ClientConn)
	c := akmessage.NewRpcClient(conn)
	if c == nil {
		return nil, errors.New("invalid node session.")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	data, err := proto.Marshal(arg.(protoreflect.ProtoMessage))
	if err != nil {
		return nil, err
	}

	var (
		rsp *akmessage.RpcResponse
	)

	common.Dosafe(func(){
		
		rsp, err = c.Send(ctx, &akmessage.RpcRequest{
			MsgId: 		msgID,
			ReqData: 	data,
		})
	
	}, nil)
	
	return rsp, err
}

func Call(name string, msgID akmessage.RPCMSG, arg interface{}) (*akmessage.RpcResponse, error) {
	
	if Nodes[name] == nil {
		return nil, errors.New("can not find node.")
	}

	return Nodes[name].Call(name, msgID, arg)
}
