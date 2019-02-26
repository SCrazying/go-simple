package main

import (
	"fmt"
	"time"
)

var bufferChannel = make(chan int, 10)

var write_chanel = make(chan bool)
var read_chanel = make(chan bool)

func write(max int) {

	for {
		for i := 0; i < max; i++ {
			bufferChannel <- i
			fmt.Printf("write value %d\n", i)
			time.Sleep(1 * time.Second)

		}
		write_chanel <- true
		break
	}

}

func read(max int) {
	for max > 0 {
		r := <-bufferChannel
		fmt.Printf("read value %d\n", r)
		time.Sleep(2 * time.Second)
		max--
	}
	read_chanel <- true
}
func waitfor() {
	for {
		read := <-read_chanel
		fmt.Println("read task has finished")
		write := <-write_chanel
		fmt.Println("write task has finished")
		if read && write {
			fmt.Println("task has finished")
			break
		}
	}

}
func readandwrite(max int) {

	go write(max)
	go read(max)
	go waitfor()
}
func main() {
	readandwrite(5)
	for {
	}
}
