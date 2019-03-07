package main

import (
	pb "go-simple/grpc/myexample"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

//这个服务用于统计来自客户端流数据的和，并返回到客户端
type CountServer struct{}

func (s *CountServer) CalTotal(stream pb.Count_CalTotalServer) error {

	//定义返回数据的结构
	reply := &pb.CountReply{}
	for {
		//循环接受数据
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		reply.Total += req.Num
	}
	//返回数据到客户端
	stream.SendAndClose(reply)
	return nil
}

const Host = ":2333"

func main() {
	//注册gRpc服务
	s := grpc.NewServer()
	//服务监听
	ser := &CountServer{}
	pb.RegisterCountServer(s, ser)
	//创建监听端口
	l, err := net.Listen("tcp", Host)
	if err != nil {
		log.Println(err)
	}
	//将rpc端口绑定到tcp 2333端口上
	s.Serve(l)
}
