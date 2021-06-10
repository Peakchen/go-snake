package mail

import (
	"go-snake/akmessage"
	"go-snake/core/user"
	"go-snake/core/msg"

	"github.com/Peakchen/xgameCommon/akLog"
	"google.golang.org/protobuf/proto"
)


func Register(){

	msg.RegisterActorMessageProc(uint32(akmessage.MSG_SS_REGISTER_RSP),
		(*akmessage.SS_Register_Resp)(nil),
		func(actor user.IEntityUser, pb proto.Message) {

			if actor != nil {
				actor.HandleMailInfo(pb)
			}else{
				akLog.Error("invalid actor object.")
			}
			
		})



}


func (this *Mail) HandleMailInfo(pb proto.Message){

}

func (this *Mail) HandleMailRead(pb proto.Message){

}

func (this *Mail) HandleMailTake(pb proto.Message){

}

func (this *Mail) HandleMailOneKeyRead(pb proto.Message){

}

func (this *Mail) HandleMailOneKeyTake(pb proto.Message){

}

func (this *Mail) HandleMailDelete(pb proto.Message) {

}

func (this *Mail) HandleMailOneKeyDelete(pb proto.Message) {

}