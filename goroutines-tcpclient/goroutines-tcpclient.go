package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net conn 127.0.0.1:8000")
	}
	for {
		bytes := make([]byte, 0)
		_, err := os.Stdin.Read(bytes)
		conn.Write()
	}
}
