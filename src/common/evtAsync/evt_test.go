package evtAsync

import (
	"fmt"
	"testing"
	"time"
)

var fff bool
var c int
var box = make(chan int, 10)

func fix(ch chan<- bool) {
	t := time.NewTicker(time.Second)

	for range t.C {
		ch <- true
		c++
		if c >= 20 {
			fff = true
		}
		box <- c
		//fff = true
	}

}

func TestFOR_SELECT(t *testing.T) {
	ch := make(chan bool, 1)
	go fix(ch)

	for {
		select {
		case <-ch:
			//ch = nil
			fmt.Println("aaa")
			//return
			if fff {
				return
			}
		default:
			select {
			case <-ch:
				//ch = nil
				fmt.Println("bbb")
				if fff {
					return
				}
				//return
			case d := <-box:
				fmt.Println("box: ", d)
			default:

			}
		}
	}
}
