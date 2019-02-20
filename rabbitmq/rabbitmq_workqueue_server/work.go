package main

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

func failError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func doworkthread1() {
	conn, err := amqp.Dial("amqp://admin:admin@192.168.0.16:5672/")
	failError(err)

	defer conn.Close()

	ch, err := conn.Channel()

	failError(err)

	defer ch.Close()
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failError(err)

	q, err := ch.QueueDeclare(
		"work_queue",
		true,
		false,
		false,
		false,
		nil)

	failError(err)
	wait := make(chan bool)
	msg, err := ch.Consume(q.Name,
		"",
		false, //false 不会主动应答
		false,
		false,
		false,
		nil)
	failError(err)

	go func() {

		for d := range msg {

			//fmt.Println("doworkthread1 begin")
			time.Sleep(1 * time.Second)
			d.Ack(false)
			fmt.Println("doworkthread1 done")
		}

	}()
	<-wait
}

func doworkthread2() {
	conn, err := amqp.Dial("amqp://admin:admin@192.168.0.16:5672/")
	failError(err)

	defer conn.Close()

	ch, err := conn.Channel()

	failError(err)

	defer ch.Close()

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failError(err)

	q, err := ch.QueueDeclare(
		"work_queue",
		true,
		false,
		false,
		false,
		nil)

	failError(err)
	wait := make(chan bool)
	msg, err := ch.Consume(q.Name,
		"",
		false, //false 不会主动应答
		false,
		false,
		false,
		nil)
	failError(err)

	go func() {

		for d := range msg {

			//fmt.Println("doworkthread2 start")
			time.Sleep(3 * time.Second)
			d.Ack(false)
			fmt.Println("doworkthread2 done")
		}

	}()
	<-wait
}
func main() {

	wait := make(chan bool)

	go doworkthread1()
	go doworkthread2()

	<-wait
}
