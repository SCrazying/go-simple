package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Person struct {
	Name string
	Age  int
}

func (this Person) PutString() string {
	return "Person name is " + this.Name + ",age is" + strconv.Itoa(this.Age)
}

func FormatPerson(v interface{}) string {
	var this = reflect.TypeOf(v)
	fmt.Println(this.Name())
	return ""
	//return this.PutString()
}

func main() {
	person := Person{"name", 1222}
	FormatPerson(person)
}
