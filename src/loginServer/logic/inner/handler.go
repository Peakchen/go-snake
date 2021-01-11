package inner

import (
	"go-snake/akmessage"
	"go-snake/loginServer/entityBase"
	"go-snake/loginServer/msg"

	"github.com/Peakchen/xgameCommon/akLog"
	"google.golang.org/protobuf/proto"
)

func init() {
	msg.RegisterActorMessageProc(uint32(akmessage.MSG_SS_REGISTER_RSP),
		(*akmessage.SS_Register_Resp)(nil),
		func(actor entityBase.IEntityUser, pb proto.Message) {
			actor.HandlerRegisterResp(pb)
		})

	msg.RegisterActorMessageProc(uint32(akmessage.MSG_SS_HEARTBEAT_RSP),
		(*akmessage.SS_HeartBeat_Rsp)(nil),
		func(actor entityBase.IEntityUser, pb proto.Message) {
			actor.HandlerHeartBeatResp(pb)
		})
}

func (this *ServerInner) HandlerRegisterResp(pb proto.Message) {
	akLog.FmtPrintln("register finish....")
}

func (this *ServerInner) HandlerHeartBeatResp(pb proto.Message) {
	akLog.FmtPrintln("recv Heart Beat....")
}
