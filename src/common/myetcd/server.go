package myetcd

import (
    "context"
    "time"
    "net"
    "go-snake/akmessage"
    "go-snake/common"
    "fmt"
"os"

    "go.etcd.io/etcd/clientv3"
	//"github.com/coreos/etcd/clientv3"
    "github.com/Peakchen/xgameCommon/akLog"
	//"go.etcd.io/etcd/client/v3"
    "go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
    "google.golang.org/grpc"
    "google.golang.org/grpc/grpclog"

)

type etcdServer struct {
    etcdAddr string
    nodeAddr string
}

func NewEtcdServer(etcdHost, nodeHost string, name string, service akmessage.RpcServer){
	es := &etcdServer{
        etcdAddr: etcdHost,
        nodeAddr: nodeHost,
    }
    clientv3.SetLogger(grpclog.NewLoggerV2(os.Stderr, os.Stderr, os.Stderr))
	common.DosafeRoutine(func(){
        es.accept(service)
        }, func(){})
    common.DosafeRoutine(func(){
        es.keepalive(name)
        }, func(){})
}

func (this *etcdServer) accept(service akmessage.RpcServer){
	listener, err := net.Listen("tcp", this.nodeAddr)
    if err != nil {
        akLog.Fail("ListenTCP error:", err)
    }
    s := grpc.NewServer()
    akmessage.RegisterRpcServer(s, service)
    if err := s.Serve(listener); err != nil {
		akLog.Fail("failed to serve: %v", err)
	}
}

func (this *etcdServer) keepalive(name string){
    client, err := clientv3.New(clientv3.Config{
        Endpoints:   []string{this.etcdAddr},
        DialTimeout: 5 * time.Second,
    })
    if err != nil {
		akLog.Fail("create etcd object fail, err", err)
		return
    }
 
    kv := clientv3.NewKV(client)
    lease := clientv3.NewLease(client)
    var curLeaseId clientv3.LeaseID
 
	tick := time.NewTicker(time.Second*5)
	defer tick.Stop()

    common.Dosafe(func(){
        for range tick.C{
            if curLeaseId == 0 {
                leaseResp, err := lease.Grant(context.TODO(), 10)
                if err != nil {
                    akLog.Error("err: ", err)
                    continue
                }
    
                key := name + ":"+ fmt.Sprintf("%d", leaseResp.ID)
                if _, err := kv.Put(context.TODO(), key, this.nodeAddr, clientv3.WithLease(leaseResp.ID)); err != nil {
                    akLog.Error("err: ", key, err)
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
    }, nil)
}