package evtAsync

import (
	"go-snake/common"
	"time"

	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/panjf2000/ants/v2"
)

const (
	ET_FUNC = iota
	ET_TASK
	ET_EVENT
)

type EventContent struct {
	ET int
}

type EvtMgr struct {
	fnch    chan func()
	content chan *EventContent
	eopts   *EventOption
	poolfn  *ants.Pool
}

var (
	_exitCh = make(chan bool, 1)
	_main   *EvtMgr
)

func NewMainEvtMgr(opts ...OptFn) {
	p, _ := ants.NewPool(10 * 1024)
	_main = &EvtMgr{
		poolfn: p,
		//fnch:    make(chan func(), 3*1000),
		content: make(chan *EventContent, 1000),
		eopts:   loadEventOpts(opts...),
	}
	common.DosafeRoutine(_main.loop, func() {
		time.Sleep(time.Second)
	})
}

func Stop() {
	_exitCh <- true
}

func SendEvtFn(fn func()) {
	// if len(_main.fnch) >= 3*cap(_main.fnch)/4 {
	// 	akLog.Info("warning fnch...")
	// }
	//_main.fnch <- fn

	_main.poolfn.Submit(fn)
}

func SendEvtContent(ec *EventContent) {
	_main.content <- ec
}

func (this *EvtMgr) loop() {
	tick := time.NewTicker(100 * time.Millisecond)
	defer tick.Stop()

loop:
	for {
		select {
		case <-tick.C:

		case fn := <-this.fnch:
			common.Dosafe(fn, nil)
		case e := <-this.content:
			switch e.ET {
			case ET_FUNC:

			case ET_TASK:
				common.Dosafe(func() {
					this.eopts.TaskHandler(e)
				}, nil)
			case ET_EVENT:
				common.Dosafe(func() {
					this.eopts.EventHandler(e)
				}, nil)
			default:
				akLog.Error("invalid event type: ", e.ET)
			}
		case <-_exitCh:
			break loop
		}
	}

}
