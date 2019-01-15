package main

import (
	"fmt"
	"go-simple/protobuf/login"

	"github.com/golang/protobuf/proto"
)

func main() {
	msg_request := &login.LoginRequest{
		Username: "小明",
		Password: "1234",
		Type:     login.TYPE_student,
	}

	in_data, err := proto.Marshal(msg_request)
	if err != nil {
		fmt.Println("Marshal err :", err)
	}
	msg_encoding := &login.LoginRequest{}
	err = proto.Unmarshal(in_data, msg_encoding)
	if err != nil {
		fmt.Println("Unmarshaling error: ", err)
	}
	fmt.Printf("msg username: %s\n", msg_encoding.GetUsername())
	fmt.Printf("msg password: %s\n", msg_encoding.GetPassword())
	fmt.Printf("msg type: %d\n", msg_encoding.GetType())
}
