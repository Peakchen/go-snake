package model

import (
	"go-snake/akmessage"
	"go-snake/robot/model/login"
	"reflect"
	"testing"

	"github.com/Peakchen/xgameCommon/akLog"
	"google.golang.org/protobuf/proto"
)

func TestM(t *testing.T) {
	var login = &login.Login{}
	refFn := reflect.ValueOf(login).MethodByName("SC_ACC_REGISTER")
	if refFn.IsValid() {
		//var params []reflect.Value
		//params = append(params, reflect.ValueOf(&akmessage.SC_AccRegister{}))
		//refFn.Call(params)

		data, err := proto.Marshal(&akmessage.SC_AccRegister{})
		if err != nil {
			akLog.Error("Marshal fail")
			return
		}
		dst := reflect.New(refFn.Type().In(0).Elem()).Interface().(proto.Message)
		err = proto.Unmarshal(data, dst)
		if err != nil {
			akLog.Error("unmarshal fail")
			return
		}
		refFn.Call([]reflect.Value{reflect.ValueOf(dst)})
	}
	//akLog.FmtPrintln(reflect.ValueOf(login).MethodByName("SC_ACC_REGISTER").IsValid())
}
