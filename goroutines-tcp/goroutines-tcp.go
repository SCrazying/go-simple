package main

import (
	"fmt"
	"net"
)

func handerconnect(conn net.Conn) {
	defer conn.Close()
	for {
		buffer := make([]byte, 0)
		_, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("reader buffer error")
		}
		fmt.Println("recv from " + conn.RemoteAddr().String() + string(buffer))
	}
}
func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	fmt.Println("start listen localhost:8080")
	if err != nil {
		fmt.Println("net listen error")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept error")
			continue
		}
		handerconnect(conn)

	}
}
