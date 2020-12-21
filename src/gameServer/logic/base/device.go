package base

import (
	"go-snake/akmessage"
	"go-snake/gameServer/logic/Device/device_model"
)

type IDevice interface {
	//func
	LoadDevices(dev *device_model.ClientDevice)

	//message
	SyncClientDevice(msg *akmessage.CS_SyncClientDeviceInfo)
}
