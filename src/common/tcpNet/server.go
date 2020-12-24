package tcpNet

import (
	"go-snake/akmessage"
	"go-snake/common"
	"go-snake/common/mixNet"
	"net"
	"os"

	"github.com/Peakchen/xgameCommon/akLog"
)

type myTcpServer struct {
}

func NewTcpServer(host string, st akmessage.ServerType, extFn ...OptionFn) {
	tcpSvr := &myTcpServer{}
	tcpSvr.StartTcpListen(host, mixNet.GetSessionMgr(), st, SortOptions(extFn...))
}

func (this *myTcpServer) StartTcpListen(host string, mgr mixNet.SessionMgrIf, st akmessage.ServerType, extFns *ExtFnsOption) {
	os.Setenv("GOTRACEBACK", "crash")
	common.DosafeRoutine(func() {
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

		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				akLog.Error("accept fail,err: ", err)
				continue
			}
			if mgr.IsClose() {
				akLog.Info("server close tcp...")
				return
			}
			akLog.FmtPrintln("new tcp session, addr: ", conn.RemoteAddr())
			NewTcpSession(conn, st, make(chan bool, 1), mgr, extFns)
		}
	}, nil)
}
