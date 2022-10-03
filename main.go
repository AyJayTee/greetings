package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/helloworld", helloWorld)
	http.HandleFunc("/", hello)

	http.ListenAndServe(":8080", nil)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[1:]

	if name != "" {
		fmt.Fprintf(w, "Hello %s! Nice to meet you!", name)
		return
	}

	fmt.Fprintln(w, "Hello! Nice to meet you!")
}
