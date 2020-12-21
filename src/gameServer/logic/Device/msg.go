package Device

import (
	"go-snake/akmessage"
	"go-snake/common/webNet"
	"go-snake/gameServer/logic/base"
)

func init() {
	base.Register(func(entity EntityUser) {
		webNet.MsgRegister(akmessage.MSG_CS_SYNC_CLIENT_DEVICE_INFO, entity.SyncClientDevice)
	})

}

func (this *DeviceInfo) SyncClientDevice(msg *akmessage.CS_SyncClientDeviceInfo) {

	return
}
