package common

import (
	"runtime"
	"runtime/debug"

	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/tool"
)

func Dosafe(fn func(), exitfn func()) {
	defer ExceptionStack(exitfn)
	fn()
}

func DosafeRoutine(fn func(), exitfn func()) {
	defer ExceptionStack(exitfn)
	go fn()
}

func ExceptionStack(fn func()) {
	if fn != nil {
		fn()
	}
	err := recover()
	switch err.(type) {
	case runtime.Error:
		akLog.Error("runtime error: ", err, string(debug.Stack()))
	default:
		return
	}
}

func SafeExit(fn func(), exitfn func()) {
	defer ExceptionStack(exitfn)

	fn()
}

func SignalExit(fn func()) {
	tool.SignalExit(fn)
}
