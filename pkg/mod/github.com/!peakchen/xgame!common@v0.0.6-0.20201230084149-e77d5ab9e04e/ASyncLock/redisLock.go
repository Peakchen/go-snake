package AsyncLock

//add by stefan
// redis version for high frequency and short duration.

import (
	"gopkg.in/redsync.v1"
	//"github.com/garyburd/redigo/redis"
	//"fmt"
	"time"
)

var (
	_redislocks   = map[string]*redsync.Mutex{}
	_redisSyncobj *redsync.Redsync
)

func NewAsyncLock(pools []redsync.Pool) {
	_redisSyncobj = redsync.New(pools)
}

func AddAsyncLock(key, Name string) {
	lockid := key + ":" + Name
	if _, ok := _redislocks[lockid]; !ok {
		_redislocks[lockid] = _redisSyncobj.NewMutex(lockid,
			redsync.SetExpiry(time.Duration(10*time.Second)),
			redsync.SetRetryDelay(time.Duration(1*time.Second)))
	}
	_redislocks[lockid].Lock()
	return
}

func ReleaseAsyncLock(key, Name string) {
	lockid := key + ":" + Name
	if _, ok := _redislocks[lockid]; !ok {
		return
	}
	_redislocks[lockid].Unlock()
	return
}
