package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"go-simple/micro/hello_micro/handler"
	"go-simple/micro/hello_micro/subscriber"

	example "go-simple/micro/hello_micro/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("com.skyfin.srv.hello_micro"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("com.skyfin.srv.hello_micro", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("com.skyfin.srv.hello_micro", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
