package main

import (
	"context"
	"github.com/smallnest/rpcx/server"
	Arith2 "go-simple/rpcx/proto/Arith"
)

type Arith struct {
}

func (A *Arith) Mul(ctx context.Context, request *Arith2.ArithRequest, reply *Arith2.ArithReply) error {
	reply.C = request.A * request.B
	return nil
}
func main() {

	s := server.NewServer()
	s.Register(&Arith{}, "")
	s.Serve("tcp", ":2333")
}
