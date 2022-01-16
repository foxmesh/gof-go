// Package observer 观察者模式(发布订阅模式)
// 定义：在对象之间定义一个一对多的依赖，当一个对象状态改变的时候，所有依赖的对象都会自动收到通知

package observer

// 基础实现
type IObserver interface {
	Update(msg string)
}

type ISubject interface {
	Register(observer IObserver)
	Remove(observer IObserver)
	Notify(msg string)
}

type Subject struct {
	observers []IObserver
}

func (s *Subject) Register(observer IObserver)  {
	s.observers = append(s.observers, observer)
}

func (s *Subject) Remove(observer IObserver)  {
	for i,v:=range s.observers{
		if v==observer{
		  s.observers = append(s.observers[:i], s.observers[i+1:]...)
		}
	}
}

func (s *Subject) Notify(msg string)  {
	for _,ob:=range s.observers{
		ob.Update(msg)
	}
}

