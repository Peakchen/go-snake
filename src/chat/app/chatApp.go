package app

import (
	"fmt"
	"go-snake/akmessage"
	"go-snake/common/evtAsync"
	"go-snake/common/messageBase"
	"go-snake/common/mixNet"
	"go-snake/common/tcpNet"
	"go-snake/chat/base"
	"go-snake/chat/entityBase"
	"go-snake/chat/msg"
	"reflect"

	"google.golang.org/protobuf/proto"

	"github.com/Peakchen/xgameCommon/akLog"
)

func Init() {
	mixNet.SetApp(NewApp())
}

//gate2 <-> server
type ChatApp struct {
	roles uint32

	session *tcpNet.TcpSession

	isclose bool
}

func NewApp() *ChatApp {
	return &ChatApp{
		session: nil,
	}
}

func (this *ChatApp) Online(nt messageBase.NetType, sess interface{}) {
	evtAsync.SendEvtFn(func() {
		switch nt {
		case messageBase.Net_WS:

		case messageBase.Net_TCP:
			this.session = sess.(*tcpNet.TcpSession)
			this.session.SendMsg(messageBase.GetActorRegisterReq(this.session.GetSessionID(), this.session.GetType()))
		}
	})
}

func (this *ChatApp) Offline(nt messageBase.NetType, id string) {
	evtAsync.SendEvtFn(func() {
		switch nt {
		case messageBase.Net_WS:

		case messageBase.Net_TCP:
			this.session = nil
		}
	})
}

func (this *ChatApp) Bind(sid string, id int64) {

}

func (this *ChatApp) CS_SendInner(sid string, id uint32, data []byte) {

}

func (this *ChatApp) SendClient(sid string, id uint32, data []byte) {

}

//gate->*
func (this *ChatApp) Handler(sid string, data []byte) {
	evtAsync.SendEvtFn(func() {
		sspt := messageBase.SSPackTool()
		err := sspt.UnPack(data)
		if err != nil {
			akLog.Error("ss upack fail.")
			return
		}

		msgid := sspt.GetMsgID()
		dstData := sspt.GetData()

		if msgid == uint32(akmessage.MSG_SS_ROUTE) {
			ssroute := messageBase.GetSSRoute()
			err := messageBase.Codec().Unmarshal(dstData, ssroute)
			if err != nil {
				akLog.Error(fmt.Errorf("unmarshal message fail, err: %v.", err))
				return
			}
			cspt := messageBase.CSPackTool()
			err = cspt.UnPack(ssroute.Data)
			if err != nil {
				akLog.Error(fmt.Errorf("cs unpack message fail, err: %v.", err, msgid))
				return
			}

			msgid = cspt.GetMsgID()
			dstData = cspt.GetData()
		}

		sessid := sspt.GetSessID()
		uid := sspt.GetUID()

		entity := base.GetUserByID(uid)
		switch msgid {
		case uint32(akmessage.MSG_SS_REGISTER_RSP),
			uint32(akmessage.MSG_SS_HEARTBEAT_RSP):
			if entity == nil {
				entity = entityBase.NewEntity(sessid, uid)
				base.AddUser(uid, entity)
			}
		}

		if entity == nil {
			akLog.Error("invalid uid: ", uid)
			return
		}

		content := msg.GetActorMessageProc(msgid)
		if content == nil {
			akLog.Error("invalid msgid: ", msgid)
			return
		}

		dst := reflect.New(content.RefPb.Elem()).Interface().(proto.Message)
		err = messageBase.Codec().Unmarshal(dstData, dst)
		if err != nil {
			akLog.Error(fmt.Errorf("unmarshal message fail, err: %v.", err))
			return
		}

		entity.SetSessionID(sessid)

		var refs []reflect.Value
		refs = append(refs, reflect.ValueOf(entity))
		refs = append(refs, reflect.ValueOf(dst))
		content.RefProc.Call(refs)
	})
}

//*->gate
func (this *ChatApp) SS_SendInner(sid string, id uint32, data []byte) {
	evtAsync.SendEvtFn(func() {
		if this.session == nil {
			akLog.Error("session disconnetced..., msg not send, id: ", id)
			return
		}
		this.session.SendMsg(data)
	})
}

func (this *ChatApp) Close() {
	this.session.Stop()
	this.session = nil
	this.isclose = true
}

func (this *ChatApp) IsClose() bool {
	return this.isclose
}
