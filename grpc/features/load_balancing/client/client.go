package main

import (
	"context"
	"fmt"
	pb "go-simple/grpc/features/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

var (
	addrs = []string{":2333", ":2334", ":2335", ":2336"}
)

type featuresServer struct{}

func (s *featuresServer) Dosomething(ctx context.Context, request *pb.Request) (*pb.Reply, error) {
	reply := &pb.Reply{}
	return reply, nil
}
func startServer(addr string) error {
	s := grpc.NewServer()
	pb.RegisterFeaturesServer(s, &featuresServer{})

	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	err = s.Serve(l)
	if err != nil {
		return err
	}
	fmt.Println(addr, " start")
	return nil
}
func main() {

	var wg sync.WaitGroup
	for _, addr := range addrs {
		wg.Add(1)
		go func(string) {
			defer wg.Done()
			err := startServer(addr)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(addr + " exit")
		}(addr)
	}
	wg.Wait()
}
