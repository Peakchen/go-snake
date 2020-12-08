package Device

import (
	"ak-remote/akmessage"
	"ak-remote/common/myWebSocket"
)

func init() {
	myWebSocket.MsgRegister(akmessage.MSG_CS_HEARTBEAT)
}


