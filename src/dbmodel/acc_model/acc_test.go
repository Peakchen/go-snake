package accdb

import (
	"go-snake/common/akOrm"
	"reflect"
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
	exist, err := akOrm.HasExistAcc(&Acc{}, userName, pwd)
	if err != nil {
		akLog.FmtPrintln("err: ", err)
		return
	}
	if exist {
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
	exist, err := akOrm.HasExistAcc(&Acc{}, userName, pwd)
	if err != nil {
		akLog.FmtPrintln("err: ", err)
		return
	}
	if exist {
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

func TestTableName(t *testing.T) {
	acc := &Acc{}
	akLog.FmtPrintln(reflect.TypeOf(*acc).Name())
}
