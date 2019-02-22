package main

import (
	"log"

	"github.com/streadway/amqp"
)

func failError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
func main() {

	conn, err := amqp.Dial("amqp://admin:admin@192.168.0.16:5672/")

	failError(err)
	defer conn.Close()

	ch, err := conn.Channel()
	failError(err)

	defer ch.Close()

	err = ch.ExchangeDeclare(
		"hello_exchange",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	failError(err)

	q, err := ch.QueueDeclare(
		"test",
		true,
		false,
		false,
		false,
		nil,
	)
	failError(err)
	err = ch.QueueBind(q.Name, "test_key", "hello_exchange", false, nil)

	failError(err)

	ch.Publish(
		"hello_exchange",
		"test_key",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("test_exchange"),
			AppId:       "test_yyy",
		},
	)
}
