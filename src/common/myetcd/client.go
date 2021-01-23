package myetcd

import (
	"time"

	"github.com/coreos/etcd/clientv3"

)

type etcdClient struct {
	addr string
	cli clientv3.Client
	serviceRef reflect.Value
}

func NewEtcdClient(host string, timeOuts time.Duration, service EtcdClient){
	client, err := clientv3.New(clientv3.Config{
		EndPoints: []string{host},
		DailTimeout: timeOuts,
	})
	if err != nil {
		akLog.Fail("can not create etcd client, host: ", host)
		return
	}
	ec := &etcdClient{
		addr : host,
		cli: client,
		serviceRef: reflect.ValueOf(service),
	}
	common.DosafeRoutine(ec.ListenUpdateNodes, func(){

	})
}

func (this *etcdClient) ListenUpdateNodes(){
	collectNodes := func(){
		kv := clientv3.NewKV(this.cli)
		rsp, err := kv.Get(context.TODO(),  this.serviceRef.MethodByName("Name").Call(nil), clientv3.WithPrefix())
		if err != nil {
			akLog.Error("get etcd nodes fail, node: ", this.serviceRef.MethodByName("Name").Call(nil))
			return
		}
		
		for _, kv := range rsp.Kvs {
			this.serviceRef.MethodByName("Update").Call([]reflect.Value{
				reflect.ValueOf(Kv.Key),
				reflect.ValueOf(Kv.Value),
			})
		}
	}

	common.Dosafe(collectNodes, func(){})

	ticker := time.NewTicker(3*time.Second)
	defer ticker.Stop()

	watcher := clientv3.NewWatcher(this.cli)
	watchCh := watcher.Watch(context.TODO(), this.serviceRef.MethodByName("Name").Call(nil), clientv3.WithPrefix())
 	for {
		select {
		case <-ticker.C:
			common.Dosafe(collectNodes, func(){})
		case chs := <-watchCh:
			for rsp := range chs{
				for _,e := range rsp.Events{
					switch e.Type{
					case mvccpb.PUT:
						this.serviceRef.MethodByName("Update").Call([]reflect.Value{
							reflect.ValueOf(e.Kv.Key),
							reflect.ValueOf(e.Kv.Value),
						})
					case mvccpb.DELETE:
						this.serviceRef.MethodByName("Delete").Call([]reflect.Value{
							reflect.ValueOf(e.Kv.Key),
						})
					}
				}	 
			}
		}
	}
}