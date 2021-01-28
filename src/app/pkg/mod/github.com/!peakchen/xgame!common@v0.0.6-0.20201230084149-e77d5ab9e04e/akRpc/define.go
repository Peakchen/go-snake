package akRpc

import (
	"reflect"
)

const (
	rpcdealline    int32 = 60 * 5 // five min ove time
	ActChanMaxSize int   = 1000   // act call params
)

type TModelActV1 struct {
	actid  string
	mod    interface{}
	modt   string
	modf   reflect.Value
	params []reflect.Value
}

type TModelActV2 struct {
	modf   reflect.Value
	params []reflect.Value
}

type TActRet struct {
	actid string
	rets  []reflect.Value
}
