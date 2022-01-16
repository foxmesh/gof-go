// Package flyweight 享元模式
// 目的：共享对象，以节省内存，前提是被共享的对象是不可变的
// 实现方式：主要通过工厂模式，在工厂类中，通过一个Map类来存储已创建的享元对象，以达到复用的目的

package flyweight

// ChessPieceUnit 象棋棋子是不可变的，所以可以享元
type ChessPieceUnit struct {
	ID int
	Name string
	Color string
}

var units = map[int]*ChessPieceUnit{
	1:{ID: 1,Name: "炮",Color: "red",
	},
	// ...
}

func NewChessPieceUnit(id int) *ChessPieceUnit {
	return units[id]
}
