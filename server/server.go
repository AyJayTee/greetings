package server

import (
	"fmt"
	"net/http"

	"github.com/AyJayTee/greetings/lib"
)

func ServiceStart() {
	http.HandleFunc("/helloworld", helloWorld)
	http.HandleFunc("/", hello)

	http.ListenAndServe("localhost:8080", nil)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[1:]

	if name == "" {
		name = "World"
	}

	fmt.Fprintln(w, lib.Hello(name))
}
