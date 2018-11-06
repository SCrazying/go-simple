package main

import (
	"fmt"
	"reflect"
)

type BookImp interface {
	PrintInfo() string
}

type Book struct {
	Name   string
	Widget string
}

type ChineseBook struct {
	Book
	Param1 string
}

type MathBook struct {
	Book
	Param2 string
}

func (this Book) PrintInfo() string {
	return "Name : " + this.Name + " Widget : " + this.Widget
}
func (this ChineseBook) PrintInfo() string {
	return "Name : " + this.Name + " Widget : " + this.Widget + " Param1 : " + this.Param1
}
func (this MathBook) PrintInfo() string {
	return "Name : " + this.Name + " Widget : " + this.Widget + " Param2 : " + this.Param2
}

type AnyThing interface {
}

type List []AnyThing

type Person struct {
	Name string
	Age  int
}

type PersonImpl interface {
	GetName() string
	GetAge() int
	SetName(v string)
	SetAge(a int)
	TestFun(v string)
	TestFunNotArgs()
}

func (this Person) GetName() string {
	return this.Name
}

func (this Person) GetAge() int {
	return this.Age
}

func (this Person) SetName(v string) {
	this.Name = v
}
func (this Person) SetRealName(v string) {
	this.Name = v
}

func (this Person) SetAge(a int) {
	this.Age = a
}
func (this Person) TestFun(v string) {
	fmt.Println(v)
}
func (this Person) TestFunNotArgs() {
	fmt.Println("not args")
}

type T struct {
	Age  int
	Name string
}

func main() {

	book := Book{"Book", "10"}
	chineseBook := ChineseBook{Book{"ChineseBook", "12"}, "param1"}
	mathBook := MathBook{Book{"MathBook", "13"}, "param2"}
	fmt.Println(book)
	fmt.Println(chineseBook)
	fmt.Println(mathBook)
	fmt.Println(book.PrintInfo())
	fmt.Println(chineseBook.PrintInfo())
	fmt.Println(mathBook.PrintInfo())

	var testBook BookImp
	testBook = book
	fmt.Println(testBook.PrintInfo())

	var testChineseBook BookImp
	testChineseBook = chineseBook
	fmt.Println(testChineseBook.PrintInfo())

	//类型判断
	list := make(List, 3)
	list[0] = book
	list[1] = chineseBook
	list[2] = mathBook

	for _, val := range list {
		switch val.(type) {
		case Book:
			fmt.Println("Book")
			break
		case ChineseBook:
			fmt.Println("ChineseBook")
			break
		case MathBook:
			fmt.Println("MathBook")
			break
		}
	}

	//反射
	var person interface{}
	person = Person{"name", 1222}

	var t = reflect.TypeOf(person)
	var v = reflect.ValueOf(person)
	fmt.Println(t)
	fmt.Println(v)

	//t0 := t.Elem().Field(0).Tag
	fmt.Println("------------")
	for i := 0; i < t.NumField(); i++ {
		tt := t.Field(i)
		vv := v.Field(i).Interface()
		fmt.Print(tt.Type)
		fmt.Printf("%s: %v = %v\n", tt.Name, tt.Type, vv)

	}
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}

	var data float64 = 1.23456789
	//要修改值必须是指针 不然 panic: reflect: reflect.Value.SetFloat using unaddressable value
	data_v := reflect.ValueOf(&data)
	//通过Elem获取原始值对应的对象
	newValue := data_v.Elem()
	newValue.SetFloat(9.87654321)
	fmt.Println(data)

	fmt.Println("------------")
	//调用Person方法

	reflect_vv := reflect.ValueOf(person)

	//带参数的方法调用
	TestFun := reflect_vv.MethodByName("TestFun")
	args := []reflect.Value{reflect.ValueOf("test")}
	TestFun.Call(args)
	//不带方法的调用
	TestFunNotArgs := reflect_vv.MethodByName("TestFunNotArgs")
	notargs := make([]reflect.Value, 0)
	TestFunNotArgs.Call(notargs)
	//通过方法修改结构体值

	pppp := Person{"name", 1222}
	pp := reflect.ValueOf(&pppp)
	pp.Elem().Field(0).SetString("dada")
	fmt.Println(pppp)

	pppp.SetName("xiaoming")
	fmt.Println(pppp)
	pppp.SetRealName("xiaohong")
	fmt.Println(pppp)
	pppp.Name = "hello"
	fmt.Println(pppp)
	elem := pp.Elem()
	SetName := elem.MethodByName("SetName")
	args = []reflect.Value{reflect.ValueOf("fuck")}
	SetName.Call(args)
	fmt.Println(pppp)

	var _name = pp.MethodByName("GetName").Call(make([]reflect.Value, 0))
	fmt.Println(_name)

	tttt := T{12, "someone-life"}
	ss := reflect.ValueOf(&tttt).Elem()

	ss.Field(0).SetInt(123)
	fmt.Println(ss)
}
