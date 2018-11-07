package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	//解析参数
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.Scheme)

	w.Write([]byte("sayHello"))

	for k, v := range r.Form {
		fmt.Print("key : " + k + " ")
		fmt.Println("val : ", strings.Join(v, ""))
	}

}
func main() {

	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9527", nil)
	if err != nil {
		log.Fatal("ListenAndServe :", err)
	}
}
