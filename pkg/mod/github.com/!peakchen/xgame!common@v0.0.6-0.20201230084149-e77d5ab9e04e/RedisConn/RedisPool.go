// add by stefan

package RedisConn

import (
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

func Subscriber(pool *redis.Pool, subcontent interface{}) {
	if pool == nil {
		return
	}

	for {
		c := pool.Get()
		psc := redis.PubSubConn{Conn: c}
		psc.Subscribe(subcontent)

		switch recv := psc.Receive().(type) {
		case redis.Message:
			log.Printf("channel: %s, message: %s.", recv.Channel, recv.Data)
		case redis.Subscription:
			log.Printf("channel: %s, Count: %d, kind: %s.", recv.Channel, recv.Count, recv.Kind)
		case error:
			log.Printf("Error.")
			c.Close()
			time.Sleep(Sec_five)
			break
		default:
			log.Printf("default noting.")
		}
	}
}
