package utils

import (
	"github.com/segmentio/ksuid"
)

//github.com/segmentio/ksuid:  1lY8IDBB1Jqigln4d3Kd4IB0KwM
func GetOnlyString_v4() string {
	return ksuid.New().String()
}
