package tcpNet

import (
	"go-snake/akmessage"
	"go-snake/common"
	"go-snake/common/mixNet"
	"net"
	"os"
	"time"

	"github.com/Peakchen/xgameCommon/akLog"
)

type myTcpClient struct {
	stop   chan bool
	st     akmessage.ServerType
	extFns *ExtFnsOption
}

func NewTcpClient(host string, st akmessage.ServerType, extFn ...OptionFn) {
	cli := new(myTcpClient)
	cli.stop = make(chan bool, 1)
	cli.st = st
	cli.extFns = SortOptions(extFn...)
	if !cli.connect(host, mixNet.GetSessionMgr()) {
		cli.stop <- true
	}
	common.DosafeRoutine(func() {
		cli.checkDisconnect(host, mixNet.GetSessionMgr())
	}, nil)
}

func (this *myTcpClient) connect(host string, mgr mixNet.SessionMgrIf) bool {
	os.Setenv("GOTRACEBACK", "crash")

	tcpAddr, err := net.ResolveTCPAddr("tcp4", host)
	if err != nil {
		akLog.Error("resolve tcp error: ", err.Error(), host)
		return false
	}

	c, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		akLog.Error("net dial err: ", err, tcpAddr)
		return false
	}

	NewTcpSession(c, this.st, this.stop, mgr, this.extFns)
	return true
}

func (this *myTcpClient) checkDisconnect(host string, mgr mixNet.SessionMgrIf) {
	for {
		select {
		case <-this.stop:
			if mgr.IsClose() {
				akLog.Info("client close tcp...")
				return
			}
			akLog.FmtPrintln("begin reconnect...")
			if !this.connect(host, mgr) {
				time.Sleep(3 * time.Second)
				this.stop <- true
			}
		}
	}
}
