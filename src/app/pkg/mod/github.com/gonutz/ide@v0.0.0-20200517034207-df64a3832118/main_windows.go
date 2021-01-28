//+build windows

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"syscall"
	"time"
	"unicode/utf16"
	"unsafe"

	"github.com/gonutz/ide/font"
	"github.com/gonutz/ide/w32"
)

var globalGraphics graphics

func main() {
	defer handlePanics()
	runtime.LockOSThread()
	hideConsoleWindow()
	window := createWindow()

	graphics, err := newD3d9Graphics(window, font.TTF, 20)
	if err != nil {
		panic(err)
	}
	defer graphics.close()
	globalGraphics = graphics

	w32.SetTimer(window, 1, 50)

	var msg w32.MSG
	for w32.GetMessage(&msg, 0, 0, 0) > 0 {
		w32.TranslateMessage(&msg)
		w32.DispatchMessage(&msg)
	}
}

var text string

func handleOSMessage(window, message, w, l uintptr) uintptr {
	switch message {
	case w32.WM_TIMER:
		// TODO for now this is rendering some moving rectangles for redering;
		// eventually this will update the GUI if re-drawing is necessary
		globalGraphics.rect(0, 0, 100000, 100000, 0xFF072727)
		globalGraphics.rect(10, 10, 200, 200, 0xFFFFFFFF)
		globalGraphics.text(
			[]byte(text),
			10, 10,
			rect(10, 10, 200, 200),
			0xFF000000,
		)
		err := globalGraphics.present()
		if err != nil {
			panic(err)
		}
		return 0
	case w32.WM_DESTROY:
		w32.PostQuitMessage(0)
		return 0
	case w32.WM_COMMAND:
		cmd := w & 0xFFFF
		fmt.Println("            command:", cmd)
		return 0
	case w32.WM_SYSCOMMAND:
		fmt.Printf("WM_SYSCOMMAND %x %x %s\n", w, l, keyState())
		return w32.DefWindowProc(window, message, w, l)
	case w32.WM_KEYDOWN:
		fmt.Printf("WM_KEYDOWN %x(%s) %x %s\n", w, string(rune(w)), l, keyState())
		return w32.DefWindowProc(window, message, w, l)
	case w32.WM_CHAR:
		r := utf16.Decode([]uint16{uint16(w)})[0]
		repeatCount := int(l & 0xFFFF)
		scanCode := int(l&0xFF0000) >> 16
		extended := l&1<<24 != 0
		altDown := l&1<<29 != 0
		wasDownBefore := l&1<<30 != 0
		beingReleased := l&1<<31 != 0
		fmt.Printf(`WM_CHAR
           key  %s (%d)
  repeat count  %d
     scan code  %d
      extended  %v
      alt down  %v
  key was down  %v
   is released  %v
`,
			string(r), int(r),
			repeatCount,
			scanCode,
			extended,
			altDown,
			wasDownBefore,
			beingReleased,
		)
		return w32.DefWindowProc(window, message, w, l)
	default:
		return w32.DefWindowProc(window, message, w, l)
	}
}

func keyState() string {
	result := ""
	const mask = 1 << 15
	if w32.GetKeyState(w32.VK_SHIFT)&mask != 0 {
		result += " shift"
	}
	if w32.GetKeyState(w32.VK_CONTROL)&mask != 0 {
		result += " control"
	}
	if w32.GetKeyState(w32.VK_MENU)&mask != 0 {
		result += " alt"
	}
	return result
}

func handlePanics() {
	// After a panic the user/developer is shown the stack trace. To be sure
	// that the message is seen, it is not only printed to stdout but also saved
	// to disk and a message box pops up.
	if err := recover(); err != nil {
		message := fmt.Sprintf("panic: %v\nstack:\n\n%s\n", err, debug.Stack())

		// print to standard output
		fmt.Println(message)

		// write as a log file to disk
		logFile := filepath.Join(
			os.Getenv("APPDATA"),
			"ide_log_"+time.Now().Format("2006_01_02__15_04_05")+".txt",
		)
		ioutil.WriteFile(logFile, []byte(message), 0666)

		// open crash log file with the default text viewer
		exec.Command("cmd", "/C", logFile).Start()

		// pop up a message box
		w32.MessageBox(
			0,
			message,
			"Go IDE crashed",
			w32.MB_OK|w32.MB_ICONERROR|w32.MB_TOPMOST,
		)
	}
}

// hideConsoleWindow closes the console window that opens if you 'go build' on
// Windows without specifying -ldflags "-H=windowsgui.
func hideConsoleWindow() {
	console := w32.GetConsoleWindow()
	if console != 0 {
		// If this application is the process that created the console window,
		// then this program was not compiled with the -H=windowsgui flag and on
		// start-up it created a console along with the main application window.
		// In this case hide the console window. See
		// http://stackoverflow.com/questions/9009333/how-to-check-if-the-program-is-run-from-a-console
		_, consoleProcID := w32.GetWindowThreadProcessId(console)
		if w32.GetCurrentProcessId() == consoleProcID {
			w32.ShowWindowAsync(console, w32.SW_HIDE)
		}
	}
}

func createWindow() uintptr {
	const windowClassName = "GoIDEWindowClass"
	class := w32.WNDCLASSEX{
		WndProc:   syscall.NewCallback(handleOSMessage),
		Cursor:    w32.LoadCursor(0, w32.IDC_ARROW),
		ClassName: syscall.StringToUTF16Ptr(windowClassName),
	}
	class.Size = uint32(unsafe.Sizeof(class))

	atom := w32.RegisterClassEx(&class)
	if atom == 0 {
		panic("RegisterClassEx failed")
	}

	window := w32.CreateWindowEx(
		0,
		windowClassName,
		"Go IDE",
		w32.WS_OVERLAPPEDWINDOW|w32.WS_VISIBLE,
		10, 10, 850, 800,
		0, 0, 0, 0,
	)
	if window == 0 {
		panic("CreateWindowEx failed")
	}

	// the icon is contained in the .exe file as a resource, load it and set it
	// as the window icon so it appears in the top-left corner of the window and
	// when you alt+tab between windows
	const iconResourceID = 10
	icon := uintptr(w32.LoadImage(
		w32.GetModuleHandle(""),
		w32.MakeIntResource(iconResourceID),
		w32.IMAGE_ICON,
		0,
		0,
		w32.LR_DEFAULTSIZE|w32.LR_SHARED,
	))
	if icon != 0 {
		w32.SendMessage(window, w32.WM_SETICON, w32.ICON_SMALL, icon)
		w32.SendMessage(window, w32.WM_SETICON, w32.ICON_SMALL2, icon)
		w32.SendMessage(window, w32.WM_SETICON, w32.ICON_BIG, icon)
	}

	return window
}
