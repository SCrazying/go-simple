package main

import (
	"context"
	"fmt"
	"go-simple/grpc/myexample"
	"google.golang.org/grpc"
	"io"
	"log"
)

const host = "127.0.0.1:2333"

func main() {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		fmt.Println("connnect failed", err)
	}
	defer conn.Close()

	cli := myexample.NewRouterClient(conn)
	request := &myexample.RouteRequest{A: 10, B: 20}
	reply, err := cli.C2S(context.Background(), request)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%d * %d = %d\n", request.A, request.B, reply.C)

	fmt.Println("---------------------------------")

	reply_stream, err := cli.C2Sstream(context.Background(), request)

	if err != nil {

	}
	for {
		//阻塞式等待应答，同步等待调用rpc
		//使用 sync.Pool{}对象池在一个线程中处理连接过后的对象数据  不阻塞主线程
		reply, err = reply_stream.Recv()
		if err == io.EOF {
			//接受数据完成
			break
		}
		fmt.Printf("ans = %d\n", reply.C)
	}

	fmt.Println("---------------------------------")

	var requests = make([]*myexample.RouteRequest, 0)
	requests = append(requests, &myexample.RouteRequest{A: 1, B: 2})
	requests = append(requests, &myexample.RouteRequest{A: 3, B: 4})
	requests = append(requests, &myexample.RouteRequest{A: 5, B: 6})

	stream, err := cli.Cstream2S(context.Background())

	for _, req := range requests {
		stream.SendMsg(req)
	}

	reply, err = stream.CloseAndRecv()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("ans = %d\n", reply.C)

	fmt.Println("---------------------------------")

	stream1, err := cli.Cstream2Sstream(context.Background())

	for _, req := range requests {
		stream1.SendMsg(req)
	}
	stream1.CloseSend()
	for {
		replystream, err := stream1.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}

		fmt.Printf(replystream.Recv)
	}

}
