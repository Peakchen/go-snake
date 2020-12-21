package tcpNet

import (
	"go-snake/akmessage"
	"go-snake/common"
	"go-snake/common/mixNet"
	"net"
	"os"

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
	cli.connect(host, mixNet.GetSessionMgr())
}

func (this *myTcpClient) connect(host string, mgr mixNet.SessionMgrIf) {
	os.Setenv("GOTRACEBACK", "crash")

	tcpAddr, err := net.ResolveTCPAddr("tcp4", host)
	if err != nil {
		akLog.Fail("resolve tcp error: ", err.Error(), host)
		return
	}

	c, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		akLog.Fail("net dial err: ", err, tcpAddr)
		return
	}

	NewTcpSession(c, this.st, this.stop, mgr, this.extFns)

	common.DosafeRoutine(func() {
		this.checkDisconnect(host, mgr)
	}, func() {
		os.Exit(1)
	})
}

func (this *myTcpClient) checkDisconnect(host string, mgr mixNet.SessionMgrIf) {
	for {
		select {
		case <-this.stop:
			this.connect(host, mgr)
		}
	}
}
