package main

import (
	"context"
	pb "go-simple/grpc/myexample"
	"google.golang.org/grpc"
	"log"
	"net"
)

type ArithServer struct {
}

func (service *ArithServer) CalCircumference(ctx context.Context, req *pb.ArithRequest) (*pb.Arithreply, error) {
	reply := &pb.Arithreply{}
	reply.Circumference = (req.Width + req.Width) * 2
	return reply, nil
}

const Host = ":2333"

func main() {
	//注册gRpc服务
	s := grpc.NewServer()
	//服务监听
	ser := &ArithServer{}
	pb.RegisterArithServer(s, ser)
	//创建监听端口
	l, err := net.Listen("tcp", Host)
	if err != nil {
		log.Println(err)
	}
	//将rpc端口绑定到tcp 2333端口上
	s.Serve(l)
}
