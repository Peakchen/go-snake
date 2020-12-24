package app

import (
	"fmt"
	"go-snake/akmessage"
	"go-snake/common/messageBase"
	"go-snake/common/mixNet"
	"go-snake/common/tcpNet"
	"go-snake/common/webNet"
	"reflect"
	"sync"

	"google.golang.org/protobuf/proto"

	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/gorilla/websocket"
)

func Init() {
	mixNet.SetApp(New())
}

type S2SContext struct {
	roles    uint32
	session  *tcpNet.TcpSession
	msgActor IMessageActor
}

type C2SContext struct {
	SID      string
	session  *webNet.WebSession
	msgActor IMessageActor
}

type GateApp struct {
	c2sMt sync.RWMutex
	s2sMt sync.RWMutex

	roles uint32
	c2s   map[string]*C2SContext
	s2s   map[string]*S2SContext
	o2i   map[string]string

	isclose bool
}

func New() *GateApp {
	return &GateApp{
		c2s: make(map[string]*C2SContext),
		s2s: make(map[string]*S2SContext),
		o2i: make(map[string]string),
	}
}

// rule 1: get max roles server and role not big equal 5w
func (this *GateApp) selectServer(routeID akmessage.ServerType) string {
	var dst string
	for id, c := range this.s2s {
		if c.roles < 50000 && c.session.GetCliType() == routeID {
			dst = id
		}
	}
	return dst
}

func (this *GateApp) Online(nt messageBase.NetType, sess interface{}) {
	switch nt {
	case messageBase.Net_WS:
		this.c2sMt.Lock()
		defer this.c2sMt.Unlock()

		this.roles++

		s := sess.(*webNet.WebSession)
		this.c2s[s.GetSessionID()] = &C2SContext{
			session:  sess.(*webNet.WebSession),
			msgActor: &GateMessage{},
		}
	case messageBase.Net_TCP:
		this.s2sMt.Lock()
		defer this.s2sMt.Unlock()

		s := sess.(*tcpNet.TcpSession)
		this.s2s[s.GetSessionID()] = &S2SContext{
			session:  s,
			msgActor: &GateMessage{},
		}
	}
}

func (this *GateApp) Offline(nt messageBase.NetType, id string) {
	switch nt {
	case messageBase.Net_WS:
		this.c2sMt.Lock()
		defer this.c2sMt.Unlock()

		this.roles--
		delete(this.c2s, id)
	case messageBase.Net_TCP:
		this.s2sMt.Lock()
		defer this.s2sMt.Unlock()

		delete(this.s2s, id)
	}
}

func (this *GateApp) Bind(sid string, id int64) {
	//this.c2sMt.Lock()
	//defer this.c2sMt.Unlock()

	sessContent, ok := this.c2s[sid]
	if !ok {
		akLog.Error("can not find client session, id: ", sid)
		return
	}
	if id == 0 {
		akLog.Error("invalid uid: ", id)
		return
	}
	sessContent.session.Bind(id)
}

//c->[gate1<=>gate2]->s
func (this *GateApp) CS_SendInner(sid string, id uint32, data []byte) {
	//this.c2sMt.Lock()
	//defer this.c2sMt.Unlock()

	c, ok := this.c2s[sid]
	if !ok {
		akLog.Error("can not find client session, sid: ", sid)
		return
	}
	c.msgActor.SetSession(sid)

	cspt := messageBase.CSPackTool()
	err := cspt.UnPack(data)
	if err != nil {
		akLog.Error(err)
		return
	}

	defer cspt.Reset()

	routeID := akmessage.GetMsgRoute(akmessage.MSG(cspt.GetMsgID()))
	switch routeID {
	case akmessage.ServerType_Gate:
		content := GetActorMessageProc(cspt.GetMsgID())
		if content != nil {
			dst := reflect.New(content.refPb.Elem()).Interface().(proto.Message)
			err := messageBase.Codec().Unmarshal(cspt.GetData(), dst)
			if err != nil {
				akLog.Error(fmt.Errorf("unmarshal message fail, err: %v.", err))
				return
			}
			var refs []reflect.Value
			refs = append(refs, reflect.ValueOf(c.msgActor))
			refs = append(refs, reflect.ValueOf(dst))
			content.refProc.Call(refs)
		}
		return
	default:

	}
	selectSID := this.selectServer(routeID)
	if len(selectSID) == 0 {
		akLog.Error("can select server, info: ", cspt.GetMsgID(), routeID)
		return
	}

	//this.s2sMt.RLock()
	//defer this.s2sMt.RUnlock()

	s, ok := this.s2s[selectSID]
	if ok {
		c.SID = selectSID
	} else {
		akLog.Error("can not find server select server, selectSID: ", selectSID)
		return
	}
	packMsg := messageBase.SSPackMsg_pb(sid, s.session.GetUID(), akmessage.MSG_SS_ROUTE, &akmessage.SS_SSRoute{Data: data})
	if packMsg == nil {
		return
	}
	s.session.SendMsg(packMsg)
}

