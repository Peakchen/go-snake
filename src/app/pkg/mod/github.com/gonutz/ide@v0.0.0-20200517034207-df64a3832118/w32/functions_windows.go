package w32

import (
	"syscall"
	"unsafe"
)

var (
	user32   = syscall.NewLazyDLL("user32.dll")
	kernel32 = syscall.NewLazyDLL("kernel32.dll")
)

var (
	defWindowProc            = user32.NewProc("DefWindowProcW")
	postQuitMessage          = user32.NewProc("PostQuitMessage")
	loadCursor               = user32.NewProc("LoadCursorW")
	registerClassEx          = user32.NewProc("RegisterClassExW")
	createWindowEx           = user32.NewProc("CreateWindowExW")
	getMessage               = user32.NewProc("GetMessageW")
	translateMessage         = user32.NewProc("TranslateMessage")
	dispatchMessage          = user32.NewProc("DispatchMessageW")
	messageBox               = user32.NewProc("MessageBoxW")
	loadImage                = user32.NewProc("LoadImageW")
	sendMessage              = user32.NewProc("SendMessageW")
	getWindowThreadProcessId = user32.NewProc("GetWindowThreadProcessId")
	showWindowAsync          = user32.NewProc("ShowWindowAsync")
	setTimer                 = user32.NewProc("SetTimer")
	getClientRect            = user32.NewProc("GetClientRect")
	registerRawInputDevices  = user32.NewProc("RegisterRawInputDevices")
	getKeyState              = user32.NewProc("GetKeyState")
	createAcceleratorTable   = user32.NewProc("CreateAcceleratorTableW")
	translateAccelerator     = user32.NewProc("TranslateAccelerator")
	setCapture               = user32.NewProc("SetCapture")
	releaseCapture           = user32.NewProc("ReleaseCapture")

	getModuleHandle     = kernel32.NewProc("GetModuleHandleW")
	getConsoleWindow    = kernel32.NewProc("GetConsoleWindow")
	getCurrentProcessId = kernel32.NewProc("GetCurrentProcessId")
)

func MakeIntResource(id uint16) *uint16 {
	return (*uint16)(unsafe.Pointer(uintptr(id)))
}

func DefWindowProc(window, msg, wParam, lParam uintptr) uintptr {
	ret, _, _ := defWindowProc.Call(window, msg, wParam, lParam)
	return ret
}

func PostQuitMessage(exitCode int) {
	postQuitMessage.Call(uintptr(exitCode))
}

func LoadCursor(instance uintptr, cursor int) uintptr {
	ret, _, _ := loadCursor.Call(instance, uintptr(cursor))
	return ret
}

func RegisterClassEx(wndClassEx *WNDCLASSEX) uintptr {
	ret, _, _ := registerClassEx.Call(uintptr(unsafe.Pointer(wndClassEx)))
	return ret
}

func CreateWindowEx(
	exStyle uintptr,
	className string,
	windowName string,
	style uintptr,
	x, y, width, height int,
	parent uintptr,
	menu uintptr,
	instance uintptr,
	param uintptr,
) uintptr {
	var classNamePtr *uint16
	if className != "" {
		classNamePtr = syscall.StringToUTF16Ptr(className)
	}

	var windowNamePtr *uint16
	if windowName != "" {
		windowNamePtr = syscall.StringToUTF16Ptr(windowName)
	}

	ret, _, _ := createWindowEx.Call(
		exStyle,
		uintptr(unsafe.Pointer(classNamePtr)),
		uintptr(unsafe.Pointer(windowNamePtr)),
		style,
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		parent,
		menu,
		instance,
		param,
	)
	return ret
}

func GetMessage(message *MSG, window, msgFilterMin, msgFilterMax uintptr) int {
	ret, _, _ := getMessage.Call(
		uintptr(unsafe.Pointer(message)),
		window,
		msgFilterMin,
		msgFilterMax,
	)
	return int(ret)
}

func TranslateMessage(message *MSG) bool {
	ret, _, _ := translateMessage.Call(uintptr(unsafe.Pointer(message)))
	return ret != 0
}

func DispatchMessage(message *MSG) uintptr {
	ret, _, _ := dispatchMessage.Call(uintptr(unsafe.Pointer(message)))
	return ret
}

func MessageBox(window uintptr, message, caption string, flags uint) uintptr {
	ret, _, _ := messageBox.Call(
		window,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(message))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(caption))),
		uintptr(flags),
	)
	return ret
}

func LoadImage(instance uintptr, name *uint16, typ, w, h, load uintptr) uintptr {
	ret, _, _ := loadImage.Call(
		instance,
		uintptr(unsafe.Pointer(name)),
		typ,
		w,
		h,
		load,
	)
	return ret
}

func SendMessage(window, message, wParam, lParam uintptr) uintptr {
	ret, _, _ := sendMessage.Call(
		window,
		message,
		wParam,
		lParam,
	)
	return ret
}

func GetWindowThreadProcessId(hwnd uintptr) (uintptr, uint32) {
	var processId uint32
	ret, _, _ := getWindowThreadProcessId.Call(
		hwnd,
		uintptr(unsafe.Pointer(&processId)),
	)
	return ret, processId
}

func ShowWindowAsync(window, commandShow uintptr) bool {
	ret, _, _ := showWindowAsync.Call(window, commandShow)
	return ret != 0
}

func SetTimer(window, idEvent, elapse uintptr) uintptr {
	ret, _, _ := setTimer.Call(
		window,
		idEvent,
		elapse,
		0,
	)
	return ret
}

func GetClientRect(window uintptr) (RECT, bool) {
	var r RECT
	ret, _, _ := getClientRect.Call(window, uintptr(unsafe.Pointer(&r)))
	return r, ret != 0
}

func RegisterRawInputDevices(devices []RAWINPUTDEVICE) bool {
	if len(devices) == 0 {
		return false
	}
	ret, _, _ := registerRawInputDevices.Call(
		uintptr(unsafe.Pointer(&devices[0])),
		uintptr(len(devices)),
		uintptr(unsafe.Sizeof(devices[0])),
	)
	return ret != 0
}

func GetKeyState(key uintptr) uint16 {
	ret, _, _ := getKeyState.Call(key)
	return uint16(ret)
}

func CreateAcceleratorTable(acc []ACCEL) uintptr {
	if len(acc) == 0 {
		return 0
	}
	ret, _, _ := createAcceleratorTable.Call(
		uintptr(unsafe.Pointer(&acc[0])),
		uintptr(len(acc)),
	)
	return ret
}

func TranslateAccelerator(window uintptr, accTableHandle uintptr, msg *MSG) bool {
	ret, _, _ := translateAccelerator.Call(
		window,
		accTableHandle,
		uintptr(unsafe.Pointer(msg)),
	)
	return ret != 0
}

func SetCapture(window uintptr) uintptr {
	ret, _, _ := setCapture.Call(window)
	return ret
}

func ReleaseCapture() bool {
	ret, _, _ := releaseCapture.Call()
	return ret != 0
}

func GetModuleHandle(moduleName string) uintptr {
	var name uintptr
	if moduleName != "" {
		name = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(moduleName)))
	}
	ret, _, _ := getModuleHandle.Call(name)
	return ret
}

func GetConsoleWindow() uintptr {
	ret, _, _ := getConsoleWindow.Call()
	return ret
}

func GetCurrentProcessId() uint32 {
	id, _, _ := getCurrentProcessId.Call()
	return uint32(id)
}
