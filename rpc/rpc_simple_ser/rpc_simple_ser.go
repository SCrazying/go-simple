package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
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

type Service struct {
}

func (s *Service) Dologin(request *MessageLoginRequest, reply *MessageLoginReply) error {

	fmt.Println("dologin")
	if request.UserName == "xiaoming" && request.Password == "xiaoming" {
		reply.StatusCode = 0
		reply.Msg = " login success"

		return nil
	}
	reply.StatusCode = -1
	reply.Msg = "login faild"
	return nil
}

func main() {
	err := rpc.Register(&Service{})
	if err != nil {
		log.Fatalln("register service faided,err is ", err)
	}
	rpc.HandleHTTP()
	listen, err := net.Listen("tcp", "192.168.0.16:2333")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	go http.Serve(listen, nil)
	fmt.Println("listen 192.168.0.16:2333 start")
	wait := make(chan bool)
	<-wait
}
