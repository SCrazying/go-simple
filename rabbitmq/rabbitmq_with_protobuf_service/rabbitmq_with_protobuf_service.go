package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
	user2 "go-simple/rabbitmq/rabbitmq_with_protobuf_client/proto/user"
)

func main() {
	//连接rabbitmq
	const host = "amqp://admin:admin@192.168.2.33:5672/"
	conn, err := amqp.Dial(host)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	fmt.Println("service start")
	channel, err := conn.Channel()
	if err != nil {
		fmt.Println("get channel false,error is", err)
	}
	defer channel.Close()
	queue, err := channel.QueueDeclare("protobuffer", false, false, false, false, nil)
	if err != nil {
		fmt.Println("declarequeue false,error is ", err)
	}
	msgs, err := channel.Consume(queue.Name, "", false, false, false, false, nil)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			user := &user2.User{}
			err := proto.Unmarshal(d.Body, user)
			if err != nil {
				fmt.Println("Unmarshal protobuf false,err is ", err)
			} else {
				fmt.Println("username : ", user.Username)
				fmt.Println("password : ", user.Password)
				fmt.Println("role : ", user.Role)
			}
			d.Ack(true)
		}
	}()

	<-forever
}
