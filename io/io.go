package main

import (
	"fmt"
	"io"
	"os"
)

func ReadStdin(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

//read bytes from file
func ReadFile(path string) {
	var file io.File

}
func main() {
	data, err := ReadStdin(os.Stdin, 11)
	if err == nil {
		fmt.Println(data)

	} else {
		fmt.Println("error ")
	}

}
