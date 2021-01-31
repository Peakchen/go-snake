package myetcd

import (
	"context"
	"errors"
	"go-snake/akmessage"
	"go-snake/common"
	"reflect"
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
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"

)

type etcdService struct {
	addr   string
	cli    *clientv3.Client
	rpcRef reflect.Value
}

var (
	_ec *etcdService
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
	_ec = &etcdService{
		addr:   host,
		cli:    client,
		rpcRef: reflect.ValueOf(newRpcMgr(service, nodehost)),
	}
	common.DosafeRoutine(_ec.ListenUpdateNodes, func() {})
}

func (this *etcdService) ListenUpdateNodes() {
	nodename := this.rpcRef.MethodByName("Name").Call([]reflect.Value{})[0].Interface().(string)
	collectNodes := func() {
		kv := clientv3.NewKV(this.cli)
		rsp, err := kv.Get(context.TODO(), nodename, clientv3.WithPrefix())
		if err != nil {
			akLog.Error("get etcd nodes fail, node: ", nodename)
			return
		}

		for _, kv := range rsp.Kvs {
			akLog.FmtPrintln("node kvs: ", string(kv.Key), string(kv.Value))
			ret := this.rpcRef.MethodByName("Update").Call([]reflect.Value{
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

	watcher := clientv3.NewWatcher(this.cli)
	watchCh := watcher.Watch(context.TODO(), nodename, clientv3.WithPrefix())
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
					ret := this.rpcRef.MethodByName("Update").Call([]reflect.Value{
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
					ret := this.rpcRef.MethodByName("Delete").Call([]reflect.Value{
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

func (this *etcdService) Call(name string, arg interface{}) (*akmessage.RpcResponse, error) {
	ret := this.rpcRef.MethodByName("GetNodeConn").Call([]reflect.Value{
		reflect.ValueOf(arg),
	})
	if ret[0].Interface() == nil {
		return nil, errors.New("can not get node session or not create node session.")
	}
	conn := ret[0].Interface().(*grpc.ClientConn)
	c := akmessage.NewRpcClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	data, err := proto.Marshal(arg.(protoreflect.ProtoMessage))
	if err != nil {
		return nil, err
	}
	rsp, err := c.Send(ctx, &akmessage.RpcRequest{
		ReqData: data,
	})
	return rsp, err
}

func Call(name string, arg interface{}) (*akmessage.RpcResponse, error) {
	return _ec.Call(name, arg)
}
