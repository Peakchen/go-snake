package common

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"

	"github.com/Peakchen/xgameCommon/akLog"
)

func Dosafe(fn func(), exitfn func()) {
	defer ExceptionStack(exitfn)
	fn()
}

func DosafeRoutine(fn func(), exitfn func()) {
	defer ExceptionStack(exitfn)
	go fn()
}

func callerDebug() (str string) {
	str += "\n"
	i := 2
	for {
		pc, file, line, ok := runtime.Caller(i)
		if !ok || i > 10 {
			break
		}
		str += fmt.Sprintf("\t stack: %d %v [file: %s] [func: %s] [line: %d]\n", i-1, ok, file, runtime.FuncForPC(pc).Name(), line)
		i++
	}
	return str + "\n"
}

const (
	separator = "==========================="
)

func ExceptionStack(fn func()) {
	err := recover()
	if err != nil {
		errstr := fmt.Sprintf("\n%s runtime error: %v\n traceback:\n", separator, err)
		errstr += callerDebug()
		errstr += separator + "\n"
		akLog.Error(errstr, string(debug.Stack()))
	}
}

func SafeExit() {
	akLog.Error("\nunknow exception, exit: \n", separator, callerDebug(), string(debug.Stack()), separator+"\n")
	os.Exit(1)
}
