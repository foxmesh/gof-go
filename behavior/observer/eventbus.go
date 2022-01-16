package observer

import (
	"fmt"
	"reflect"
	"sync"
)

// 事件总线方式实现观察者模式

// Bus 事件总线
type Bus interface {
	Subscribe(topic string, handler interface{}) error
	Publish(topic string, args ... interface{}) error
}

// AsyncEventBus 异步事件总线
type AsyncEventBus struct {
 handlers map[string][]reflect.Value
 lock sync.Mutex
}

func NewAsyncEventBus() *AsyncEventBus  {
	return &AsyncEventBus{
		handlers: map[string][]reflect.Value{},
		lock:sync.Mutex{},
	}
}

func (b *AsyncEventBus) Subscribe(topic string, handler interface{})error  {
	b.lock.Lock()
	defer b.lock.Unlock()

	v := reflect.ValueOf(handler)
	if v.Type().Kind() != reflect.Func{
		return fmt.Errorf("handler is not a function")
	}

	fn,ok:=b.handlers[topic]
	if !ok{
		fn = []reflect.Value{}
	}

	fn = append(fn, v)
	b.handlers[topic]=fn
	return nil
}

func (b *AsyncEventBus ) Publish(topic string, args ... interface{}) error {
	handlers, ok := b.handlers[topic]
	if !ok {
		return fmt.Errorf("not found handlers in topic:%s", topic)
	}

	params :=make([]reflect.Value, len(args))
	for _,arg:=range args{
		params = append(params,reflect.ValueOf(arg))
	}

	for i := range handlers{
		go handlers[i].Call(params)
	}

	return nil
}