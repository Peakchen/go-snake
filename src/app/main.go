package main

import (
	"context"
	"go-snake/app/in"
	"go-snake/app/mgr"
	"go-snake/common"
	"go-snake/common/akOrm"
	"go-snake/common/evtAsync"
	"go-snake/common/mixNet"
	"net/http"
	"time"

	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/pprof"
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

			pprof.Run(context.Background())
			common.DosafeRoutine(func() {
				http.ListenAndServe(params.Scfg.PprofHost, nil)
			}, nil)

			tool.SignalExit(func() {
				//1.close session from c-s,then s-s to protect message consume
				mixNet.GetApp().Close()
				//2.exit main goroutine
				evtAsync.Stop()
				//3.save db
				akOrm.Stop()
				//4.sleep 3 sec then exit process
				time.Sleep(3 * time.Second)
				common.SafeExit()
			})
		} else {
			panic("invalid app name.")
		}
	}, common.SafeExit)
}
