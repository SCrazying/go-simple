package main

import (
	"fmt"
	"strconv"
)

type HumanInterface interface {
	PrintName()
	PrintAge()
}

type Human struct {
	name string
	age  int
}
type ManInterface interface {
	PrintPhone()
}
type Man struct {
	human Human
	phone string
}

type WomanInterface interface {
	Printlike()
}
type Woman struct {
	Human
	like string
}

func (this Human) PrintName() {
	fmt.Println(this.name)
}
func (this Human) PrintAge() {
	fmt.Println(this.age)
}
func (this Man) PrintName() {
	fmt.Println(this.human.name)
}
func (this Man) PrintAge() {
	fmt.Println(this.human.age)
}
func (this Man) PrintPhone() {
	fmt.Println(this.phone)
}

func (this Woman) PrintName() {
	fmt.Println(this.name)
}
func (this Woman) PrintAge() {
	fmt.Println(this.age)
}
func (this Woman) PrintLike() {
	fmt.Println(this.like)
}

//空接口支持任何类型的参数传递
func Show(this interface{}) {
	//this.PrintName()
}

//重新实现Stringer interface String()
func (this Human) String() string {
	age := strconv.Itoa(this.age)
	return "name is " + this.name + ",age is " + age
}

func main() {
	var human = Human{"human", 18}
	human.PrintName()
	human.PrintAge()

	man := Man{Human{"man", 18}, "123456789"}
	man.PrintName()
	man.PrintAge()
	man.PrintPhone()

	woman := Woman{Human{"woman", 18}, "like shopping"}
	woman.PrintName()
	woman.PrintAge()
	woman.PrintLike()

	//空接口支持任何类型的数据
	var ii interface{}
	val := 1
	val_s := "test"
	ii = val
	fmt.Println(ii)
	ii = val_s
	fmt.Println(ii)

	//输出Human数据
	fmt.Println(human)

}
