package utils

// add by stefan

import (
	"github.com/axgle/mahonia"
)

func GBKToUTF8(src string) string {
	return mahonia.NewDecoder("utf8").ConvertString(src)
}

func UTF8ToGBK(src string) string {
	return mahonia.NewEncoder("gbk").ConvertString(src)
}

/*
	for chinese encoding.
*/
func UTF8ToGB18030(src string) string {
	return mahonia.NewDecoder("GB18030").ConvertString(src)
}
