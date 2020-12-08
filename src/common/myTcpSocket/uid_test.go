package myTcpSocket

import (
	"testing"

	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/utils"
)

func TestUID(t *testing.T) {
	akLog.FmtPrintln("uid: ", utils.GetUUID())
}
