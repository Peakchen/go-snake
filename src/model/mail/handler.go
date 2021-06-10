package mail

import (
	"go-snake/akmessage"
	"go-snake/core/user"
	"go-snake/core/msg"

	"github.com/Peakchen/xgameCommon/akLog"
	"google.golang.org/protobuf/proto"
)


func Register(){

	msg.RegisterActorMessageProc(uint32(akmessage.MSG_CS_MAILINFO), (*akmessage.CS_MailInfo)(nil),
	func(actor user.IEntityUser, pb proto.Message) {
		if actor != nil {	actor.HandleMailInfo(pb)	}else{	akLog.Error("invalid actor object.")	}
	})

	msg.RegisterActorMessageProc(uint32(akmessage.MSG_CS_MAILREAD), (*akmessage.CS_MailRead)(nil),
	func(actor user.IEntityUser, pb proto.Message) {
		if actor != nil {	actor.HandleMailRead(pb)	}else{	akLog.Error("invalid actor object.")	}
	})

	msg.RegisterActorMessageProc(uint32(akmessage.MSG_CS_MAILTAKE), (*akmessage.CS_MailTake)(nil),
	func(actor user.IEntityUser, pb proto.Message) {
		if actor != nil {	actor.HandleMailTake(pb)	}else{	akLog.Error("invalid actor object.")	}
	})

	msg.RegisterActorMessageProc(uint32(akmessage.MSG_CS_MAILONEKEYREAD), (*akmessage.CS_MailOneKeyRead)(nil),
	func(actor user.IEntityUser, pb proto.Message) {
		if actor != nil {	actor.HandleMailOneKeyRead(pb)	}else{	akLog.Error("invalid actor object.")	}
	})

	msg.RegisterActorMessageProc(uint32(akmessage.MSG_CS_MAILONEKEYTAKE), (*akmessage.CS_MailOneKeyTake)(nil),
	func(actor user.IEntityUser, pb proto.Message) {
		if actor != nil {	actor.HandleMailOneKeyTake(pb)	}else{	akLog.Error("invalid actor object.")	}
	})

	msg.RegisterActorMessageProc(uint32(akmessage.MSG_CS_MAILDELETE), (*akmessage.CS_MailDelete)(nil),
	func(actor user.IEntityUser, pb proto.Message) {
		if actor != nil {	actor.HandleMailDelete(pb)	}else{	akLog.Error("invalid actor object.")	}
	})

	msg.RegisterActorMessageProc(uint32(akmessage.MSG_CS_MAILONEKEYDELETE), (*akmessage.CS_MailOneKeyDelete)(nil),
	func(actor user.IEntityUser, pb proto.Message) {
		if actor != nil {	actor.HandleMailOneKeyDelete(pb)	}else{	akLog.Error("invalid actor object.")	}
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