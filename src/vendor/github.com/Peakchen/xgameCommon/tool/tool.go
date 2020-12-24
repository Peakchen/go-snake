package tool

// add by stefan

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"
	//"runtime"
	//"syscall"
)

func CheckPortUsed(port int) bool {
	args := []string{"cmd", "/c", "netstat", "-an", "|", "findstr", strconv.Itoa(port)}
	cmd := exec.Command(args[0], args[1:]...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error() + ": " + stderr.String())
		return false
	}

	return true
}

func CmdHide() {
	handler, err := os.OpenFile("hide.vbs", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}

	vbscontext := `Set ws = CreateObject("Wscript.Shell") 
				ws.run "cmd /c zxcad.exe",vbhide`

	handler.WriteString(vbscontext)
	handler.Close()

	args := []string{"cmd", "/c", "hide.vbs"}
	cmd := exec.Command(args[0], args[1:]...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println(err.Error() + ": " + stderr.String())
		return
	}

	err = os.Remove("hide.vbs")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func BatHide(param []string) {
	handler, err := os.OpenFile("hide.bat", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}

	batpath := "\"%~nx0 h\""
	batContext := "@echo off \r\n" +
		"if %1 == h goto begin \r\n" +
		//"start " + param[0] + " \r\n" +
		"mshta vbscriptcreateobject(wscript.shell).run(" + batpath + ",0) \r\n" +
		//"mshta vbscriptcreateobject(wscript.shell).run(%~nx0 h,0)(window.close)&&exit \r\n" +
		param[0] + " \r\n" +
		":begin \r\n" +
		"REM"

	handler.WriteString(batContext)
	handler.Close()

	args := []string{"cmd", "/c", "hide.bat", param[1]}
	cmd := exec.Command(args[0], args[1:]...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println(err.Error() + ": " + stderr.String())
		return
	}

	err = os.Remove("hide.bat")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func SignalExit(fn func()) {
	chsignal := make(chan os.Signal, 1)
	//listen sign: ctrl+c, kill
	signal.Notify(chsignal, os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)

signalLoop:
	for {
		select {
		case s := <-chsignal:
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGQUIT:
				if fn != nil {
					fn()
					break signalLoop
				}
			}
		}
	}
	fmt.Println("signalLoop out....")
}
