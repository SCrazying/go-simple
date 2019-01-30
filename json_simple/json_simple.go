package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Desc     string `json:"desc"`
}

type People struct {
	Persons []Person `json:"person"`
}

func main() {
	//常用变量解析
	num := 10
	numBuffer, err := json.Marshal(num)
	if err != nil {
		log.Fatalln("json Marshal err ----- ", err)
	}
	fmt.Println(string(numBuffer))
	//切片对象转为数组
	s := make([]int, 10, 10)
	s = append(s, 12)
	sliceBuffer, err := json.Marshal(s)
	if err != nil {
		log.Fatalln("json Marshal err ----- ", err)
	}
	fmt.Println(string(sliceBuffer))
	//map对象转为object
	mmap := make(map[int]string)
	mmap[10] = "sss"
	mmap[12] = "zzz"
	mapBuffer, err := json.Marshal(mmap)
	if err != nil {
		log.Fatalln("json Marshal err ----- ", err)
	}
	fmt.Println(string(mapBuffer))
	//struct 转为object
	var person = Person{10, "yyy", "test"}
	//Marshal 默认只支持public 变量，默认的值为变量字符串,使用`json:"id"`可以修改类型值 利用tag
	buffer, err := json.Marshal(person)
	if err != nil {
		log.Fatalln("json Marshal err ----- ", err)
	}
	fmt.Println(string(buffer))

	//list 转为 数组对象
	var people People
	people.Persons = append(people.Persons, Person{10, "xiaoming", "爱游戏"})
	people.Persons = append(people.Persons, Person{10, "xiaohong", "爱学习"})
	listBuffer, err := json.Marshal(people)
	if err != nil {
		log.Fatalln("json Marshal err ----- ", err)
	}
	fmt.Println(string(listBuffer))
	//利用已知类型进行json解析
	var rPeople People
	err = json.Unmarshal(listBuffer, &rPeople)
	if err != nil {
		log.Fatalln("json Unmarshal err ----- ", err)
	}
	fmt.Println(rPeople)
	//使用interface进行数据解析,JSON包中采用map[string]interface{}和[]interface{}结构来存储任意的JSON对象和数组
	//首先解析到[]interface{}中
	var listPerson interface{}
	err = json.Unmarshal(listBuffer, &listPerson)
	if err != nil {
		log.Fatalln("json Unmarshal err ----- ", err)
	}
	fmt.Println(listPerson)
	lPeople := listPerson.(map[string]interface{})
	for k, v := range lPeople {
		//fmt.Println(k)
		//fmt.Println(v)
		switch v_type := v.(type) {
		case string:
			break
		case int:
			break
		case interface{}:
			fmt.Println(k, "is an array:", v_type)
			//			for kk, vv := range v_type {

			//			}
			break
		}
		//		for vv := range v {
		//			switch vvv := vv.(type) {
		//			case string:
		//				break
		//			case int:
		//				break
		//			case interface{}:
		//				break
		//			default:
		//				fmt.Println(kk, vv, "is of a type I don't know how to handle")

		//			}
		//		}
	}

}
