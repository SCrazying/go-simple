package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func doLogin(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		//主动校验form
		r.ParseForm()
		name := r.Form.Get("username")
		pwd := r.Form.Get("password")

		fmt.Println(name, pwd)
		w.Write([]byte("post"))
	} else {

		w.Write([]byte("dsadas"))
	}
}
func main() {

	http.HandleFunc("/login", doLogin)
	http.ListenAndServe(":8080", nil)
}
