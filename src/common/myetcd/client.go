package myetcd

import (
	"time"
	"reflect"
	"context"
	"go-snake/akmessage"
	"errors"
	"go-snake/common"

	"go.etcd.io/etcd/clientv3"
	//"github.com/coreos/etcd/clientv3"
	"google.golang.org/protobuf/proto"
	//"github.com/coreos/etcd/mvcc/mvccpb"
	//"go.etcd.io/etcd/mvcc/mvccpb"
	//"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/api/v3/mvccpb"
	"github.com/Peakchen/xgameCommon/akLog"

)

type etcdService struct {
	addr string
	cli *clientv3.Client
	rpcRef reflect.Value
}

var (
	_ec *etcdService
)

func NewEtcdClient(host string, timeOuts int, service string){
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{host},
		DialTimeout: time.Duration(timeOuts)*time.Second,
	})
	if err != nil {
		akLog.Fail("can not create etcd client, host: ", host)
		return
	}
	_ec = &etcdService{
		addr : host,
		cli: client,
		rpcRef: reflect.ValueOf(newRpcMgr(service)),
	}
	common.DosafeRoutine(_ec.ListenUpdateNodes, func(){})
}

func (this *etcdService) ListenUpdateNodes(){
	collectNodes := func(){
		kv := clientv3.NewKV(this.cli)
		rsp, err := kv.Get(context.TODO(),  this.rpcRef.MethodByName("Name").Call([]reflect.Value{})[0], clientv3.WithPrefix())
		if err != nil {
			akLog.Error("get etcd nodes fail, node: ", this.rpcRef.MethodByName("Name").Call([]reflect.Value{})[0])
			return
		}
		
		for _, kv := range rsp.Kvs {
			ret := this.rpcRef.MethodByName("Update").Call([]reflect.Value{
				reflect.ValueOf(kv.Key),
				reflect.ValueOf(kv.Value),
			})
			if ret[0].Interface().(error) != nil {
				akLog.Error("service update: ", ret[0].Interface().(error))
			}
		}
	}

	common.Dosafe(func(){ 
		collectNodes()
		}, nil)

	ticker := time.NewTicker(3*time.Second)
	defer ticker.Stop()

	watcher := clientv3.NewWatcher(this.cli)
	watchCh := watcher.Watch(context.TODO(), this.rpcRef.MethodByName("Name").Call([]reflect.Value{}), clientv3.WithPrefix())
 	for {
		select {
		case <-ticker.C:
			common.Dosafe(func(){
				collectNodes()
				}, nil)
		case rsp := <-watchCh:
			for _,e := range rsp.Events{
				switch e.Type{
				case mvccpb.PUT:
					ret := this.rpcRef.MethodByName("Update").Call([]reflect.Value{
						reflect.ValueOf(e.Kv.Key),
						reflect.ValueOf(e.Kv.Value),
					})
					if ret[0].Interface().(error) != nil {
						akLog.Error("service update: ", ret[0].Interface().(error))
					}
				case mvccpb.DELETE:
					ret := this.rpcRef.MethodByName("Delete").Call([]reflect.Value{
						reflect.ValueOf(e.Kv.Key),
					})

					if ret[0].Interface().(error) != nil {
						akLog.Error("service delete: ", ret[0].Interface().(error))
					}
				}
			}	 
			
		}
	}
}

func (this *etcdService) Call(name string, arg interface{})(*akmessage.RpcResponse, error){
	ret := this.rpcRef.MethodByName("GetNodeConn").Call([]reflect.Value{
			reflect.ValueOf(arg),
		})
	if ret.Interface() == nil {
		return nil,  errors.New("can not get node session or not create node session.")
	}	
	conn := ret.Interface().(net.Conn)
	c := akmessage.NewRpcClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	
	data, err := proto.Marshal(arg)
	if err != nil {
		return nil, err
	}
	rsp, err := c.Send(ctx, &akmessage.RpcRequest{
		ReqData: data,
	})
	return rsp, err
}

func Call(name string, arg interface{}) (*akmessage.RpcResponse, error){
	return _ec.Call(name, arg)
}