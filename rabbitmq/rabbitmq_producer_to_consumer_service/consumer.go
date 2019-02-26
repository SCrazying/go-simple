package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func failError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
func main() {

	conn, err := amqp.Dial("amqp://admin:admin@192.168.0.16:5672")

	failError(err)
	defer conn.Close()

	ch, err := conn.Channel()
	failError(err)
	defer ch.Close()

	msg, err := ch.Consume(
		"test",
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	failError(err)

	wait := make(chan bool)
	go func() {

		for d := range msg {
			fmt.Println(string(d.Body))
			d.Ack(true)
		}
	}()
	<-wait
}
