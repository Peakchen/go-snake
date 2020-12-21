package mixNet

import (
	"reflect"
	"testing"

	"github.com/Peakchen/xgameCommon/akLog"

	"google.golang.org/protobuf/proto"
)

type testCallBackFn func(int, proto.Message)

func msgCall(id int, msg proto.Message) {

}

func testRegister(id int, fn testCallBackFn) {
	fnt := reflect.TypeOf(fn)
	if fnt.NumIn() != 2 {
		akLog.FmtPrintln("invalid fn params: ", fnt.NumIn())
		return
	}
	var ins []reflect.Type
	for i := 0; i < fnt.NumIn(); i++ {
		in := fnt.In(i)
		akLog.FmtPrintln(in.Name(), in.Kind().String())
		ins = append(ins, in)
	}

}

func TestMsg(t *testing.T) {
	testRegister(1, msgCall)
}
