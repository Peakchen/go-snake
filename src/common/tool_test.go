package common

import (
	"testing"

	"github.com/Peakchen/xgameCommon/akLog"
)

func TestTool(t *testing.T) {
	SafeExit()
}

func TestLogInfo(t *testing.T) {
	akLog.Info("session close...")
}
