package app

import (
	"fmt"
	"go-snake/akmessage"
	"go-snake/common/messageBase"
	"go-snake/common/mixNet"
	"go-snake/common/tcpNet"
	"go-snake/loginServer/base"
	"go-snake/loginServer/entityMgr"
	"go-snake/loginServer/msg"
	"reflect"

	"google.golang.org/protobuf/proto"

	"github.com/Peakchen/xgameCommon/akLog"
)

func Init() {
	mixNet.SetApp(NewApp())
}

//gate2 <-> game server
type LoginApp struct {
	roles uint32

	session *tcpNet.TcpSession
}

func NewApp() *LoginApp {
	return &LoginApp{
		session: nil,
	}
}

func (this *LoginApp) Online(nt messageBase.NetType, sess interface{}) {
	switch nt {
	case messageBase.Net_WS:

	case messageBase.Net_TCP:
		this.session = sess.(*tcpNet.TcpSession)
		this.session.SendMsg(messageBase.GetActorRegisterReq(this.session.GetSessionID(), this.session.GetType()))
	}
}

func (this *LoginApp) Offline(nt messageBase.NetType, id string) {
	switch nt {
	case messageBase.Net_WS:

	case messageBase.Net_TCP:
		this.session = nil
	}
}

func (this *LoginApp) Bind(sid string, id int64) {

}

func (this *LoginApp) CS_SendInner(sid string, id uint32, data []byte) {

}

func (this *LoginApp) SendClient(sid string, id uint32, data []byte) {

}

//gate2->login
func (this *LoginApp) Handler(sid string, data []byte) {
	if this.session == nil {
		akLog.FmtPrintln("session disconnect...")
		return
	}

	sspt := messageBase.SSPackTool()
	err := sspt.UnPack(data)
	if err != nil {
		akLog.Error(err)
		return
	}

	msgid := sspt.GetMsgID()
	dstData := sspt.GetData()

	if sspt.GetMsgID() == uint32(akmessage.MSG_SS_ROUTE) {
		ssroute := messageBase.GetSSRoute()
		err := messageBase.Codec().Unmarshal(dstData, ssroute)
		if err != nil {
			akLog.Error(fmt.Errorf("unmarshal message fail, err: %v.", err))
			return
		}
		cspt := messageBase.CSPackTool()
		err = cspt.UnPack(ssroute.Data)
		if err != nil {
			akLog.Error(fmt.Errorf("cs unpack message fail, err: %v.", err))
			return
		}

		msgid = cspt.GetMsgID()
		dstData = cspt.GetData()
	}

	content := msg.GetActorMessageProc(msgid)
	user := base.GetUserByID(sspt.GetUID())
	if content != nil {
		dst := reflect.New(content.RefPb.Elem()).Interface().(proto.Message)
		err := messageBase.Codec().Unmarshal(dstData, dst)
		if err != nil {
			akLog.Error(fmt.Errorf("unmarshal message fail, err: %v.", err))
			return
		}

		switch msgid {
		case uint32(akmessage.MSG_CS_ACC_REGISTER),
			uint32(akmessage.MSG_CS_LOGIN),
			uint32(akmessage.MSG_SS_REGISTER_RSP):
			if user == nil {
				user = entityMgr.NewEntity(sspt.GetSessID())
				base.GetEntityMgr().SetEntityByID(user.GetID(), user)
			}
		default:
		}

		if user == nil {
			akLog.Error("can not get entity, mid:", sspt.GetMsgID())
			return
		}

		user.SetSessionID(sspt.GetSessID())

		var refs []reflect.Value
		refs = append(refs, reflect.ValueOf(user))
		refs = append(refs, reflect.ValueOf(dst))
		content.RefProc.Call(refs)
	}
}

//login->gate2
func (this *LoginApp) SS_SendInner(sid string, id uint32, data []byte) {
	if this.session == nil {
		akLog.Error("session disconnetced..., msg not send, id: ", id)
		return
	}
	this.session.SendMsg(data)
}
