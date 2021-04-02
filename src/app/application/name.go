package application

import (
	"strconv"
	"time"
)

var AppStr string

func SetAppName(origin string) {
	AppStr = origin + ":" + strconv.Itoa(int(time.Now().UnixNano()))
}

func GetAppName() string{
	return AppStr
}