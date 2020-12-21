package Device

import (
	"go-snake/gameServer/logic/Device/device_model"
	"go-snake/gameServer/logic/Entity"
	"go-snake/gameServer/logic/base"
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
