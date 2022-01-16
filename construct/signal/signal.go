package signal

import "sync"

type Object struct {
}

var (
	instance *Object
	once     sync.Once
)

// 饿汉模式
func init() {
	instance = &Object{}
}

func GetInstanceWithHunger() *Object {
	return instance
}

// GetInstallWithLazy 懒汉模式
func GetInstallWithLazy() *Object {
	if instance == nil {
		once.Do(func() {
			instance = new(Object)
		})
	}
	return instance
}
