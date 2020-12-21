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

	_ "go-snake/loginServer/logic/account"
	_ "go-snake/loginServer/logic/innner"

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

func (this *LoginApp) Bind(id int64) {

}

func (this *LoginApp) SendInner(sid string, id uint32, data []byte) {
	if this.session == nil || data == nil {
		return
	}
	this.session.SendMsg(data)
}

func (this *LoginApp) SendClient(sid string, id uint32, data []byte) {

}

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

	content := msg.GetActorMessageProc(sspt.GetMsgID())
	user := base.GetEntityMgr().GetEntityByID(sspt.GetUID())
	if content != nil {
		dst := reflect.New(content.RefPb.Elem()).Interface().(proto.Message)
		err := proto.Unmarshal(sspt.GetData(), dst)
		if err != nil {
			akLog.Error(fmt.Errorf("unmarshal message fail, err: %v.", err))
			return
		}

		switch sspt.GetMsgID() {
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
