package main

import (
	"go-snake/app/in"
	"go-snake/app/mgr"
	"go-snake/common"
	"go-snake/common/akOrm"
	"go-snake/common/mixNet"
	"time"

	"github.com/Peakchen/xgameCommon/akLog"

	"github.com/Peakchen/xgameCommon/tool"
)

func main() {
	common.Dosafe(func() {
		params := in.ParseInput()
		if params == nil {
			panic("invalid input params.")
		}
		akLog.FmtPrintln("input params: ", *params)
		app := mgr.GetApp(params.AppName)
		if app != nil {
			app.Init()
			app.Run(params)

			tool.SignalExit(func() {
				//1.close session from c-s,then s-s to protect message consume
				mixNet.GetApp().Close()
				//2.save db
				akOrm.Stop()
				//3.exit process
				time.Sleep(3 * time.Second)
				common.SafeExit()
			})
		} else {
			panic("invalid app name.")
		}
	}, common.SafeExit)
}
