package main

import (
	"fmt"
	pb "go-simple/grpc/myexample"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

type CharServer struct{}

var lock sync.Mutex

func (s *CharServer) OnChat(stream pb.Chat_OnChatServer) error {
	lock.Lock()
	defer lock.Unlock()
	ctx := stream.Context()
	select {
	case <-ctx.Done():
		//处理异常关闭信号
		fmt.Println("server connext close")
		return ctx.Err()
	default:
		fmt.Println("server start")
		wait := make(chan bool)
		go func() {
			for {
				req, err := stream.Recv()
				if err != nil {
					break
				}
				reply := &pb.ChatReply{}
				//数据回显
				reply.Reply = "reply : " + req.Data
				err = stream.Send(reply)
				if err != nil {
					log.Fatalln(err)
				}
			}
			wait <- true
		}()
		<-wait
	}
	fmt.Println("server finish")
	return nil
}
func main() {
	const Host = ":2333"
	s := grpc.NewServer()
	service := &CharServer{}

	pb.RegisterChatServer(s, service)

	l, err := net.Listen("tcp", Host)
	if err != nil {
		log.Fatalln(err)
	}
	s.Serve(l)
}
