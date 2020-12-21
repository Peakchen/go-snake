package acc_model

import (
	"go-snake/common/akOrm"
	"testing"
	"time"

	"github.com/Peakchen/xgameCommon/akLog"
)

func init() {
	akOrm.OpenDB("root", "OPbySYJ1FHfl9376ybgzZHtWqt2rcA51", "127.0.0.1:3306", "test")
	time.Sleep(5 * time.Second)
}

func TestAcc(t *testing.T) {
	userName := "111"
	pwd := "222"
	if akOrm.HasExistAcc(&Acc{}, userName, pwd) {
		akLog.FmtPrintln("acc is existed.")
		return
	}
	acc := NewAcc(userName, pwd)
	if acc == nil {
		akLog.FmtPrintln("acc is existed.")
	} else {
		akLog.FmtPrintln("acc: ", acc)
	}

}

func TestAccs(t *testing.T) {
	userName := "333"
	pwd := "444"
	if akOrm.HasExistAcc(&Acc{}, userName, pwd) {
		akLog.FmtPrintln("acc is existed.")
		return
	}
	acc := NewAcc(userName, pwd)
	if acc == nil {
		akLog.FmtPrintln("acc is existed.")
	} else {
		time.Sleep(1 * time.Second)
		accs := acc.Load()
		for _, a := range accs {
			akLog.FmtPrintln("acc: ", a)
		}

	}
}
