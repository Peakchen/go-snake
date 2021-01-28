package Kcpnet

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/akNet"
	"github.com/Peakchen/xgameCommon/aktime"
	"github.com/Peakchen/xgameCommon/define"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_HeartBeat"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_MainModule"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_Server"
	"github.com/Peakchen/xgameCommon/stacktrace"
	"github.com/golang/protobuf/proto"
)

type KcpClientSession struct {
	conn         net.Conn
	readCh       chan []byte
	writeCh      chan []byte
	remoteAddr   string
	pack         IMessagePack
	isAlive      bool
	offCh        chan *KcpClientSession
	closeOnce    sync.Once
	SvrType      define.ERouteId
	RegPoint     define.ERouteId
	StrIdentify  string
	Name         string
	exCollection *ExternalCollection //for external expand data
}

func NewKcpClientSession(c net.Conn, offCh chan *KcpClientSession, exCol *ExternalCollection) *KcpClientSession {
	return &KcpClientSession{
		conn:         c,
		readCh:       make(chan []byte, 1000),
		writeCh:      make(chan []byte, 1000),
		remoteAddr:   c.RemoteAddr().String(),
		pack:         &KcpClientProtocol{},
		offCh:        offCh,
		isAlive:      true,
		exCollection: exCol,
	}
}

func (this *KcpClientSession) Alive() bool {
	return this.isAlive
}

func (this *KcpClientSession) close() {
	if this == nil {
		return
	}

	this.closeOnce.Do(func() {
		leftplayers := GPlayerStaticis.SubPlayer(this.SvrType)
		akLog.FmtPrintf("server session close, svr: %v, regpoint: %v, left players: %v.", this.SvrType, this.RegPoint, leftplayers)
		GClient2ServerSession.RemoveSession(this.remoteAddr)
		this.isAlive = false
		this.offCh <- this

		this.conn.Close()
	})

}

func (this *KcpClientSession) SetSendCache(data []byte) {
	this.writeCh <- data
}

func (this *KcpClientSession) heartbeatloop() {
	defer func() {
		this.close()
	}()

	ticker := time.NewTicker(time.Duration(cstKeepLiveHeartBeatSec) * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if this.RegPoint == 0 || len(this.StrIdentify) == 0 {
				continue
			}
			sendHeartBeat(this)
		}
	}
}

func (this *KcpClientSession) Handler() {
	go this.readloop()
	go this.writeloop()
}

func (this *KcpClientSession) readloop() {
	defer func() {
		this.close()
	}()

	for {
		select {
		case <-this.readCh:
			this.dispatch()
		default:
			this.conn.SetReadDeadline(aktime.Now().Add(pongWait))
			//是否加个消息队列处理 ?
			this.read()
		}
	}
}

func (this *KcpClientSession) read() (succ bool) {
	defer func() {
		stacktrace.Catchcrash()
	}()

	if len(this.StrIdentify) == 0 &&
		(this.SvrType == define.ERouteId_ER_ESG || this.SvrType == define.ERouteId_ER_ISG) {
		succ = UnPackExternalMsg(this.conn, this.pack)
		if !succ {
			return
		}
		this.pack.SetRemoteAddr(this.remoteAddr)
	} else {
		succ = UnPackInnerMsg(this.conn, this.pack)
		if !succ {
			return
		}
		this.StrIdentify = this.pack.GetIdentify()
	}

	this.readCh <- []byte{1}
	return
}

func (this *KcpClientSession) dispatch() (succ bool) {
	defer func() {
		stacktrace.Catchcrash()
	}()

	var route define.ERouteId
	mainID, _ := this.pack.GetMessageID()
	if (mainID == uint16(MSG_MainModule.MAINMSG_SERVER) ||
		mainID == uint16(MSG_MainModule.MAINMSG_LOGIN)) && len(this.StrIdentify) == 0 {
		this.StrIdentify = this.remoteAddr
	}
	if len(this.pack.GetIdentify()) == 0 {
		this.pack.SetIdentify(this.StrIdentify)
	}
	if mainID == uint16(MSG_MainModule.MAINMSG_LOGIN) {
		route = define.ERouteId_ER_Login
	} else if mainID >= uint16(MSG_MainModule.MAINMSG_PLAYER) {
		route = define.ERouteId_ER_Game
	}
	if mainID != uint16(MSG_MainModule.MAINMSG_SERVER) &&
		mainID != uint16(MSG_MainModule.MAINMSG_HEARTBEAT) &&
		(this.SvrType == define.ERouteId_ER_ISG) {
		//akLog.FmtPrintf("[client] StrIdentify: %v.", this.StrIdentify)
		succ = innerMsgRouteAct(akNet.ESessionType_Client, route, mainID, this.pack.GetSrcMsg())
	} else {
		succ = this.checkmsgProc(route) //路由消息回调处理
	}
	return
}

func (this *KcpClientSession) writeloop() {
	defer func() {
		this.close()
	}()

	for {
		select {
		case data := <-this.writeCh:
			this.conn.SetReadDeadline(aktime.Now().Add(pongWait))
			this.WriteMessage(data)
		}
	}
}

