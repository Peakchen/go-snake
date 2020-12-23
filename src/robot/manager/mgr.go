package manager

import (
	"go-snake/akmessage"
	"go-snake/common/messageBase"
	"go-snake/robot/RoboIF"
	"reflect"

	"github.com/Peakchen/xgameCommon/akLog"
	"google.golang.org/protobuf/proto"
)

var robotModels = make(map[string]RoboIF.IRobotModel)

func RegisterModel(m RoboIF.IRobotModel)   { robotModels[m.Name()] = m }
func GetModel(n string) RoboIF.IRobotModel { return robotModels[n] }

//初始化模块，以及发消息
func RangeModels(v reflect.Value) {
	for _, md := range robotModels {
		md.Init(v)
		md.Enter()
	}
}

//接收消息
func RangeRecv(v []reflect.Value) {
	for _, md := range robotModels {
		if md.Recv(func(obj interface{}) bool {
			mid := v[0].Interface().(uint32)
			data := v[1].Interface().([]byte)

			sMsgName, ok := akmessage.MSG_name[int32(mid)]
			if !ok {
				return false
			}
			refFn := reflect.ValueOf(obj).MethodByName(sMsgName)
			if !refFn.IsValid() {
				return false
			}
			dst := reflect.New(refFn.Type().In(0).Elem()).Interface().(proto.Message)
			err := messageBase.Codec().Unmarshal(data, dst)
			if err != nil {
				akLog.Error("unmarshal fail,id: ", mid)
				return false
			}
			refFn.Call([]reflect.Value{reflect.ValueOf(dst)})
			return true
		}) {
			return
		}
	}
}
