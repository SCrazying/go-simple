package main

import (
	"fmt"
	"strconv"
)

func main() {

	var val interface{}
	//string to int
	val, err := strconv.Atoi("10")
	if err != nil {
		fmt.Println("strconmv atoi error")
	}
	fmt.Print(val)
	fmt.Printf(" %T\n", val)
	//string to int64
	val, _ = strconv.ParseInt("100001", 10, 64)
	fmt.Print(val)
	fmt.Printf(" %T\n", val)
	//int to string
	val = strconv.Itoa(100)
	fmt.Print(val)
	fmt.Printf(" %T\n", val)
	//int64 to String
	val = strconv.FormatInt(100001, 10)
	fmt.Print(val)
	fmt.Printf(" %T\n", val)
	//int64 to int ,to long catch error
	val = int(111111111111111111)
	fmt.Print(val)
	fmt.Printf(" %T\n", val)
	//int to int64
	val = int64(100001)
	fmt.Print(val)
	fmt.Printf(" %T\n", val)
	//interface to any type,val type is int64
	value64, ok := val.(int64)
	if !ok {
		fmt.Println("convert error")
	}
	fmt.Print(value64)
	fmt.Printf(" %T\n", value64)

	//interface类型和转换类型不对报错
	_, okk := val.(int)
	if !okk {
		fmt.Println("convert error")
	}

	var temp interface{}
	temp = value64
	fmt.Print(temp)
	fmt.Printf(" %T\n", temp)
	//string to byte
	var strbuff = "哈喽 go"
	bytes := []byte(strbuff)
	fmt.Println(bytes)
	sss := fmt.Sprintf("%s", bytes)
	fmt.Println(sss)
	strs := string(bytes)
	fmt.Println(strs)
}
