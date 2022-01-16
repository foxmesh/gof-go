// Package builder 建造者模式用来处理复杂对象的构造，在go 中，一般将必填参数直接传递，可选参数通过传递可变方法进行创建，如 WithOption
// 应用场景：
//   1、类型中属性比较多
//   2、结构中属性之间有一定的依赖关系，或者是约束条件
//   3、存在必选和非必选的属性
//   4、希望创建不可变的对象
// 和工厂模式的区别：
//   1、工厂模式：用于创建类型相关的不同对象
//   2、建造者模式：用于创建参数复杂的对象
package builder

import "errors"

type ResourcePoolConfig struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

type ResourcePoolConfigOption struct {
	maxTotal int
	maxIdle  int
	minIdle  int
}

type ResourcePoolConfigOptFun func(option *ResourcePoolConfigOption)

func NewResourcePoolConfig(name string, opts ...ResourcePoolConfigOptFun) (*ResourcePoolConfig, error) {
	if name == "" {
		return nil, errors.New("name can not be empty")
	}

	option := &ResourcePoolConfigOption{
		maxTotal: 10,
		maxIdle:  9,
		minIdle:  1,
	}

	for _, opt := range opts {
		opt(option)
	}

	return &ResourcePoolConfig{
		name:     name,
		maxTotal: option.maxTotal,
		maxIdle:  option.maxIdle,
		minIdle:  option.minIdle,
	}, nil
}
