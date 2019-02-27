package main

import (
	"errors"
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"log"
	"time"
)

func callback(evt zk.Event) {
	fmt.Println(">>>>>>>>>>>>>>>>")
	fmt.Println("path", evt.Path)
	fmt.Println("type", evt.Type.String())
	fmt.Println("state", evt.State.String())
	fmt.Println("<<<<<<<<<<<<<<<<")
}

//global listener
func globalListen() {

	//host info
	var hosts = []string{"192.168.0.16:2181"}

	// event callback
	option := zk.WithEventCallback(callback)
	conn, connChan, err := zk.Connect(hosts, time.Second*50, option)
	if err != nil {
		log.Fatalln("---> ", err)
	}
	defer conn.Close()
	for {
		isConnnect := false
		select {

		case connEvent := <-connChan:
			if connEvent.State == zk.StateConnected {
				isConnnect = true
			}
		case _ = <-time.After(time.Second * 3):
			errors.New("connect zk server timeout")
			break

		}
		if isConnnect {
			break
		}
	}
	//use function ExistsW  watch /test,just only once
	isExist, _, e, err := conn.ExistsW("/test")
	if isExist {
		select {
		case ch_event := <-e:
			if ch_event.Type == zk.EventNodeCreated {
				fmt.Println("zk.EventNodeCreated")
			} else if ch_event.Type == zk.EventNodeDeleted {
				fmt.Println("zk.EventNodeDeleted")
			} else if ch_event.Type == zk.EventNodeDataChanged {
				fmt.Println("zk.EventNodeDataChanged")
			}
		}
	} else {
		fmt.Println("/test  not exist")
	}
}

//part listener
func partListen() {
	var hosts = []string{"192.168.0.16:2181"}

	conn, _, err := zk.Connect(hosts, time.Second)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	isExist, _, event, err := conn.ExistsW("/test")
	wait := make(chan bool)
	if isExist {

		go func(<-chan zk.Event) {
			evt := <-event
			fmt.Println(">>>>>>>>>>>>>>>>")
			fmt.Println("path", evt.Path)
			fmt.Println("type", evt.Type.String())
			fmt.Println("state", evt.State.String())
			fmt.Println("<<<<<<<<<<<<<<<<")
			wait <- true
		}(event)
	} else {
		fmt.Println("/test is not exist")
		wait <- true
	}
	<-wait
}

//loop listener

func loop(conn *zk.Conn, path string) error {

	isExist, _, event, err := conn.ExistsW(path)
	if err != nil {
		return err
	}
	wait := make(chan bool)
	if isExist {
		evt := <-event
		fmt.Println(">>>>>>>>>>>>>>>>")
		fmt.Println("path", evt.Path)
		fmt.Println("type", evt.Type.String())
		fmt.Println("state", evt.State.String())
		fmt.Println("<<<<<<<<<<<<<<<<")
		go loop(conn, path)
		wait <- true

	} else {
		fmt.Println("/test is not exist")
		wait <- true
	}
	<-wait
	return nil
}
func loopListen() {
	var hosts = []string{"192.168.0.16:2181"}

	conn, _, err := zk.Connect(hosts, time.Second)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go loop(conn, "/test")
	wait := make(chan bool)
	<-wait
}
func main() {
	loopListen()
}
