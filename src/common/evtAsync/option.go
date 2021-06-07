package evtAsync

type OptFn func(*EventOption)

//具体事件实现
type IEvent interface {
	TaskHandler(t interface{})
	EventHandler(e interface{})
}

type EventOption struct {
	Task  	interface{}
	Event 	interface{}
	Func 	interface{}
}

func (this *EventOption) TaskHandler(t interface{}) {

}

func (this *EventOption) EventHandler(e interface{}) {

}

func loadEventOpts(optFns ...OptFn) *EventOption {
	eopts := new(EventOption)
	for _, f := range optFns {
		f(eopts)
	}
	return eopts
}

func WithTask(t interface{}) OptFn {
	return func(ep *EventOption) {
		ep.Task = t
	}
}

func WithEvent(e interface{}) OptFn {
	return func(ep *EventOption) {
		ep.Event = e
	}
}
