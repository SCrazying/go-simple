package main

import (
	"context"
	"fmt"
	pb "go-simple/grpc/myexample"
	"google.golang.org/grpc"
	"log"
)

const Host = ":2333"

func main() {

	//连接服务器，使用加密的形式
	conn, err := grpc.Dial(Host, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	//创建proto定义的客户端
	cl := pb.NewCountClient(conn)
	stream, err := cl.CalTotal(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	//使用流的形式发送数据
	stream.Send(&pb.CountRequest{Num: 1})
	stream.Send(&pb.CountRequest{Num: 2})
	stream.Send(&pb.CountRequest{Num: 3})
	stream.Send(&pb.CountRequest{Num: 4})
	stream.Send(&pb.CountRequest{Num: 5})
	stream.CloseSend()

	reply := &pb.CountReply{}
	err = stream.RecvMsg(reply)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("count is %d\n", reply.Total)

}
