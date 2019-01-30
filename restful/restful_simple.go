package main

import (
	"net/http"
)

func Dologin(rsp http.ResponseWriter, rqs *http.Request) {
	rsp.Write([]byte("dologin"))
}

func Welcome(rsp http.ResponseWriter, rqs *http.Request) {
	rsp.Write([]byte("hello world"))
}
func main() {
	http.HandleFunc("/login", Dologin)
	http.HandleFunc("/", Welcome)
	http.ListenAndServe("127.0.0.1:9527", nil)
	//http.HandleFunc("\login",&Dologin())
}
