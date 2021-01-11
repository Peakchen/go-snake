package akOrm

import (
	"container/list"
	"fmt"
	"sync"
	"testing"
	"time"
)

var operlist = list.New()
var mutex sync.RWMutex
var oper = make(chan *DBOper, 1000)

func consumer() {
	tick := time.NewTicker(500 * time.Millisecond)
	defer tick.Stop()

	for {
		select {
		case <-oper:
			fmt.Println("recv: ", time.Now().Unix())
		case <-tick.C:
			mutex.Lock()
			var count = operlist.Len()
			fmt.Println("now: ", count, time.Now().Unix())
			for i := 0; i < count; i++ {
				op := operlist.Front()
				v := op.Value.(*DBOper)
				if v == nil {
					fmt.Println("elem is nil.")
				}
				operlist.Remove(op)
			}
			fmt.Println("left: ", operlist.Len(), time.Now().Unix())
			mutex.Unlock()
		}
	}
}

func producer() {
	tick := time.NewTicker(5 * time.Millisecond)
	defer tick.Stop()

	for range tick.C {
		// oper <- &DBOper{
		// 	Type: ORM_UPDATE,
		// 	Data: nil,
		// }
		mutex.Lock()
		operlist.PushBack(&DBOper{
			Type: ORM_UPDATE,
			Data: nil,
		})
		mutex.Unlock()
	}
}

func TestDoubleList(t *testing.T) {
	go consumer()
	go producer()
	select {}
}
