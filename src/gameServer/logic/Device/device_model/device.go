package device_model

import (
	"gorm.io/gorm"
)

type ClientDevice struct {
}

const (
	Identify_ClientDevice = "ClientDevice"
)

func (this *ClientDevice) Identify() string {
	return Identify_ClientDevice
}

func LoadDeviceData(session *gorm.Session) {

}
