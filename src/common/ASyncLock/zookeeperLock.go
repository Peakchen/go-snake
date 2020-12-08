package AsyncLock

//add by stefan
// zookeeper version for low frequency and long duration.

import (
	"github.com/samuel/go-zookeeper/zk"
	//"net"
	"time"
	"sync"
)

var (
	zkconn   *zk.Conn
	_zklocks  sync.Map
)

//ip -> ip:port
func NewZKLock(ips []string) {
	var err error
	zkconn, _, err = zk.Connect(ips, time.Second) //default 1s
	if err != nil {
		panic(err)
	}
}

func AddZKLock(key, Name string) (succ bool) {
	lockKey := key + ":" + Name
	zl := zk.NewLock(zkconn, "/"+lockKey, zk.WorldACL(zk.PermAll))
	if err := zl.Lock(); err != nil {
		panic(err)
	}

	succ = true
	_zklocks.Store(lockKey, zl)
	return
}

func ReleaseZKLock(key, Name string) {
	lockKey := key + ":" + Name
	data, exist := _zklocks.Load(lockKey)
	if exist && data != nil{
		_zklocks.Delete(lockKey)
		lock := data.(*zk.Lock)
		lock.Unlock()
	}
}
