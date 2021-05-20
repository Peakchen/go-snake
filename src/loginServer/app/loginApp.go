package app

import (
	"fmt"
	"go-snake/akmessage"
	"go-snake/common"
	"go-snake/common/evtAsync"
	"go-snake/common/messageBase"
	"go-snake/common/mixNet"
	"go-snake/common/tcpNet"
	"go-snake/core/usermgr"
	userEntity "go-snake/core/user"
	"go-snake/core/msg"
	"reflect"
	"sync"

	"google.golang.org/protobuf/proto"

	"github.com/Peakchen/xgameCommon/akLog"
)

func Init() {
	mixNet.SetApp(NewApp())
}

//gate2 <-> game server
type LoginApp struct {
	roles uint32

	session sync.Map //*tcpNet.TcpSession

	isclose bool
}

func NewApp() *LoginApp {
	return &LoginApp{}
}

func (this *LoginApp) Online(nt messageBase.NetType, sess interface{}) {
	evtAsync.SendEvtFn(func() {
		switch nt {
		case messageBase.Net_WS:

		case messageBase.Net_TCP:
			session := sess.(*tcpNet.TcpSession)
			this.session.Store(session.GetSessionID(), session)
			session.SendMsg(messageBase.GetActorRegisterReq(session.GetSessionID(), session.GetType()))
		}
	})
}

func (this *LoginApp) Offline(nt messageBase.NetType, id string) {
	evtAsync.SendEvtFn(func() {
		switch nt {
		case messageBase.Net_WS:

		case messageBase.Net_TCP:
			this.session.Delete(id)
			akLog.Info("tcp offline: ", id)
		}
	})
}

func (this *LoginApp) Bind(sid string, id int64) {

}

func (this *LoginApp) CS_SendInner(sid string, id uint32, data []byte) {

}

func (this *LoginApp) SendClient(sid string, id uint32, data []byte) {
}

//gate2->login
func (this *LoginApp) Handler(sid string, data []byte) {
	common.Dosafe(func() {
		sessIf, _ := this.session.Load(sid)
		if sessIf == nil {
			akLog.Info("session disconnect, sid: ", sid)
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

		if len(sspt.GetSessID()) == 0 {
			akLog.Error("session id is null, id: ", msgid)
			return
		}

		if msgid == uint32(akmessage.MSG_SS_ROUTE) && len(dstData) > 0 && dstData != nil {
			ssroute := messageBase.GetSSRoute()
			err := messageBase.Codec().Unmarshal(dstData, ssroute)
			if err != nil {
				akLog.Error(fmt.Errorf("unmarshal message fail, err: %v, dstData: %v.", err, dstData))
				return
			}
			cspt := messageBase.CSPackTool()
			err = cspt.UnPack(ssroute.Data)
			if err != nil {
				akLog.Error(fmt.Errorf("cs unpack message fail, err: %v,msgid: %v.", err, msgid))
				return
			}

			msgid = cspt.GetMsgID()
			dstData = cspt.GetData()
		}

		content := msg.GetActorMessageProc(msgid)
		user := usermgr.GetUserByID(sspt.GetUID())
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
				uint32(akmessage.MSG_SS_REGISTER_RSP),
				uint32(akmessage.MSG_SS_HEARTBEAT_RSP):
				if user == nil {
					user = userEntity.NewEntityBySid(sspt.GetSessID())
					usermgr.AddUser(user.GetID(), user)
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
	}, nil)
}

//login->gate2
func (this *LoginApp) SS_SendInner(sid string, id uint32, data []byte) {
	//evtAsync.SendEvtFn(func() {
	this.session.Range(func(k interface{}, v interface{}) bool {
		v.(*tcpNet.TcpSession).SendMsg(data)
		return true
	})
}

func (this *LoginApp) Close() {

	this.session.Range(func(k interface{}, v interface{}) bool {
		v.(*tcpNet.TcpSession).Stop()
		return false
	})

	this.isclose = true
}

func (this *LoginApp) IsClose() bool {
	return this.isclose
}
