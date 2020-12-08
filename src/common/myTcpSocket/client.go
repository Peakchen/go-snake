package myTcpSocket

import (
	"ak-remote/common"
	"ak-remote/common/config"
	"ak-remote/common/mixNet"
	"net"
	"os"

	"github.com/Peakchen/xgameCommon/akLog"
)

type myTcpClient struct {
	stop chan bool
}

func NewTcpClient() {
	cli := new(myTcpClient)
	cli.stop = make(chan bool, 1)
	cli.connect(config.GetServerConfig().Host, mixNet.GetSessionMgr())
}

func (this *myTcpClient) connect(host string, mgr mixNet.SessionMgrIf) {
	os.Setenv("GOTRACEBACK", "crash")

	tcpAddr, err := net.ResolveTCPAddr("tcp4", host)
	if err != nil {
		akLog.Fail("resolve tcp error: ", err.Error())
		return
	}

	c, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		akLog.Fail("net dial err: ", err)
		return
	}

	NewTcpSession(c, this.stop, mgr, ClientMsgProc)

	go common.DosafeRoutine(func() {
		this.checkDisconnect(host, mgr)
	}, nil)
}

func (this *myTcpClient) checkDisconnect(host string, mgr mixNet.SessionMgrIf) {
	for {
		select {
		case <-this.stop:
			this.connect(host, mgr)
		}
	}
}
