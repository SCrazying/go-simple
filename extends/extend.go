package main

import (
	"fmt"
)

type Shape struct {
	Name string
	desc string
}

type Rectangle struct {
	Shape
	width  int
	height int
}

type Round struct {
	Shape  //匿名函数支持继承后的方法重写
	radius float64
}

//方法继承
func (this Shape) PrintName() string {
	return this.Name
}

//重写函数
func (this Rectangle) PrintName() string {
	return "shape name is " + this.Name
}

func main() {
	shape := Shape{"shape", "this is a base shape"}
	fmt.Println(shape.PrintName())
	round := Round{Shape{"round", "this is a rectangle"}, 10}
	fmt.Println(round.PrintName())
	rectangle := Rectangle{Shape{"rectangle", "this is a rectangle"}, 10, 10}
	fmt.Println(rectangle.PrintName())
}
