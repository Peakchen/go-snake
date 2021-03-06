package inner

import (
	"go-snake/akmessage"
	"go-snake/core/user"
	"go-snake/core/msg"

	"github.com/Peakchen/xgameCommon/akLog"
	"google.golang.org/protobuf/proto"
)

func Register() {

	msg.RegisterActorMessageProc(uint32(akmessage.MSG_SS_REGISTER_RSP),
		(*akmessage.SS_Register_Resp)(nil),
		func(actor user.IEntityUser, pb proto.Message) {

			if actor != nil {
				actor.HandlerRegisterResp(pb)
			}else{
				akLog.Error("invalid actor object.")
			}
			
		})

	msg.RegisterActorMessageProc(uint32(akmessage.MSG_SS_HEARTBEAT_RSP),
		(*akmessage.SS_HeartBeat_Rsp)(nil),
		func(actor user.IEntityUser, pb proto.Message) {

			if actor != nil {
				actor.HandlerHeartBeatResp(pb)
			}else{
				akLog.Error("invalid actor object.")
			}

		})
}

func (this *ServerInner) HandlerRegisterResp(pb proto.Message) {
	akLog.FmtPrintln("register finish....")
}

func (this *ServerInner) HandlerHeartBeatResp(pb proto.Message) {
	akLog.FmtPrintln("recv Heart Beat....")
}
