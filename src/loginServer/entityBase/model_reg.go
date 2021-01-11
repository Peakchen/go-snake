package entityBase

type ModelFn func(entity IEntityUser) interface{}

var _ms = map[int]ModelFn{}

func RegisterModel(modelID int, fn ModelFn) {
	_ms[modelID] = fn
}

func RangeModels(fn func(int, ModelFn)) {
	for id, cb := range _ms {
		fn(id, cb)
	}
}

