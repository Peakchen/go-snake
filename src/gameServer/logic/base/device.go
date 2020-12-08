package base

import (
	"ak-remote/akmessage"
	"ak-remote/gameServer/logic/Device/device_model"
)

type IDevice interface {
	//func
	LoadDevices(dev *device_model.ClientDevice)

	//message
	SyncClientDevice(msg *akmessage.CS_SyncClientDeviceInfo)
}
