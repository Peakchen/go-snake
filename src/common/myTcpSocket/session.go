package myTcpSocket

import (
	"ak-remote/common"
	"ak-remote/common/mixNet"
	"net"
	"time"

	"github.com/Peakchen/xgameCommon/aktime"

	"github.com/Peakchen/xgameCommon/akLog"

	"github.com/Peakchen/xgameCommon/utils"
)

//for client connect server.

type AkTcpSession struct {
	id     string
	stop   chan bool
	sendCh chan []byte
	conn   *net.TCPConn

	Sessmgr mixNet.SessionMgrIf
	MsgFn   func(*net.TCPConn) bool
}

func NewTcpSession(c *net.TCPConn, stop chan bool, mgr mixNet.SessionMgrIf, fn func(*net.TCPConn) bool) *AkTcpSession {
	session := &AkTcpSession{
		id:      utils.GetUUID(),
		stop:    stop,
		sendCh:  make(chan []byte, 100),
		conn:    c,
		Sessmgr: mgr,
		MsgFn:   fn,
	}
	mgr.AddTcpSession(session.id, session)
	session.handler()
	return session
}

func (this *AkTcpSession) GetSessionID() string {
	return this.id
}

func (this *AkTcpSession) handler() {
	this.conn.SetKeepAlive(true)

	common.DosafeRoutine(func() { this.readloop() }, func() { this.Stop() })
	common.DosafeRoutine(func() { this.writeloop() }, func() { this.Stop() })
}

func (this *AkTcpSession) readloop() {
	defer this.Stop()

	for {
		this.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		if !this.MsgFn(this.conn) {
			return
		}
	}
}

func (this *AkTcpSession) writeloop() {
	defer this.Stop()

	for {
		select {
		case data := <-this.sendCh:
			this.conn.SetWriteDeadline(aktime.Now().Add(15 * time.Second))
			_, err := this.conn.Write(data)
			if err != nil {
				akLog.Error("send data fail, err: ", err)
			}
		}
	}
}

func (this *AkTcpSession) Stop() {
	this.stop <- true
	this.Sessmgr.RemoveTcpSession(this.id)
}

func (this *AkTcpSession) SendMsg(data []byte) {
	this.sendCh <- data
}
