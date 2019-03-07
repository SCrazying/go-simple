package main

import (
	pb "go-simple/grpc/myexample"
	"google.golang.org/grpc"

	"log"
	"net"
)

//这个服务用于来自客户端的num ，并且返回递归返回结果
type FibonacciServer struct{}

//使用闭包的形式来计算fibonacci
func Fibonacci() func() int32 {

	//缓存数据
	var sum int32 = 0
	var a int32 = 0
	var b int32 = 1
	return func() int32 {
		sum = a + b
		a = b
		b = sum
		return sum
	}
}

func (s *FibonacciServer) CalFibonacci(request *pb.FibonacciRequest, stream pb.Fibonacci_CalFibonacciServer) error {

	var total = int32(request.Num)
	f := Fibonacci()
	for i := int32(0); i <= total; i++ {
		if i == 0 || i == 1 {
			stream.SendMsg(&pb.FibonacciReply{Ans: i})
		} else {
			stream.SendMsg(&pb.FibonacciReply{Ans: f()})
		}

	}
	return nil
}

const Host = ":2333"

func main() {
	//注册gRpc服务
	s := grpc.NewServer()
	//服务监听
	ser := &FibonacciServer{}
	pb.RegisterFibonacciServer(s, ser)
	//创建监听端口
	l, err := net.Listen("tcp", Host)
	if err != nil {
		log.Println(err)
	}
	//将rpc端口绑定到tcp 2333端口上
	s.Serve(l)
}
