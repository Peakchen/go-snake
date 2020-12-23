package base

import "reflect"

type RunModelFn func(v reflect.Value)

type ModelRecvFn func(obj interface{}) bool
