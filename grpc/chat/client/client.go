package main

import (
	"context"
	"fmt"
	pb "go-simple/grpc/myexample"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	const Host = ":2333"
	conn, err := grpc.Dial(Host, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	chatClient := pb.NewChatClient(conn)
	client, err := chatClient.OnChat(context.Background())

	go func() {
		for i := 0; i < 10; i++ {
			err = client.Send(&pb.ChatRequest{Data: "test"})
			if err != nil {
				log.Fatalln(err)
			}
			time.Sleep(time.Second)
		}
		client.CloseSend()
	}()

	wait := make(chan bool)
	go func() {
		for {
			reply, err := client.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln(err)
				break
			}
			fmt.Println(reply.Reply)
		}
		wait <- true
		//
		fmt.Println("接受数据完成")
	}()
	<-wait
}
