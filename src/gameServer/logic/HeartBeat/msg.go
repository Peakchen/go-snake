package Device

import (
	"go-snake/akmessage"
	"go-snake/common/webNet"
)

func init() {
	webNet.MsgRegister(akmessage.MSG_CS_HEARTBEAT)
}


