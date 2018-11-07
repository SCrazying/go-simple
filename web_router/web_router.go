package main

import (
	"net/http"
)

type MyMux struct {
}

func (this *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayHello(w, r)
		return
	}
	http.NotFound(w, r)
	return
}
func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("sayHello"))
}
func main() {
	mux := &MyMux{}
	//http.HandleFunc("/", mux)
	http.ListenAndServe(":8080", mux)
}
