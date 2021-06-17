package friend

import (
	"go-snake/akmessage"
	"go-snake/core/user"
	"go-snake/core/msg"

	"github.com/Peakchen/xgameCommon/akLog"
	"google.golang.org/protobuf/proto"
)


func Register(){

	msg.RegisterActorMessageProc(uint32(akmessage.MSG_CS_FRIENDLIST), (*akmessage.CS_FriendList)(nil),
	func(actor user.IEntityUser, pb proto.Message) {
		if actor != nil {	actor.HandleFriendList(pb)	}else{	akLog.Error("invalid actor object.")	}
	})

	msg.RegisterActorMessageProc(uint32(akmessage.MSG_CS_FRIENDADD), (*akmessage.CS_FriendAdd)(nil),
	func(actor user.IEntityUser, pb proto.Message) {
		if actor != nil {	actor.HandleFriendAdd(pb)	}else{	akLog.Error("invalid actor object.")	}
	})

	msg.RegisterActorMessageProc(uint32(akmessage.MSG_CS_FRIENDDELETE), (*akmessage.CS_FriendDelete)(nil),
	func(actor user.IEntityUser, pb proto.Message) {
		if actor != nil {	actor.HandleFriendDelete(pb)	}else{	akLog.Error("invalid actor object.")	}
	})

	msg.RegisterActorMessageProc(uint32(akmessage.MSG_CS_FRIENDQUERY), (*akmessage.CS_FriendQuery)(nil),
	func(actor user.IEntityUser, pb proto.Message) {
		if actor != nil {	actor.HandleFriendQuery(pb)	}else{	akLog.Error("invalid actor object.")	}
	})

	msg.RegisterActorMessageProc(uint32(akmessage.MSG_CS_FRIENDBLACK), (*akmessage.CS_FriendBlack)(nil),
	func(actor user.IEntityUser, pb proto.Message) {
		if actor != nil {	actor.HandleFriendBlack(pb)	}else{	akLog.Error("invalid actor object.")	}
	})

}

func (this *Friend) HandleFriendList(pb proto.Message){

}

func (this *Friend) HandleFriendAdd(pb proto.Message){

}

func (this *Friend) HandleFriendDelete(pb proto.Message){

}

func (this *Friend) HandleFriendQuery(pb proto.Message){

}

func (this *Friend) HandleFriendBlack(pb proto.Message){

}