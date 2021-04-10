package Kcpnet

import (
	//"aktime"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/akNet"
	"github.com/Peakchen/xgameCommon/aktime"
	"github.com/Peakchen/xgameCommon/define"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_MainModule"
	"github.com/Peakchen/xgameCommon/stacktrace"
	"github.com/golang/protobuf/proto"
)

type KcpServerSession struct {
	conn              net.Conn
	readCh            chan bool
	writeCh           chan []byte
	RemoteAddr        string
	pack              IMessagePack
	offCh             chan *KcpServerSession
	isAlive           bool
	closeOnce         sync.Once
	SvrType           define.ERouteId
	RegPoint          define.ERouteId
	heartBeatDeadline int64
	kcpConfig         *KcpSvrConfig
	StrIdentify       string
	Name              string
	exCollection      *ExternalCollection //for external expand data
	versionNo         int32               //
}

func NewKcpSvrSession(c net.Conn, offCh chan *KcpServerSession, kcpcfg *KcpSvrConfig, svrType define.ERouteId, exCol *ExternalCollection) *KcpServerSession {
	return &KcpServerSession{
		conn:         c,
		readCh:       make(chan bool, 1000),
		writeCh:      make(chan []byte, 1000),
		RemoteAddr:   c.RemoteAddr().String(),
		pack:         &KcpServerProtocol{},
		offCh:        offCh,
		kcpConfig:    kcpcfg,
		isAlive:      true,
		exCollection: exCol,
		SvrType:      svrType,
	}
}

func (this *KcpServerSession) Handler() {
	go this.readloop()
	go this.writeloop()
}

func (this *KcpServerSession) close() {
	this.closeOnce.Do(func() {
		leftplayers := GPlayerStaticis.SubPlayer(this.SvrType)
		akLog.FmtPrintf("server session close, svr: %v, regpoint: %v, left players: %v.", this.SvrType, this.RegPoint, leftplayers)
		this.isAlive = false
		this.offCh <- this
		this.conn.Close()
	})
}

func (this *KcpServerSession) heartBeatloop(sw *sync.WaitGroup) {

	ticker := time.NewTicker(time.Duration(cstCheckHeartBeatMonitorSec) * time.Second)

	defer func() {
		ticker.Stop()
		sw.Done()
	}()

	for {
		select {
		// case <-this.ctx.Done():
		// 	return
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
				this.close()
				this.heartBeatDeadline = 0
			}
		}
	}
}

func (this *KcpServerSession) readloop() {

	defer func() {
		this.close()
	}()

	for {
		select {
		case rspcliented := <-this.readCh:
			this.dispatch(rspcliented)
		default:
			this.conn.SetReadDeadline(aktime.Now().Add(this.kcpConfig.udp_readDeadline))
			//是否加个消息队列处理 ?
			this.read()
		}
	}
}

func (this *KcpServerSession) read() (succ bool) {

	defer func() {
		stacktrace.Catchcrash()
	}()

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

func (this *KcpServerSession) dispatch(responseCliented bool) (succ bool) {
	defer func() {
		stacktrace.Catchcrash()
	}()

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
		if this.SvrType == define.ERouteId_ER_ESG {
			succ = externalRouteAct(route, this, responseCliented, this.exCollection)
			if succ && this.exCollection != nil {
				this.exCollection.SetExternalClient(GClient2ServerSession)
			}
		} else {
			GPlayerStaticis.AddPlayer(this.RemoteAddr)
			succ = innerMsgRouteAct(akNet.ESessionType_Server, route, mainID, this.pack.GetSrcMsg())
		}
	} else {
		succ = msgCallBack(this) //路由消息回调处理
	}
	return
}

func (this *KcpServerSession) writeloop() {

	defer func() {
		this.close()
	}()

	for {
		select {
		case data := <-this.writeCh:
			this.conn.SetWriteDeadline(aktime.Now().Add(this.kcpConfig.udp_writeDeadline))
			this.WriteMessage(data)
		}
	}
}

func (this *KcpServerSession) WriteMessage(data []byte) (succ bool) {
	n, err := this.conn.Write(data)
	if err != nil {
		akLog.Error("send reply data fail, size: %v, err: %v.", n, err)
		return
	}
	succ = true
	return
}

func (this *KcpServerSession) Alive() bool {
	return this.isAlive
}

func (this *KcpServerSession) SetSendCache(data []byte) {
	this.writeCh <- data
}

func (this *KcpServerSession) GetPack() (obj IMessagePack) {
	return this.pack
}

func (this *KcpServerSession) GetRemoteAddr() string {
	return this.RemoteAddr
}

func (this *KcpServerSession) Push(RegPoint define.ERouteId) {
	this.RegPoint = RegPoint
	GServer2ServerSession.AddSession(this.RemoteAddr, this)
}

func (this *KcpServerSession) SetIdentify(StrIdentify string) {
	session := GServer2ServerSession.GetSessionByIdentify(this.StrIdentify)
	if session != nil {
		GServer2ServerSession.RemoveSession(this.StrIdentify)
		this.StrIdentify = StrIdentify
		GServer2ServerSession.AddSession(StrIdentify, session)
	}
}

func (this *KcpServerSession) Offline() {
	if this.exCollection != nil && this.SvrType == define.ERouteId_ER_ESG {
		if this.exCollection.GetCenterClient() != nil {
			sendCenterSvr4Leave(this, this.exCollection)
		}
	}
}

func (this *KcpServerSession) SendSvrClientMsg(mainid, subid uint16, msg proto.Message) (succ bool, err error) {
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

func (this *KcpServerSession) SendInnerSvrMsg(mainid, subid uint16, msg proto.Message) (succ bool, err error) {
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

func (this *KcpServerSession) SendInnerClientMsg(mainid, subid uint16, msg proto.Message) (succ bool, err error) {
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

func (this *KcpServerSession) SendInnerBroadcastMsg(mainid, subid uint16, msg proto.Message) (succ bool, err error) {
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

func (this *KcpServerSession) GetIdentify() string {
	return this.StrIdentify
}

func (this *KcpServerSession) GetRegPoint() (RegPoint define.ERouteId) {
	return this.RegPoint
}

func (this *KcpServerSession) GetModuleName() string {
	return this.Name
}

func (this *KcpServerSession) IsUser() bool {
	return this.RegPoint == 0
}

func (this *KcpServerSession) RefreshHeartBeat(mainid, subid uint16) bool {
	return true
}

func (this *KcpServerSession) GetExternalCollection() *ExternalCollection {
	return this.exCollection
}

func (this *KcpServerSession) GetVer() int32 {
	return this.versionNo
}

func (this *KcpServerSession) SetVer(ver int32) {
	this.versionNo = ver
}

func (this *KcpServerSession) GetSvrType() (t define.ERouteId) {
	return this.SvrType
}
