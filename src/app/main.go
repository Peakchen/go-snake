package main

import (
	"go-snake/app/in"
	"go-snake/app/mgr"
	"go-snake/common"
	"time"

	"github.com/Peakchen/xgameCommon/akLog"

	"github.com/Peakchen/xgameCommon/tool"
)

func main() {
	common.Dosafe(func() {
		params := in.ParseInput()
		akLog.FmtPrintln("input params: ", *params)
		app := mgr.GetApp(params.AppName)
		if app != nil {
			app.Init()
			app.Run(params)
		}
	}, nil)
	tool.SignalExit(func() {
		time.Sleep(time.Second)
	})
	akLog.FmtPrintln("end.")
}
