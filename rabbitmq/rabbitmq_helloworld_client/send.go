package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s, %s", msg, err)

	}
}

func main() {
	conn, err := amqp.Dial("amqp://admin:admin@192.168.0.16:5672/")

	failOnError(err, "Do not connect RabbitMQ, err is ")

	defer conn.Close()

	ch, err := conn.Channel()

	failOnError(err, "Create Channel false,err is ")

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello_world",
		false,
		false,
		false,
		false,
		nil)

	fmt.Println(q.Name)

	body := "hello_world"
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "publish false")
}
