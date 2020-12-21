package evt

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
}

func NewEvtMgr() *EvtMgr {
	evtM := &EvtMgr{
		fnch:    make(chan func(), 1000),
		stop:    make(chan struct{}),
		content: make(chan *EventContent, 1000),
	}

	return evtM
}

func (this *EvtMgr) SendEvt(fn func()) {
	this.fnch <- fn
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

			case ET_EVENT:

			default:
				akLog.Error("invalid event type. ")
			}
		case <-this.stop:
			break loop
		}
	}

}
