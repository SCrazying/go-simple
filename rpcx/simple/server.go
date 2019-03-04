package main

import (
	"context"
	"github.com/smallnest/rpcx/server"
)

type Args struct {
	A int
	B int
}
type Reply struct {
	C int
}

type Arith struct {
}

func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	//模拟超时
	//time.Sleep(time.Second * 10)
	return nil
}
func main() {

	s := server.NewServer()
	//s.RegisterName("Arith", &Arith{}, "")
	s.Register(&Arith{}, "")
	s.Serve("tcp", ":2333")
}
