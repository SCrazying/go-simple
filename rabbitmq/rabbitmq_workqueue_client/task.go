package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func failError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
func main() {

	conn, err := amqp.Dial("amqp://admin:admin@192.168.0.16:5672/")
	failError(err)

	defer conn.Close()

	ch, err := conn.Channel()

	failError(err)

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"work_queue",
		true,
		false,
		false,
		false,
		nil)

	failError(err)

	wait := make(chan bool)
	for i := 0; i < 10; i++ {
		ch.Publish("",
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte("work test"),
			})
	}

	<-wait

}
