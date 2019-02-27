package main

import (
	"fmt"

	"os"
)

func main() {
	fmt.Printf("%d %s %T\n", 100, "dsadsadsa", "dsadsadsa")
	data := fmt.Sprintf("%d %s \n", 10, "test")
	fmt.Print(data)

	lenn, err := fmt.Fprint(os.Stdout, "da ", 11)
	if err != nil {
		fmt.Println("Fprint error")
	}
	fmt.Println(lenn)

	len, errr := fmt.Fprintln(os.Stdout, "da ", 11)
	if errr != nil {
		fmt.Println("Fprint error")
	}
	fmt.Println(len)

}
