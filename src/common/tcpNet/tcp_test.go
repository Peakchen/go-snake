package tcpNet

import (
	"testing"
	"time"

	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/utils"
)

func TestUID(t *testing.T) {
	uid := utils.GetUUID()
	akLog.FmtPrintln("uid: ", uid, len(uid))
}

func tcploop(ch chan bool) {
	session := &TcpSession{
		stop: ch,
	}
	tick := time.NewTicker(3 * time.Second)
	for {
		select {
		// case <-session.stop:
		// 	akLog.FmtPrintln("session stop...")
		case <-tick.C:
			session.stop <- true
			tick.Stop()
			break
		}
	}
}

func TestCh(t *testing.T) {

	ch := make(chan bool, 1)
	fch := make(chan func(), 1)
	go tcploop(ch)
	for {
		select {
		case f := <-fch:
			f()
		case <-ch:
			akLog.FmtPrintln("session stop...")
			//fch <- func() {
			//akLog.FmtPrintln("session f...")
			time.Sleep(3 * time.Second)
			ch <- true
			//}

		}
	}
}
