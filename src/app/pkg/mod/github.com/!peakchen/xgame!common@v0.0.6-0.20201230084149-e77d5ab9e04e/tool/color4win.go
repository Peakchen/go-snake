package tool

//add by stefan for linux windows.

import (
	"fmt"
	"syscall"
)

var (
	sysKernel32    *syscall.LazyDLL  = syscall.NewLazyDLL(`kernel32.dll`)
	sysProc        *syscall.LazyProc = sysKernel32.NewProc(`SetConsoleTextAttribute`)
	sysCloseHandle *syscall.LazyProc = sysKernel32.NewProc(`CloseHandle`)
)

//Color
const (
	WinFontColor_Black        = 0x01
	WinFontColor_Blue         = 0x02
	WinFontColor_Green        = 0x03
	WinFontColor_Cyan         = 0x04
	WinFontColor_Red          = 0x05
	WinFontColor_Purple       = 0x06
	WinFontColor_Yellow       = 0x07
	WinFontColor_Light_gray   = 0x08
	WinFontColor_Gray         = 0x09
	WinFontColor_Light_blue   = 0xa
	WinFontColor_Light_green  = 0xb
	WinFontColor_Light_cyan   = 0xc
	WinFontColor_Light_red    = 0xd
	WinFontColor_Light_purple = 0xf
	WinFontColor_Light_yellow = 0x10
	WinFontColor_White        = 0x11
)

func WinColorPrint(color int, str string) {
	handle, _, _ := sysProc.Call(uintptr(syscall.Stdout), uintptr(color))
	defer sysCloseHandle.Call(handle)

	fmt.Println(str)
}

func WinColorChange(color int) {
	handle, _, _ := sysProc.Call(uintptr(syscall.Stdout), uintptr(color))
	defer sysCloseHandle.Call(handle)
}
