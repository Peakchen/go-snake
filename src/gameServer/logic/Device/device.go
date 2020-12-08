package Device

import (
	"ak-remote/gameServer/logic/Device/device_model"
	"ak-remote/gameServer/logic/Entity"
	"ak-remote/gameServer/logic/base"
)

type DeviceInfo struct {
	Entity.EntityUser

	Devices []*device_model.ClientDevice
}

func init() {
	base.SetDevice(&DeviceInfo{})
}

func (this *DeviceInfo) LoadDevices(dev *device_model.ClientDevice) {

}
