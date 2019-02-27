package main

import (
	"fmt"
	"time"
)

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func progress(delay time.Duration) {
	for {
		for _, v := range `-\|/` {
			fmt.Printf("%c", v)
			time.Sleep(delay)
		}
	}
}
func main() {
	go progress(100 * time.Millisecond)
	ans := fib(45)
	fmt.Printf("\nans is %d \n", ans)
}
