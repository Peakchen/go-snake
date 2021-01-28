// add by stefan

package akNet

import (
	"fmt"
	"net"
	"sync/atomic"
	"time"

	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/aktime"
	"github.com/Peakchen/xgameCommon/define"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_HeartBeat"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_MainModule"

	//"S2SMessage"
	"context"
	"sync"

	"github.com/Peakchen/xgameCommon/stacktrace"
	"github.com/golang/protobuf/proto"
	//. "define"
)

type SvrTcpSession struct {
	sync.Mutex

	RemoteAddr string
	isAlive    bool
	// The net connection.
	conn *net.TCPConn
	// Buffered channel of outbound messages.
	send   chan []byte
	readCh chan bool
	// send/recv
	sw  sync.WaitGroup
	ctx context.Context
	// person offline flag
	off chan *SvrTcpSession
	//message pack
	pack IMessagePack
	// session id
	SessionID uint64
	//Dest point
	SvrType define.ERouteId
	//src point
	RegPoint define.ERouteId
	//person StrIdentify
	StrIdentify       string
	heartBeatDeadline int64
	Name              string
}

func NewSvrSession(addr string,
	conn *net.TCPConn,
	ctx context.Context,
	SvrType define.ERouteId,
	off chan *SvrTcpSession,
	pack IMessagePack,
	procName string) *SvrTcpSession {
	return &SvrTcpSession{
		RemoteAddr: addr,
		conn:       conn,
		send:       make(chan []byte, maxMessageSize),
		readCh:     make(chan bool, maxMessageSize),
		isAlive:    false,
		ctx:        ctx,
		pack:       pack,
		off:        off,
		SvrType:    SvrType,
		Name:       procName,
	}
}

func (this *SvrTcpSession) Alive() bool {
	return this.isAlive
}

func (this *SvrTcpSession) close(sw *sync.WaitGroup) {
	if this == nil {
		return
	}

	akLog.FmtPrintf("session close, svr: %v, regpoint: %v, cache size: %v.", this.SvrType, this.RegPoint, len(this.send))
	select {
	case this.off <- this:
	}

	GServer2ServerSession.RemoveSession(this.RemoteAddr)
	//close(this.send)
	this.conn.CloseRead()
	this.conn.CloseWrite()
	this.conn.Close()
}

func (this *SvrTcpSession) SetSendCache(data []byte) {
	this.send <- data
}

func (this *SvrTcpSession) sendloop(sw *sync.WaitGroup) {
	defer func() {
		sw.Done()
		this.close(sw)
	}()

	for {
		select {
		case <-this.ctx.Done():
			return
		case data := <-this.send:
			if !this.WriteMessage(data) {
				return
			}
		}
	}
}

func (this *SvrTcpSession) recvloop(sw *sync.WaitGroup) {
	defer func() {
		sw.Done()
		this.close(sw)
	}()

	for {
		select {
		case <-this.ctx.Done():
			return
		case brecvClient := <-this.readCh:
			this.Invoke(brecvClient)
		default:
			if !this.readMessage() {
				return
			}
		}
	}
}

func (this *SvrTcpSession) heartBeatloop(sw *sync.WaitGroup) {
	defer func() {
		sw.Done()
	}()

	ticker := time.NewTicker(time.Duration(cstCheckHeartBeatMonitorSec) * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-this.ctx.Done():
			return
		case <-ticker.C:
			if this.heartBeatDeadline == 0 {
				continue
			}

			var disconnectionSec int
			if this.RegPoint == 0 {
				disconnectionSec = cstClientDisconnectionSec
			} else {
				disconnectionSec = cstSvrDisconnectionSec
			}

			if aktime.Now().Unix()-this.heartBeatDeadline >= int64(disconnectionSec) {
				//close connection...
				this.close(sw)
				this.heartBeatDeadline = 0
			}
		}
	}
}

