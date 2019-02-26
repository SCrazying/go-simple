package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)
var thread1exit = make(chan bool)
var thread2exit = make(chan bool)

func thread1() {
	fmt.Println("thread begin")
	for i := 0; i < 20; i++ {

		if i == 10 {
			ch <- 10
			break
		}
		fmt.Println("thread1 running")
		time.Sleep(time.Millisecond * 100)
	}
	fmt.Println("thread finish")
	thread1exit <- true

}
func thread2() {
	fmt.Println("thread2 begin")
	for {
		ans := <-ch
		if ans == 10 {
			break
		}
		fmt.Println("wait thread1")
	}
	fmt.Println("thread2 finish")
	thread2exit <- true
}

func main() {
	fmt.Println("main goroutines start")
	go thread1()
	go thread2()
	for {
		thread1ans := <-thread1exit
		thread2ans := <-thread2exit
		if thread1ans && thread2ans {
			break
		}
	}
	fmt.Println("main goroutines exit")
}
