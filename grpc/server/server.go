package main

import (
	"context"
	"fmt"
	"go-simple/grpc/myexample"
	"net"

	"google.golang.org/grpc"
)

const port = ":43211"

type LoginServer struct {
}

func (l *LoginServer) DoLogin(ctx context.Context, r *myexample.LoginRequest) (*myexample.LoginReply, error) {
	return &myexample.LoginReply{Buffer: "test buffer"}, nil
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("list ", port, err)
	}
	s := grpc.NewServer()
	myexample.RegisterLoginServer(s, &LoginServer{})
	s.Serve(lis)
}
