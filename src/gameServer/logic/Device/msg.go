package Device

import (
	"ak-remote/akmessage"
	"ak-remote/common/myWebSocket"
	"ak-remote/gameServer/logic/base"
)

func init() {
	base.Register(func(entity EntityUser) {
		myWebSocket.MsgRegister(akmessage.MSG_CS_SYNC_CLIENT_DEVICE_INFO, entity.SyncClientDevice)
	})

}

func (this *DeviceInfo) SyncClientDevice(msg *akmessage.CS_SyncClientDeviceInfo) {

	return
}
