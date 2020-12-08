package Entity

import (
	"ak-remote/akmessage"
	"ak-remote/common/mixNet"
	"ak-remote/common/myTcpSocket"
	"ak-remote/gameServer/logic/base"

	"github.com/Peakchen/xgameCommon/akLog"

	"google.golang.org/protobuf/proto"
)

type EntityUser struct {
	base.IEntityUser
}

func (this *EntityUser) xxx() {

}

func (this *EntityUser) SendMsg(msgID akmessage.MSG, pb proto.Message) {
	sess := mixNet.GetSessionMgr().GetTcpSession(this.GetSessionID())
	if sess != nil {
		data := myTcpSocket.PackMsg(msgID, pb)
		if data == nil {
			akLog.Error("pack message fail.")
			return
		}
		sess.(*myTcpSocket.AkTcpSession).SendMsg(data)
	}
}
