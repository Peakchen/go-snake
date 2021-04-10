package common_test

import (
	"testing"
	"go-snake/common"

	"github.com/Peakchen/xgameCommon/akLog"

)

func TestTool(t *testing.T) {
	//common.SafeExit()
}

func TestLogInfo(t *testing.T) {
	akLog.Info("session close...")
}

func TestRoutine(t *testing.T) {
	common.DosafeRoutine(func(){
		
		type TA struct {
			A int
		}

		var ta *TA
		akLog.FmtPrintln(ta.A)

	},nil)
}
