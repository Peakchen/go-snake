package messageBase

import (
	"fmt"
	"reflect"

	"google.golang.org/protobuf/proto"

	"github.com/Peakchen/xgameCommon/akLog"
)

type TMessageProc struct {
	proc       reflect.Value
	paramTypes []reflect.Type
}

func (this *TMessageProc) Call(data []byte) error {
	dst := reflect.New(this.paramTypes[1].Elem()).Interface()
	err := proto.Unmarshal(data, dst.(proto.Message))
	if err != nil {
		return fmt.Errorf("unmarshal message fail, err: %v.", err)
	}
	msg := dst.(proto.Message)
	params := []reflect.Value{
		reflect.ValueOf(msg),
	}

	this.proc.Call(params)
	return nil
}

var (
	msgs = map[uint32]*TMessageProc{}
)

func MsgRegister(id uint32, proc interface{}) {

	cbref := reflect.TypeOf(proc)
	if cbref.Kind() != reflect.Func {
		akLog.FmtPrintln("proc type not is func, but is: %v.", cbref.Kind())
		return
	}

	if cbref.NumIn() != 2 {
		akLog.FmtPrintln("proc num input is not 2, but is: %v.", cbref.NumIn())
		return
	}

	paramtypes := []reflect.Type{}
	for i := 0; i < cbref.NumIn(); i++ {
		t := cbref.In(i)
		paramtypes = append(paramtypes, t)
	}

	msgs[id] = &TMessageProc{
		proc:       reflect.ValueOf(proc),
		paramTypes: paramtypes,
	}
}

func MsgHandler(id uint32) *TMessageProc {
	return msgs[id]
}

func GetMsgHandlers() map[uint32]*TMessageProc {
	return msgs
}
