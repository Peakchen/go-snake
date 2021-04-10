package timer

/*
	timer for few func call, should not use most sence for performance.
	note: 20200108 by stefan.
*/

import (
	"github.com/Peakchen/xgameCommon/akLog"
	"reflect"
	"time"
)

/*
timer:
	func
	params
	interval
	count
*/

type TAokoTimer struct {
	Func     refect.Value
	Params   []reflect.Value
	Interval int
	Count    int
	curSec   int //after call, then reset to 0
}

type TAokoTimerMgr struct {
	timers []*TAokoTimer
}

var (
	_timerMgr *TAokoTimerMgr
)

func init() {
	_timerMgr = &TAokoTimerMgr{
		timers: []*TAokoTimer{},
	}
}

func Run() {
	var sw sync.WaitGroup
	sw.Add(1)
	go timerloop()
}

func timerloop() {
	tick := time.NewTicker(time.Duration(1) * time.Second)
	var (
		removeTimers = []int{}
	)
	for {
		select {
		case <-tick.C:
			for idx, timercall := range _timerMgr.timers {
				if timercall.curSec >= timercall.Interval && timercall.Count > 0 {
					timercall.Func.Call(timercall.Params)
					timercall.Count--
					timercall.curSec = 0
				} else {
					timercall.curSec++
				}

				if timercall.Count == 0 {
					removeTimers = append(removeTimers, idx)
				}
			}

			for _, timeridx := range removeTimers {
				_timerMgr.timers = append(_timerMgr.timers, _timerMgr.timers[:timeridx], _timerMgr.timers[timeridx+1:]...)
			}
			// clean data
			removeTimers = []int{}
		}
	}
}

func (this *TAokoTimerMgr) register(fun interface{}, interval int, count int, args ...interface{}) {
	var params = []reflect.Value{}
	for _, arg := range args {
		params = append(params, reflect.ValueOf(arg))
	}

	timer := &TAokoTimer{
		Func:     reflect.ValueOf(fun),
		Params:   params,
		Interval: interval,
		Count:    count,
	}

	if cstTimerUpLimit <= len(_timerMgr.timers) {
		akLog.Error("timer list size limit arrives.")
		return
	}

	_timerMgr.timers = append(_timerMgr.timers, timer)
}
