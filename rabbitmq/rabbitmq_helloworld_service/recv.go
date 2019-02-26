package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func failError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s, %s", msg, err)
	}
}
func main() {

	conn, err := amqp.Dial("amqp://admin:admin@192.168.0.16:5672/")
	failError(err, "")
	defer conn.Close()

	ch, err := conn.Channel()
	failError(err, "")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello_world",
		false,
		false,
		false,
		false,
		nil,
	)

	failError(err, "")

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)

	failError(err, "")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Println("recv msg", string(d.Body))
		}
	}()

	<-forever
}
