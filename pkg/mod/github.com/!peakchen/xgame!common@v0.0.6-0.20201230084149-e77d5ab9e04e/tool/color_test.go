package tool

import "testing"

func TestColor(t *testing.T) {
	WinColorPrint(WinFontColor_Green, "222")
	LinuxDefaultBGPrint(LinuxForeground_Red, "1111")
	LinuxDefaultColorPrint("333")
}
