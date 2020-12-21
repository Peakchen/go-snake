package Entity

import (
	"go-snake/akmessage"
	"go-snake/common/messageBase"
	"go-snake/common/mixNet"
	"go-snake/common/tcpNet"
	"go-snake/gameServer/logic/base"

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
		data := messageBase.SSPackMsg(this.GetSessionID(), this.GetID(), msgID, pb)
		if data == nil {
			akLog.Error("pack message fail.")
			return
		}
		sess.(*tcpNet.TcpSession).SendMsg(data)
	}
}
