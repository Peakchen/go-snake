package app

import (
	"fmt"
	"go-snake/akmessage"
	//"go-snake/common/evtAsync"
	"go-snake/common/messageBase"
	"go-snake/common/mixNet"
	"go-snake/common/tcpNet"
	"go-snake/core/usermgr"
	"go-snake/core/user"
	"go-snake/core/msg"
	"reflect"

	"google.golang.org/protobuf/proto"

	"github.com/Peakchen/xgameCommon/akLog"
)

func Init() {
	mixNet.SetApp(NewApp())
}

//gate2 <-> World server
type WorldApp struct {
	roles uint32

	session *tcpNet.TcpSession

	isclose bool
}

func NewApp() *WorldApp {
	return &WorldApp{
		session: nil,
	}
}

func (this *WorldApp) Online(nt messageBase.NetType, sess interface{}) {
	
	switch nt {
	case messageBase.Net_WS:

	case messageBase.Net_TCP:
		this.session = sess.(*tcpNet.TcpSession)
		this.session.SendMsg(messageBase.GetActorRegisterReq(this.session.GetSessionID(), this.session.GetType()))
	}

}

func (this *WorldApp) Offline(nt messageBase.NetType, id string) {
	
	switch nt {
	case messageBase.Net_WS:

	case messageBase.Net_TCP:
		this.session = nil
	}

}

func (this *WorldApp) Bind(sid string, id int64) {

}

func (this *WorldApp) CS_SendInner(sid string, id uint32, data []byte) {

}

func (this *WorldApp) SendClient(sid string, id uint32, data []byte) {

}

//gate->World
func (this *WorldApp) Handler(sid string, data []byte) {
	
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

	entity := usermgr.GetUserByID(uid)
	switch msgid {
	case uint32(akmessage.MSG_CS_ENTER_GAME_SCENE),
		uint32(akmessage.MSG_SS_REGISTER_RSP),
		uint32(akmessage.MSG_SS_HEARTBEAT_RSP):
		if entity == nil {
			entity = user.NewEntityBySid(sessid)
			usermgr.AddUser(entity.GetID(), entity)
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

}

//World->gate
func (this *WorldApp) SS_SendInner(sid string, id uint32, data []byte) {
	
	if this.session == nil {
		akLog.Error("session disconnetced..., msg not send, id: ", id)
		return
	}
	this.session.SendMsg(data)

}

func (this *WorldApp) Close() {
	this.session.Stop()
	this.session = nil
	this.isclose = true
}

func (this *WorldApp) IsClose() bool {
	return this.isclose
}
