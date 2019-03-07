package main

import (
	"context"
	"fmt"
	pb "go-simple/grpc/myexample"
	"google.golang.org/grpc"
	"io"
	"log"
)

const Host = ":2333"

func main() {

	conn, err := grpc.Dial(Host, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	cl := pb.NewFibonacciClient(conn)
	stream, err := cl.CalFibonacci(context.Background(), &pb.FibonacciRequest{Num: 10})
	if err != nil {
		log.Fatalln(err)
	}
	for {
		reply, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("count is %d\n", reply.Ans)
	}

}
