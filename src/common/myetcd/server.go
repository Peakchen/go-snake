package myetcd

import (
	"strings"
	"github.com/coreos/etcd/clientv3"
)

type etcdServer struct {
	addr string
	name string 

}

func NewEtcdServer(host, src string, service EtcdCallService){
	rpc.RegisterName(service.Name(), service)
	es := &etcdServer{
		addr: host,
		name: src,
	}
	common.DosafeRoutine(es.accept, func(){

	})

}

func (this *etcdServer) accept(service EtcdCallService){
	listener, err := net.Listen("tcp", this.addr)
    if err != nil {
        log.Fatal("ListenTCP error:", err)
    }

    conn, err := listener.Accept()
    if err != nil {
        log.Fatal("Accept error:", err)
	}
	rpc.ServeConn(conn)
}

func (this *etcdServer) keepalive(service EtcdCallService){
    client, err := clientv3.New(clientv3.Config{
        Endpoints:   []string{this.addr},
        DialTimeout: 5 * time.Second,
    })
    if err != nil {
		akLog.Fail("create etcd object fail, err", err)
		return
    }
 
    kv := clientv3.NewKV(client)
    lease := clientv3.NewLease(client)
    var curLeaseId clientv3.LeaseID
 
	tick := time.NewTicker(time.Second)
	defer tick.Stop()

    for range tick.C{
        if curLeaseId == 0 {
            leaseResp, err := lease.Grant(context.TODO(), 10)
            if err != nil {
				akLog.Error("err: ", err)
				continue
            }
 
            key := service.Name() + fmt.Sprintf("%d", leaseResp.ID)
            if _, err := kv.Put(context.TODO(), key, this.addr, clientv3.WithLease(leaseResp.ID)); err != nil {
                akLog.Error("err: ", err)
				continue
            }
            curLeaseId = leaseResp.ID
        } else {
            if _, err := lease.KeepAliveOnce(context.TODO(), curLeaseId); err == rpctypes.ErrLeaseNotFound {
                curLeaseId = 0
                continue
            }
        }
    }
}