package main

import (
	"context"
	"fmt"
	"go-simple/grpc/myexample"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"time"
)

type RouterServer struct {
}

func (r *RouterServer) C2S(ctx context.Context, request *myexample.RouteRequest) (*myexample.RouteReply, error) {

	reply := &myexample.RouteReply{}
	reply.C = request.A * request.B
	fmt.Println("call 点对点连接")
	return reply, nil
}

func (r *RouterServer) C2Sstream(request *myexample.RouteRequest, stream myexample.Router_C2SstreamServer) error {
	fmt.Println("call 服务器stream连接")
	for i := request.A; i < request.B; i++ {
		reply := &myexample.RouteReply{}
		reply.C = i * i
		stream.Send(reply)
		//模拟超时
		//time.Sleep(time.Second)
	}
	return nil
}

func (r *RouterServer) Cstream2S(stream myexample.Router_Cstream2SServer) error {

	fmt.Println("call 客户端stream连接")

	reply := &myexample.RouteReply{}
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		reply.C = reply.C + request.A + request.B
	}
	err := stream.SendMsg(reply)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (r *RouterServer) Cstream2Sstream(stream myexample.Router_Cstream2SstreamServer) error {

	fmt.Println("call 双向stream连接")

	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		reply := &myexample.RouterStream{Recv: "recv : " + request.Recv + " send : " + "test"}
		stream.Send(reply)

		time.Sleep(time.Second)
	}
	return nil
}
func main() {

	s := grpc.NewServer()

	myexample.RegisterRouterServer(s, &RouterServer{})

	listen, err := net.Listen("tcp", ":2333")
	if err != nil {
		log.Println(err)
	}
	defer listen.Close()
	fmt.Println("grpc server start")
	s.Serve(listen)

}
