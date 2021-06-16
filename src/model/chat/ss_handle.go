package chat


import (
	"go-snake/akmessage"
	"go-snake/core/user"
	"go-snake/core/msg"

	"github.com/Peakchen/xgameCommon/akLog"
	"google.golang.org/protobuf/proto"
)


func SSRegister(){

	msg.RegisterActorMessageProc(uint32(akmessage.MSG_SS_CHAT), (*akmessage.SS_Chat)(nil),
	func(actor user.IEntityUser, pb proto.Message) {
		if actor != nil {	actor.HandleChat_ss(pb)	}else{	akLog.Error("invalid actor object.")	}
	})


}


func (this *Chat) HandleChat_ss(pb proto.Message){
	
}