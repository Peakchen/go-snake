package main

import (
	"ak-remote/app/in"
	"ak-remote/app/mgr"
	"ak-remote/common"

	"github.com/Peakchen/xgameCommon/akLog"

	"github.com/Peakchen/xgameCommon/tool"
)

func main() {
	common.DosafeRoutine(func() {
		params := in.ParseInput()
		akLog.FmtPrintln("input params: ", *params)
		app := mgr.GetApp(params.AppName)
		if app != nil {
			app.Init()
			app.Run()
		}
	}, nil)
	tool.SignalExit(func() {

	})
	akLog.FmtPrintln("end.")
}
