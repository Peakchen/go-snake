package myTcpSocket

import (
	"ak-remote/common"
	"ak-remote/common/config"
	"ak-remote/common/mixNet"
	"net"
	"os"

	"github.com/Peakchen/xgameCommon/akLog"
)

type myTcpServer struct {
}

func NewTcpServer() {
	tcpSvr := &myTcpServer{}
	tcpSvr.StartTcpListen(config.GetServerConfig().Host, mixNet.GetSessionMgr())
}

func (this *myTcpServer) StartTcpListen(host string, mgr mixNet.SessionMgrIf) {
	os.Setenv("GOTRACEBACK", "crash")
	tcpAddr, err := net.ResolveTCPAddr("tcp4", host)
	if err != nil {
		akLog.Fail("resole tcp, host: ", host)
		return
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		akLog.Fail("listen fail, addr: ", tcpAddr)
		return
	}

	common.DosafeRoutine(func() {
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				akLog.Error("accept fail,err: ", err)
				continue
			}
			NewTcpSession(conn, make(chan bool, 1), mgr, ServerMsgProc)
		}
	}, func() {
		os.Exit(1)
	})
}
