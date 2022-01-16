// Package decorator 装饰器模式
// 功能：主要解决继承关系过于复杂的问题，通过组合来替代继承。它主要的作用是给原始类添加增强功能
// 实现方式:1)继承：如果装饰器类的方法大部分和源类相同，不做改变选继承
//         2)接口
// 和代理模式的区别：代码形式上几乎没什么差别
//     代理模式：主要是给原类添加无关的功能
//     装饰器模式：主要是给原类增强功能，添加的功能都是有关联的
package decorator

import "testing"

// 画画的例子，默认的 Square  只有基础的画画功能， ColorSquare  为他加上了颜色

type IDraw interface {
	Draw() string
}

type Square struct {
}

func (s Square) Draw() string {
	return "this is a square"
}

type ColorSquare struct {
	square IDraw
	color  string
}

func NewColorSquare(square IDraw, color string) ColorSquare {
	return ColorSquare{
		square: square,
		color:  color,
	}
}
func (s ColorSquare) Draw() string {
	return s.square.Draw() + ", color is " + s.color
}

func TestColorSquare_Draw(t *testing.T) {
	sq := Square{}
	csq := NewColorSquare(sq, "red")
	got := csq.Draw()

	if got != "this is a square, color is red" {
		t.Error(got)
	}
}
