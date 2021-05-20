package role

import (
	"go-snake/akmessage"
	"go-snake/core/user"
	"go-snake/core/msg"

	"github.com/Peakchen/xgameCommon/akLog"

	"google.golang.org/protobuf/proto"
)

func init() {
	msg.RegisterActorMessageProc(uint32(akmessage.MSG_CS_ENTER_GAME_SCENE), 
		(*akmessage.CS_EnterGameScene)(nil), 
		func(actor user.IEntityUser, pb proto.Message) {
		actor.HandlerEnter(pb)
	})
}

func (this *RoleCache) HandlerEnter(pb proto.Message) {
	akLog.FmtPrintln("game enter...")
	this.SendMsg(akmessage.MSG_SC_ENTER_GAME_SCENE, &akmessage.SC_EnterGameScene{})
}
