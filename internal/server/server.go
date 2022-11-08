package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AyJayTee/greetings"
)

type StoreURL struct {
	Url string `json:"url"`
}

type ReturnURL struct {
	Url string `json:"url"`
}

func ServiceStart(port string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/helloworld", helloWorld)

	// POST /store {"url": "http://medium.com"}
	// {"url": "http://localhost:8080/id/abc123"}
	mux.HandleFunc("/store", store)

	// GET /id/abc123
	// Status 302 location: http://medium.com
	mux.HandleFunc("/id/", id)

	mux.HandleFunc("/", hello)

	http.ListenAndServe(port, mux)
}

func store(w http.ResponseWriter, r *http.Request) {

	// Guard clause
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var storeUrl StoreURL

	// Deserialize JSON from body
	if err := json.NewDecoder(r.Body).Decode(&storeUrl); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create unique ID
	returnUrl := ReturnURL{Url: "http://localhost:8080/id/abc123"}

	// Return JSON payload
	if err := json.NewEncoder(w).Encode(returnUrl); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func id(w http.ResponseWriter, r *http.Request) {
	// Guard clause
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Read abc123 from path
	path := r.URL.Path[4:]

	// Guard
	if path != "abc123" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Set status and location header
	w.Header().Set("Location", "http://medium.com")
	w.WriteHeader(http.StatusFound)
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
