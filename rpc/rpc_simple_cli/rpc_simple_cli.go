package main

import (
	"encoding/json"
	"fmt"
	"net/rpc"
)

type MessageLoginRequest struct {
	UserName string
	Password string
}

type MessageLoginReply struct {
	StatusCode int
	Msg        string
}

func (reply *MessageLoginReply) ToString(args ...interface{}) string {

	datas, _ := json.Marshal(reply)
	return string(datas)
}
func main() {

	addr := "192.168.0.16:2333"
	client, err := rpc.DialHTTP("tcp", addr)
	if err != nil {

		fmt.Println(err)
	}
	var reply MessageLoginReply
	err = client.Call("Service.Dologin", &MessageLoginRequest{
		UserName: "xiaoming",
		Password: "xiaoming",
	}, &reply)
	if err != nil {
		fmt.Println("rpc call faided,error is ", err)
	} else {
		fmt.Println(reply)
	}

}
