package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	http.HandleFunc("/helloworld", helloWorld)
	http.HandleFunc("/", hello)

	http.ListenAndServe(":8080", nil)
}

func Hello(name string) string {
	if name == "" {
		return ""
	}

	return fmt.Sprintf(message(), name)
}

func message() string {
	formats := []string{"Hello %s!", "Welcome %s!", "Hi %s!"}

	index := rand.Intn(len(formats))

	return formats[index]
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[1:]

	if name == "" {
		name = "World"
	}

	fmt.Fprintln(w, Hello(name))
}