func (this *SvrTcpSession) WriteMessage(data []byte) (succ bool) {
	if !this.isAlive || len(data) == 0 {
		return
	}

	defer stacktrace.Catchcrash()

	this.conn.SetWriteDeadline(aktime.Now().Add(writeWait))
	//send...
	//akLog.FmtPrintln("[server] begin send response message to client, message length: ", len(data))
	_, err := this.conn.Write(data)
	if err != nil {
		akLog.FmtPrintln("send data fail, err: ", err)
		return false
	}

	return true
}

func (this *SvrTcpSession) readMessage() (succ bool) {
	defer func() {
		this.Unlock()
		stacktrace.Catchcrash()
	}()

	this.Lock()
	//this.conn.SetReadDeadline(aktime.Now().Add(pongWait))
	var responseCliented bool
	if this.RegPoint == 0 {
		succ = UnPackExternalMsg(this.conn, this.pack)
		if !succ {
			return
		}
		this.pack.SetRemoteAddr(this.RemoteAddr)
	} else {
		succ = UnPackInnerMsg(this.conn, this.pack)
		if !succ {
			return
		}
		this.StrIdentify = this.pack.GetIdentify()
		if this.SvrType == define.ERouteId_ER_ESG {
			responseCliented = true
		}
	}
	this.readCh <- responseCliented
	return
}

func (this *SvrTcpSession) Invoke(responseCliented bool) (succ bool) {
	var route define.ERouteId
	mainID, SubID := this.pack.GetMessageID()
	akLog.FmtPrintf("recv message: mainID: %v, subID: %v.", mainID, SubID)
	if mainID == uint16(MSG_MainModule.MAINMSG_SERVER) &&
		this.SvrType == define.ERouteId_ER_ESG {
		route = define.ERouteId_ER_ISG
		this.RegPoint = define.ERouteId_ER_ISG
		this.Push(define.ERouteId_ER_ISG) //外网关加入内网关session
		RegisterMessageRet(this)
		succ = true
		return
	}

	if (mainID == uint16(MSG_MainModule.MAINMSG_SERVER) ||
		mainID == uint16(MSG_MainModule.MAINMSG_LOGIN)) && len(this.StrIdentify) == 0 {
		this.StrIdentify = this.RemoteAddr
	}

	if len(this.pack.GetIdentify()) == 0 {
		this.pack.SetIdentify(this.StrIdentify)
	}

	if mainID == uint16(MSG_MainModule.MAINMSG_LOGIN) {
		route = define.ERouteId_ER_Login
	} else if mainID >= uint16(MSG_MainModule.MAINMSG_PLAYER) {
		route = define.ERouteId_ER_Game
	}

	if mainID != uint16(MSG_MainModule.MAINMSG_SERVER) && mainID != uint16(MSG_MainModule.MAINMSG_HEARTBEAT) &&
		(this.SvrType == define.ERouteId_ER_ESG || this.SvrType == define.ERouteId_ER_ISG) {
		//akLog.FmtPrintf("[server] Route (%v), StrIdentify: %v.", route, this.StrIdentify)
		if this.SvrType == define.ERouteId_ER_ESG {
			succ = externalRouteAct(route, this, responseCliented)
		} else {
			succ = innerMsgRouteAct(ESessionType_Server, route, mainID, this.pack.GetSrcMsg())
		}
	} else {
		succ = msgCallBack(this) //路由消息回调处理
	}
	return
}

func (this *SvrTcpSession) GetPack() (obj IMessagePack) {
	return this.pack
}

func (this *SvrTcpSession) HandleSession(sw *sync.WaitGroup) {
	this.isAlive = true
	atomic.AddUint64(&this.SessionID, 1)
	sw.Add(3)
	go this.recvloop(sw)
	go this.sendloop(sw)
	go this.heartBeatloop(sw)

	this.Name = fmt.Sprintf("server_%v", this.Name)
}

