package tcpNet

import (
	"go-snake/akmessage"
	"go-snake/common"
	"go-snake/common/mixNet"
	"net"
	"strings"
	"time"

	"github.com/Peakchen/xgameCommon/aktime"

	"github.com/Peakchen/xgameCommon/akLog"

	"github.com/Peakchen/xgameCommon/utils"
)

//for client connect server.

type TcpSession struct {
	id     string
	stop   chan bool
	sendCh chan []byte
	conn   *net.TCPConn

	Sessmgr mixNet.SessionMgrIf
	extFns  *ExtFnsOption

	uid  int64
	svrt akmessage.ServerType
	clit akmessage.ServerType
}

func NewTcpSession(c *net.TCPConn, st akmessage.ServerType, stop chan bool, mgr mixNet.SessionMgrIf, extFn *ExtFnsOption) *TcpSession {
	session := &TcpSession{
		id:      strings.Trim(utils.GetUUID(), " "),
		stop:    stop,
		sendCh:  make(chan []byte, 100),
		conn:    c,
		Sessmgr: mgr,
		extFns:  extFn,
		svrt:    st,
	}
	mgr.AddTcpSession(session.id, session)
	session.handler()
	return session
}

func (this *TcpSession) GetSessionID() string {
	return this.id
}

func (this *TcpSession) handler() {
	this.conn.SetKeepAlive(true)

	common.DosafeRoutine(func() { this.readloop() }, func() { this.Stop() })
	common.DosafeRoutine(func() { this.writeloop() }, func() { this.Stop() })
	if this.extFns.CS_HeartBeat != nil {
		common.DosafeRoutine(func() { this.heartBeat() }, func() { this.Stop() })
	}
}

func (this *TcpSession) heartBeat() {
	tick := time.NewTicker(3 * time.Second)
	defer tick.Stop()

	for range tick.C {
		this.sendCh <- this.extFns.CS_HeartBeat(this.id)
	}
}

func (this *TcpSession) readloop() {
	defer this.Stop()

	for {
		this.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		if this.extFns.Handler != nil {
			if !this.extFns.Handler(this.id, this.conn, this.Sessmgr.Handler) {
				return
			}
		}
	}
}

func (this *TcpSession) writeloop() {
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

func (this *TcpSession) Stop() {
	this.stop <- true
	this.Sessmgr.RemoveTcpSession(this.id)
}

func (this *TcpSession) SendMsg(data []byte) {
	this.sendCh <- data
}

func (this *TcpSession) Bind(uid int64) {
	this.uid = uid
}

func (this *TcpSession) GetUID() int64 { return this.uid }

func (this *TcpSession) GetType() akmessage.ServerType {
	return this.svrt
}

func (this *TcpSession) SetCliType(t akmessage.ServerType) { this.clit = t }
func (this *TcpSession) GetCliType() akmessage.ServerType  { return this.clit }
