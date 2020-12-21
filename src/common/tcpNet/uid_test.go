package tcpNet

import (
	"testing"

	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/utils"
)

func TestUID(t *testing.T) {
	uid := utils.GetUUID()
	akLog.FmtPrintln("uid: ", uid, len(uid))
}
