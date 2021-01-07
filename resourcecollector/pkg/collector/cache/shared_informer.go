package cache

// ResourceEventHandlerFuncs is an adaptor to let you easily specify as many or
// as few of the notification functions as you want while still implementing
// ResourceEventHandler.
type ResourceEventHandlerFuncs struct {
	ListFunc func(obj []interface{})
}

// OnAdd calls AddFunc if it's not nil.
func (r ResourceEventHandlerFuncs) OnList(obj []interface{}) {
	if r.ListFunc != nil {
		r.ListFunc(obj)
	}
}
