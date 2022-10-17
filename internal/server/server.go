package server

import (
	"fmt"
	"net/http"

	"github.com/AyJayTee/greetings"
)

func ServiceStart(port string) {
	http.HandleFunc("/helloworld", helloWorld)
	http.HandleFunc("/", hello)

	http.ListenAndServe(port, nil)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[1:]

	if name == "" {
		name = "World"
	}

	fmt.Fprintln(w, greetings.Hello(name))
}
