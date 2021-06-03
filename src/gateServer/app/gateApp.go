package app

import (
	"fmt"
	"go-snake/akmessage"
	"go-snake/common"
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
	sync.RWMutex

	nt 		messageBase.NetType
	roles 	uint32
	c2s   	sync.Map
	s2s   	sync.Map

	isclose bool
}

func New() *GateApp {
	return &GateApp{
	}
}

// rule 1: get max roles server and role not big equal 5w
func (this *GateApp) selectServer(routeID akmessage.ServerType) string {
	var dst string
	this.s2s.Range(func(key interface{}, value interface{}) bool {
		c := value.(*S2SContext)
		if c.roles < 50000 && c.session.GetCliType() == routeID {
			dst = key.(string)
			return false
		}
		return true
	})
	return dst
}

func (this *GateApp) Online(nt messageBase.NetType, sess interface{}) {

	this.nt = nt

	common.Dosafe(func() {

		switch nt {
		case messageBase.Net_WS:

			this.Lock()
			this.roles++
			this.Unlock()

			s := sess.(*webNet.WebSession)
			this.c2s.Store(s.GetSessionID(), &C2SContext{
				session:  sess.(*webNet.WebSession),
				msgActor: &GateMessage{},
			})

		case messageBase.Net_TCP:

			s := sess.(*tcpNet.TcpSession)
			this.s2s.Store(s.GetSessionID(), &S2SContext{
				session:  s,
				msgActor: &GateMessage{},
			})

		}

	}, nil)
}

func (this *GateApp) Offline(nt messageBase.NetType, id string) {
	common.Dosafe(func() {
		switch nt {
		case messageBase.Net_WS:

			this.Lock()
			this.roles--
			this.Unlock()

			this.c2s.Delete(id)
			akLog.Info("ws offline: ", id)

		case messageBase.Net_TCP:

			this.s2s.Delete(id)
			akLog.Info("tcp offline: ", id)
		}
	}, nil)
}

func (this *GateApp) Bind(sid string, id int64) {
	common.Dosafe(func() {
		sessContent, ok := this.c2s.Load(sid)
		if !ok {
			akLog.Error("can not find client session, id: ", sid)
			return
		}
		if id > 0 {
			sessContent.(*C2SContext).session.Bind(id)
		}
		this.c2s.Store(sid, sessContent.(*C2SContext))
	}, nil)
}

//c->[gate1<=>gate2]->s
func (this *GateApp) CS_SendInner(sid string, id uint32, data []byte) {

	common.Dosafe(func() {

		c, ok := this.c2s.Load(sid)
		if !ok {
			akLog.Error("can not find client session, sid: ", sid)
			return
		}
		
		c.(*C2SContext).msgActor.SetSession(sid)

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
				refs = append(refs, reflect.ValueOf(c.(*C2SContext).msgActor))
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

		s, ok := this.s2s.Load(selectSID)
		if ok {
			c.(*C2SContext).SID = selectSID
			this.c2s.Store(sid, c.(*C2SContext))
		} else {
			akLog.Error("can not find server select server, selectSID: ", selectSID)
			return
		}

		packMsg := messageBase.SSPackMsg_pb(sid, s.(*S2SContext).session.GetUID(), akmessage.MSG_SS_ROUTE, &akmessage.SS_SSRoute{Data: data})
		if packMsg != nil {
			s.(*S2SContext).session.SendMsg(packMsg)
		}

	}, nil)
}

//[gate2<=>gate1]->c
func (this *GateApp) SendClient(sid string, id uint32, data []byte) {
	common.Dosafe(func() {
		sessContent, ok := this.c2s.Load(sid)
		if !ok {
			akLog.Error("can not find client session, id: ", sid)
			return
		}
		sessContent.(*C2SContext).session.Write(websocket.BinaryMessage, data)
	}, nil)
}

//game/login/...->gate->client
func (this *GateApp) Handler(sid string, data []byte) {

	common.Dosafe(func() {
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

				s, ok := this.s2s.Load(sessid)
				if !ok {
					akLog.Error("s2s can not find session, sid: ", sessid)
					return
				}
				akLog.FmtPrintln("register client dst:", reg.St)
				s.(*S2SContext).session.SetCliType(reg.St)
				s.(*S2SContext).msgActor.SetSession(sessid)
				this.s2s.Store(sessid, s.(*S2SContext))

				rsp := messageBase.GetActorRegisterRsp(sessid, akmessage.ServerType_Gate)
				s.(*S2SContext).session.SendMsg(rsp)
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

					s, ok := this.s2s.Load(sid)
					if !ok {
						akLog.Error("s2s can not find session, sid: ", sid)
						return
					}

					s.(*S2SContext).msgActor.SetSession(sid)
					this.s2s.Store(sid, s.(*S2SContext))

					var refs []reflect.Value
					refs = append(refs, reflect.ValueOf(s.(*S2SContext).msgActor))
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
		this.sendClient_ssc(sspt.GetSessID(), sspt.GetMsgID(), sspt.GetData())
	}, nil)
}

func (this *GateApp) sendClient_ssc(sid string, id uint32, data []byte) {
	common.Dosafe(func() {
		sessContent, ok := this.c2s.Load(sid)
		if !ok {
			akLog.Error("can not find client session, id: ", sid)
			return
		}
		cspt := messageBase.CSPackTool()
		cspt.Init(id, data)
		out := make([]byte, len(data)+messageBase.CS_MSG_PACK_NODATA_SIZE)
		cspt.Pack(out)
		sessContent.(*C2SContext).session.Write(websocket.BinaryMessage, out)
	}, nil)
}

func (this *GateApp) SS_SendInner(sid string, id uint32, data []byte) {
	common.Dosafe(func() {
		sessContent, ok := this.s2s.Load(sid)
		if !ok {
			akLog.Info("can not find client session: ", sid, id)
			return
		}
		sessContent.(*S2SContext).session.SendMsg(data)
	}, nil)
}

func (this *GateApp) Close() {

	common.Dosafe(func() {
		this.isclose = true

		this.c2s.Range(func(k interface{}, value interface{}) bool {
			value.(*C2SContext).session.Stop()
			this.c2s.Delete(k)
			return true
		})
		this.s2s.Range(func(k interface{}, value interface{}) bool {
			value.(*S2SContext).session.Stop()
			this.s2s.Delete(k)
			return true
		})
	}, nil)

}

func (this *GateApp) IsClose() bool {
	return this.isclose
}
