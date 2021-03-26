package app

import (
	"go-snake/akmessage"

	"github.com/Peakchen/xgameCommon/akLog"
	"google.golang.org/protobuf/proto"
)

type IGateMessage interface {
	Handler_CS_HeartBeat(pb proto.Message)
	Handler_SS_HeartBeat(pb proto.Message)
	//HandlerRegister(pb proto.Message)
}

type GateMessage struct {
	GateSession
}

func init() {
	
	RegisterActorMessageProc(
		uint32(akmessage.MSG_CS_HEARTBEAT),
		(*akmessage.CS_HeartBeat)(nil),
		func(actor IMessageActor, pb proto.Message) {
			actor.Handler_CS_HeartBeat(pb)
		})

	RegisterActorMessageProc(
		uint32(akmessage.MSG_SS_HEARTBEAT_REQ),
		(*akmessage.SS_HeartBeat_Req)(nil),
		func(actor IMessageActor, pb proto.Message) {
			actor.Handler_SS_HeartBeat(pb)
		})
		
}

func (this *GateMessage) Handler_CS_HeartBeat(pb proto.Message) {
	akLog.FmtPrintln("gate client heart beat.", pb.(*akmessage.CS_HeartBeat))
	this.SendMsg_cs(uint32(akmessage.MSG_SC_HEARTBEAT), &akmessage.SC_HeartBeat{})
}

func (this *GateMessage) Handler_SS_HeartBeat(pb proto.Message) {
	akLog.FmtPrintln("gate server heart beat.", pb.(*akmessage.SS_HeartBeat_Req))
	this.SendMsg_ss(uint32(akmessage.MSG_SS_HEARTBEAT_RSP), &akmessage.SS_HeartBeat_Rsp{})
}

func (this *GateMessage) HandlerRegister(pb proto.Message) {

}
