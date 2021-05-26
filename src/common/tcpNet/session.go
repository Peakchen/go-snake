package tcpNet

import (
	"go-snake/akmessage"
	"go-snake/common"
	"go-snake/common/messageBase"
	"go-snake/common/mixNet"
	"net"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Peakchen/xgameCommon/aktime"

	"github.com/Peakchen/xgameCommon/akLog"

	"github.com/Peakchen/xgameCommon/utils"
)

//for client connect server.

type TcpSession struct {
	id       string
	stop     chan<- bool
	stopthis chan bool
	sendCh   chan []byte
	conn     *net.TCPConn
	Sessmgr mixNet.SessionMgrIf
	extFns  *ExtFnsOption
	uid  	int64
	svrt 	akmessage.ServerType
	clit 	akmessage.ServerType
	status 	uint32
	wg 		sync.WaitGroup
}

func NewTcpSession(c *net.TCPConn, st akmessage.ServerType, stop chan<- bool, mgr mixNet.SessionMgrIf, extFn *ExtFnsOption) *TcpSession {
	
	session := &TcpSession{
		id:       strings.Trim(utils.GetUUID(), " "),
		stop:     stop,
		sendCh:   make(chan []byte, 1000),
		conn:     c,
		Sessmgr:  mgr,
		extFns:   extFn,
		svrt:     st,
		stopthis: make(chan bool, 1),
	}

	akLog.Info("tcp new session: ", session.id)
	session.conn.SetKeepAlive(true)
	session.SetConnected()
	mgr.AddTcpSession(session.id, session)
	session.handler()
	
	return session
}

func (this *TcpSession) GetSessionID() string {
	return this.id
}

func (this *TcpSession) handler() {
	
	this.wg.Add(2)

	common.DosafeRoutine(this.readloop, nil)
	common.DosafeRoutine(this.writeloop, nil)

	if this.extFns.CS_HeartBeat != nil {
		common.DosafeRoutine(this.heartBeat, nil)
	}

	this.wg.Wait()
	this.close()

}

func (this *TcpSession) heartBeat() {

	tick := time.NewTicker(3 * time.Second)

	defer func() {
		this.wg.Done()
	}()

heartBeat:
	for {
		select {
		case <-this.stopthis:
			break heartBeat

		case <-tick.C:
			if this.GetStatus() == messageBase.CONNECTED {
				this.sendCh <- this.extFns.CS_HeartBeat(this.id)
			}
		}

	}
	akLog.Info("heartBeat break.")
}

func (this *TcpSession) readloop() {

	common.Dosafe(func() {
		
		defer func() {
			this.stop <- true
			this.wg.Done()
		}()

		for {
			select {
			default:
				this.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
				if this.extFns.Handler != nil {
					if !this.extFns.Handler(this.id, this.conn, this.Sessmgr.Handler) {
						return
					}
				}
			}
		}

		akLog.Info("readloop break.")

	}, nil)
}

func (this *TcpSession) writeloop() {

	common.Dosafe(func() {

		defer func() {
			this.stop <- true
			this.wg.Done()
		}()

		for {
			select {
			case data := <-this.sendCh:

				this.conn.SetWriteDeadline(aktime.Now().Add(15 * time.Second))
				_, err := this.conn.Write(data)
				if err != nil {
					akLog.Info("send data fail, err: ", err)
					return
				}

			}
		}

		akLog.Info("writeloop break.")

	}, nil)

}

func (this *TcpSession) Stop() {

	this.stopthis <- true
	this.stop <- true
	this.SetClose()
	time.Sleep(time.Second)
	this.conn.Close()
}

func (this *TcpSession) close() {
	this.Stop()
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

func (this *TcpSession) SetClose() {
	atomic.StoreUint32(&this.status, messageBase.CLOSED)
}

func (this *TcpSession) SetConnected() {
	atomic.StoreUint32(&this.status, messageBase.CONNECTED)
}

func (this *TcpSession) GetStatus() uint32 {
	return atomic.LoadUint32(&this.status)
}
