package evtAsync

import (
	"go-snake/common"
	"time"

	"github.com/Peakchen/xgameCommon/akLog"
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
	stop    chan struct{}
	content chan *EventContent
	eopts   *EventOption
}

func NewEvtMgr(opts ...OptFn) *EvtMgr {
	evtM := &EvtMgr{
		fnch:    make(chan func(), 1000),
		stop:    make(chan struct{}),
		content: make(chan *EventContent, 1000),
		eopts:   loadEventOpts(opts...),
	}
	common.DosafeRoutine(evtM.loop, func() {
		time.Sleep(time.Second)
	})
	return evtM
}

func (this *EvtMgr) SendEvtFn(fn func()) {
	this.fnch <- fn
}

func (this *EvtMgr) SendEvtContent(ec *EventContent) {
	this.content <- ec
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
		case <-this.stop:
			break loop
		}
	}

}
