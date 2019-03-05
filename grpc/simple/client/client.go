package main

import (
	"context"
	"go-simple/grpc/myexample"
	"log"

	"google.golang.org/grpc"
)

const port = ":43211"

const (
	address = "localhost:43211"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()
	c := myexample.NewLoginClient(conn)
	r, err := c.DoLogin(context.Background(), &myexample.LoginRequest{
		Username: "xiaoming",
		Password: "xiaoming",
	})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	log.Printf("%s", r.GetBuffer())

}