func (this *KcpClientSession) WriteMessage(data []byte) (succ bool) {
	n, err := this.conn.Write(data)
	if err != nil {
		akLog.Error("send reply data fail, size: %v, err: %v.", n, err)
		return
	}
	succ = true
	return
}

func (this *KcpClientSession) checkRegisterRet(route define.ERouteId) (exist bool) {
	mainID, subID := this.pack.GetMessageID()
	if mainID == uint16(MSG_MainModule.MAINMSG_SERVER) &&
		uint16(MSG_Server.SUBMSG_SC_ServerRegister) == subID {
		this.StrIdentify = this.remoteAddr
		if this.SvrType == define.ERouteId_ER_ISG {
			this.Push(define.ERouteId_ER_ESG)
		} else if this.SvrType == define.ERouteId_ER_ESG {
			this.Push(define.ERouteId_ER_CenterGate)
		} else {
			this.Push(define.ERouteId_ER_ISG)
		}

		exist = true
	}
	return
}

func (this *KcpClientSession) checkHeartBeatRet() (exist bool) {
	mainID, subID := this.pack.GetMessageID()
	if mainID == uint16(MSG_MainModule.MAINMSG_HEARTBEAT) &&
		uint16(MSG_HeartBeat.SUBMSG_SC_HeartBeat) == subID {
		exist = true
	}
	return
}

func (this *KcpClientSession) checkmsgProc(route define.ERouteId) (succ bool) {
	//akLog.FmtPrintf("recv response, route: %v.", route)
	bRegister := this.checkRegisterRet(route)
	bHeartBeat := checkHeartBeatRet(this.pack)
	if bRegister || bHeartBeat {
		succ = true
		return
	}

	succ = msgCallBack(this)
	return
}

func (this *KcpClientSession) GetPack() (obj IMessagePack) {
	return this.pack
}

func (this *KcpClientSession) Push(RegPoint define.ERouteId) {
	//akLog.FmtPrintf("[client] push new sesson, reg point: %v.", RegPoint)
	this.RegPoint = RegPoint
	GServer2ServerSession.AddSession(this.remoteAddr, this)
}

func (this *KcpClientSession) SetIdentify(StrIdentify string) {
	session := GServer2ServerSession.GetSessionByIdentify(this.StrIdentify)
	if session != nil {
		GServer2ServerSession.RemoveSession(this.StrIdentify)
		this.StrIdentify = StrIdentify
		GServer2ServerSession.AddSession(StrIdentify, session)
	}
}

func (this *KcpClientSession) Offline() {
	// notify some one server...
}

func (this *KcpClientSession) SendSvrClientMsg(mainid, subid uint16, msg proto.Message) (succ bool, err error) {
	if !this.isAlive {
		err = fmt.Errorf("[client] writeCh msg session disconnection, mainid: %v, subid: %v.", mainid, subid)
		akLog.FmtPrintln("writeCh msg err: ", err)
		return succ, err
	}

	data, err := this.pack.PackClientMsg(mainid, subid, msg)
	if err != nil {
		return succ, err
	}
	this.SetSendCache(data)
	return true, nil
}

func (this *KcpClientSession) SendInnerSvrMsg(mainid, subid uint16, msg proto.Message) (succ bool, err error) {
	if !this.isAlive {
		err = fmt.Errorf("[client] writeCh svr session disconnection, mainid: %v, subid: %v.", mainid, subid)
		akLog.FmtPrintln("writeCh msg err: ", err)
		return false, err
	}

	data, err := this.pack.PackInnerMsg(mainid, subid, msg)
	if err != nil {
		return succ, err
	}
	this.SetSendCache(data)
	return true, nil
}

func (this *KcpClientSession) SendInnerClientMsg(mainid, subid uint16, msg proto.Message) (succ bool, err error) {
	if !this.isAlive {
		err = fmt.Errorf("[client] session disconnection, mainid: %v, subid: %v.", mainid, subid)
		akLog.FmtPrintln("writeCh msg err: ", err)
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

func (this *KcpClientSession) SendInnerBroadcastMsg(mainid, subid uint16, msg proto.Message) (succ bool, err error) {
	if !this.isAlive {
		err = fmt.Errorf("[client] session disconnection, mainid: %v, subid: %v.", mainid, subid)
		akLog.FmtPrintln("writeCh msg err: ", err)
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

func (this *KcpClientSession) GetIdentify() string {
	return this.StrIdentify
}

func (this *KcpClientSession) GetRegPoint() (RegPoint define.ERouteId) {
	return this.RegPoint
}

func (this *KcpClientSession) GetRemoteAddr() string {
	return this.remoteAddr
}

func (this *KcpClientSession) IsUser() bool {
	return this.RegPoint == 0
}

func (this *KcpClientSession) RefreshHeartBeat(mainid, subid uint16) bool {
	return true
}

func (this *KcpClientSession) GetModuleName() string {
	return this.Name
}

func (this *KcpClientSession) GetExternalCollection() *ExternalCollection {
	return this.exCollection
}

func (this *KcpClientSession) GetVer() int32 {
	return 0
}

func (this *KcpClientSession) SetVer(ver int32) {

}

func (this *KcpClientSession) GetSvrType() (t define.ERouteId) {
	return this.SvrType
}
