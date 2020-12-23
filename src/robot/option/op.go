package option

import (
	"reflect"
)

type RobotOptionFn func(*RobotOption)
type RobotOption struct {
	ModelsRun  func(v reflect.Value)
	ModelsRecv func(v []reflect.Value)
}

func SortRobotOptions(opts ...RobotOptionFn) *RobotOption {
	ropts := new(RobotOption)
	for _, fn := range opts {
		fn(ropts)
	}
	return ropts
}

func WithModelsRun(fn func(v reflect.Value)) RobotOptionFn {
	return func(opts *RobotOption) {
		opts.ModelsRun = fn
	}
}

func WithModelRecv(fn func(v []reflect.Value)) RobotOptionFn {
	return func(opts *RobotOption) {
		opts.ModelsRecv = fn
	}
}
