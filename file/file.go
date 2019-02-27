package main

import (
	"fmt"
	"os"

	"io/ioutil"
	"log"
)

func FatalError(hint string, err error) {
	if err != nil {
		log.Fatal(fmt.Sprintf("error : %s", hint), err)
	}
}
func main() {
	//读取文件
	bytes, err := ioutil.ReadFile("1.txt")
	FatalError("ReadFile ", err)
	fmt.Println(string(bytes))
	fmt.Println(`------------------------------------------------`)
	//写入文件
	var buffer = "hello go"
	err = ioutil.WriteFile("3.txt", []byte(buffer), 0000)
	FatalError("ioutil.WriteFile", err)
	fmt.Println(`------------------------------------------------`)
	//使用os去read文件
	f, err := os.Open("1.txt")
	FatalError("os Open ", err)
	defer f.Close()
	bytes, err = ioutil.ReadAll(f)
	FatalError("ioutil readall ", err)
	fmt.Println(string(bytes))

	//使用os的readfile读取文件
	f, err = os.OpenFile("3.txt", os.O_RDONLY|os.O_CREATE, 0111)
	FatalError("os.OpenFile", err)

	defer f.Close()
	// bytes, err = ioutil.ReadAll(f)
	// size, err := f.Write([]byte(buffer))
	// FatalError("f.Write", err)
	// fmt.Println("write bytes size", size)

	//获取系统配置下的环境变量
	fmt.Println(os.Getenv("GOPATH"))
	//系统path 分隔符 windows \\ linux /
	fmt.Println(os.PathSeparator)

}
