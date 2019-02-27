package main

import (
	"fmt"
	"regexp"
)

func main() {
	flag, err := regexp.MatchString("^*xml.go$", "/e/code/src/go-simple/xml.go")
	if err != nil {
		fmt.Println("regexp string error")
	}
	if flag {
		fmt.Println("match true")
	}

}
