package chat


import (
	"go-snake/akmessage"
	"go-snake/core/user"
	"go-snake/core/msg"

	"github.com/Peakchen/xgameCommon/akLog"
	"google.golang.org/protobuf/proto"
)


func Register(){

	msg.RegisterActorMessageProc(uint32(akmessage.MSG_CS_CHAT), (*akmessage.CS_Chat)(nil),
	func(actor user.IEntityUser, pb proto.Message) {
		if actor != nil {	actor.HandleChat(pb)	}else{	akLog.Error("invalid actor object.")	}
	})


}

func (this *Chat) HandleChat(pb proto.Message){

}