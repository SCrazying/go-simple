package main

import (
	"context"
	"fmt"
	pb "go-simple/grpc/myexample"
	"log"

	"google.golang.org/grpc"
)

const Host = ":2333"

func main() {
	//建立rpc连接
	conn, err := grpc.Dial(Host, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	//使用rpc客户端
	c := pb.NewArithClient(conn)
	//同步阻塞调用
	reply, err := c.CalCircumference(context.Background(), &pb.ArithRequest{Width: 10, Height: 24})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Circumference = %d\n", reply.Circumference)
}