func (this *SvrTcpSession) Push(RegPoint define.ERouteId) {
	//akLog.FmtPrintf("[server] push new sesson, reg point: %v.", RegPoint)
	this.RegPoint = RegPoint
	GServer2ServerSession.AddSession(this.RemoteAddr, this)
}

func (this *SvrTcpSession) SetIdentify(StrIdentify string) {
	session := GServer2ServerSession.GetSessionByIdentify(this.StrIdentify)
	if session != nil {
		GServer2ServerSession.RemoveSession(this.StrIdentify)
		this.StrIdentify = StrIdentify
		GServer2ServerSession.AddSession(StrIdentify, session)
	}
}

func (this *SvrTcpSession) Offline() {

}

func (this *SvrTcpSession) SendSvrClientMsg(mainid, subid uint16, msg proto.Message) (succ bool, err error) {
	if !this.isAlive {
		err = fmt.Errorf("[server] send msg session disconnection, mainid: %v, subid: %v.", mainid, subid)
		akLog.FmtPrintln("send msg err: ", err)
		return false, err
	}

	data, err := this.pack.PackClientMsg(mainid, subid, msg)
	if err != nil {
		return succ, err
	}
	this.SetSendCache(data)
	return true, nil
}

func (this *SvrTcpSession) SendInnerSvrMsg(mainid, subid uint16, msg proto.Message) (succ bool, err error) {
	if !this.isAlive {
		err = fmt.Errorf("[server] send svr session disconnection, mainid: %v, subid: %v.", mainid, subid)
		akLog.FmtPrintln("send msg err: ", err)
		return false, err
	}

	data, err := this.pack.PackInnerMsg(mainid, subid, msg)
	if err != nil {
		return succ, err
	}
	this.SetSendCache(data)
	return true, nil
}

func (this *SvrTcpSession) SendInnerClientMsg(mainid, subid uint16, msg proto.Message) (succ bool, err error) {
	if !this.isAlive {
		err = fmt.Errorf("[server] session disconnection, mainid: %v, subid: %v.", mainid, subid)
		akLog.FmtPrintln("send msg err: ", err)
		return false, err
	}

	if len(this.GetIdentify()) > 0 {
		this.pack.SetIdentify(this.GetIdentify())
	}

	this.pack.SetPostType(MsgPostType_Single)

	data, err := this.pack.PackInnerMsg(mainid, subid, msg)
	if err != nil {
		return succ, err
	}
	this.SetSendCache(data)
	return true, nil
}

func (this *SvrTcpSession) SendInnerBroadcastMsg(mainid, subid uint16, msg proto.Message) (succ bool, err error) {
	if !this.isAlive {
		err = fmt.Errorf("[server] session disconnection, mainid: %v, subid: %v.", mainid, subid)
		akLog.FmtPrintln("send msg err: ", err)
		return false, err
	}

	if len(this.GetIdentify()) > 0 {
		this.pack.SetIdentify(this.GetIdentify())
	}

	this.pack.SetPostType(MsgPostType_Broadcast)

	data, err := this.pack.PackInnerMsg(mainid, subid, msg)
	if err != nil {
		return succ, err
	}
	this.SetSendCache(data)
	return true, nil
}

func (this *SvrTcpSession) GetIdentify() string {
	return this.StrIdentify
}

func (this *SvrTcpSession) GetRegPoint() (RegPoint define.ERouteId) {
	return this.RegPoint
}

func (this *SvrTcpSession) GetRemoteAddr() string {
	return this.RemoteAddr
}

func (this *SvrTcpSession) IsUser() bool {
	return this.RegPoint == 0
}

func (this *SvrTcpSession) RefreshHeartBeat(mainid, subid uint16) bool {
	if mainid == uint16(MSG_MainModule.MAINMSG_HEARTBEAT) &&
		subid == uint16(MSG_HeartBeat.SUBMSG_CS_HeartBeat) {
		this.heartBeatDeadline = aktime.Now().Unix()
	}
	return true
}

func (this *SvrTcpSession) GetModuleName() string {
	return this.Name
}
