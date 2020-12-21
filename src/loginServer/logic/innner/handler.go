package innner

import (
	"go-snake/akmessage"
	"go-snake/loginServer/entityMgr"
	"go-snake/loginServer/msg"

	"github.com/Peakchen/xgameCommon/akLog"
	"google.golang.org/protobuf/proto"
)

func init() {
	msg.RegisterActorMessageProc(uint32(akmessage.MSG_SS_REGISTER_RSP),
		(*akmessage.SS_Register_Resp)(nil),
		func(actor entityMgr.IEntityUser, pb proto.Message) {
			actor.HandlerRegisterResp(pb)
		})
}

func (this *ServerInner) HandlerRegisterResp(pb proto.Message) {
	akLog.FmtPrintln("register finish....")
}

func (this *ServerInner) HandlerHeartBeatResp(pb proto.Message) {
	akLog.FmtPrintln("recv Heart Beat....")
}
