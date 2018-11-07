package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println(v)
				o <- true
				break
			case <-time.After(5 * time.Second):
				fmt.Println("time out")
				o <- true
				break
			}
		}
	}()
	//接受到其他线程的时候响应
	go func() {
		time.Sleep(3 * time.Second)
		c <- 10
	}()
	<-o
}
