package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
	"go-simple/rabbitmq/rabbitmq_with_protobuf_client/proto/user"
)

func main() {
	//连接rabbitmq
	const host = "amqp://admin:admin@192.168.2.33:5672/"
	conn, err := amqp.Dial(host)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	channel, err := conn.Channel()
	if err != nil {
		fmt.Println("get channel false,error is", err)
	}
	defer channel.Close()
	queue, err := channel.QueueDeclare("protobuffer", false, false, false, false, nil)
	if err != nil {
		fmt.Println("declarequeue false,error is ", err)
	}

	//将实例化user的proto数据
	user := &user.User{
		Username: "scrazying",
		Password: "123456",
		Role:     user.UserRole_admin,
	}
	msg, err := proto.Marshal(user)
	if err != nil {
		fmt.Println("marshal protobuf data error", err)
	}
	//发送proto格式的数据
	err = channel.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text:json",
			Body:        []byte(msg),
		})
	if err != nil {
		fmt.Println("channel publish msg ,error is", err)
	}

}