//s->[gate2<=>gate1]->c
//s-> gate2 rid
func (this *GateApp) SendClient(sid string, id uint32, data []byte) {
	//this.c2sMt.RLock()
	//defer this.c2sMt.RUnlock()

	sessContent, ok := this.c2s[sid]
	if !ok {
		akLog.Error("can not find client session, id: ", sid)
		return
	}
	cspt := messageBase.CSPackTool()
	cspt.Init(id, data)
	out := make([]byte, len(data)+messageBase.CS_MSG_PACK_NODATA_SIZE)
	cspt.Pack(out)
	sessContent.session.Write(websocket.BinaryMessage, out)
}

//game/login/...->gate->client
func (this *GateApp) Handler(sid string, data []byte) {
	sspt := messageBase.SSPackTool()
	err := sspt.UnPack(data)
	if err != nil {
		akLog.Error("ssPack msg fail.")
		return
	}

	//服务器注册
	switch sspt.GetMsgID() {
	case uint32(akmessage.MSG_SS_REGISTER_REQ):
		func(data []byte, sessid string) {
			reg := &akmessage.SS_Register_Req{}
			err := messageBase.Codec().Unmarshal(data, reg)
			if err != nil {
				akLog.Error("proto unmarshal msg fail.")
				return
			}

			//this.s2sMt.Lock()
			//defer this.s2sMt.Unlock()

			s, ok := this.s2s[sessid]
			if !ok {
				akLog.Error("s2s can not find session, sid: ", sessid)
				return
			}
			akLog.FmtPrintln("register client dst:", reg.St)
			s.session.SetCliType(reg.St)
			s.msgActor.SetSession(sessid)
			rsp := messageBase.GetActorRegisterRsp(sessid, akmessage.ServerType_Gate)
			s.session.SendMsg(rsp)
		}(sspt.GetData(), sid)
		return
	default:

	}

	//网关自身消息回调处理
	routeID := akmessage.GetMsgRoute(akmessage.MSG(sspt.GetMsgID()))
	switch routeID {
	case akmessage.ServerType_Gate:
		func(id uint32, data []byte, sid string) {
			content := GetActorMessageProc(id)
			if content != nil {
				dst := reflect.New(content.refPb.Elem()).Interface().(proto.Message)
				err := messageBase.Codec().Unmarshal(data, dst)
				if err != nil {
					akLog.Error(fmt.Errorf("unmarshal message fail, err: %v.", err))
					return
				}
				s, ok := this.s2s[sid]
				if !ok {
					akLog.Error("s2s can not find session, sid: ", sid)
					return
				}
				s.msgActor.SetSession(sid)
				var refs []reflect.Value
				refs = append(refs, reflect.ValueOf(s.msgActor))
				refs = append(refs, reflect.ValueOf(dst))
				content.refProc.Call(refs)
			}
		}(sspt.GetMsgID(), sspt.GetData(), sid)
		return
	default:

	}
	//客户端消息转发
	switch sspt.GetMsgID() {
	case uint32(akmessage.MSG_SC_LOGIN):
		this.Bind(sspt.GetSessID(), sspt.GetUID())
	}
	this.SendClient(sspt.GetSessID(), sspt.GetMsgID(), sspt.GetData())
}

func (this *GateApp) SS_SendInner(sid string, id uint32, data []byte) {
	//this.s2sMt.RLock()
	//defer this.s2sMt.RUnlock()

	sessContent, ok := this.s2s[sid]
	if !ok {
		akLog.Error("can not find client session, id: ", sid)
		return
	}
	sessContent.session.SendMsg(data)
}

func (this *GateApp) Close() {

	this.isclose = true

	akLog.Info("begin -> c2s：%v, s2s: %v.", len(this.c2s), len(this.s2s))
	this.c2sMt.Lock()
	for _, c := range this.c2s {
		c.session.Stop()
	}
	this.c2s = map[string]*C2SContext{}
	this.c2sMt.Unlock()

	this.s2sMt.Lock()
	for _, c := range this.s2s {
		c.session.Stop()
	}
	this.s2s = map[string]*S2SContext{}
	this.s2sMt.Unlock()

	akLog.Info("after -> c2s：%v, s2s: %v.", len(this.c2s), len(this.s2s))
}

func (this *GateApp) IsClose() bool {
	return this.isclose
}
