package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/AyJayTee/greetings"
)

type StoreURL struct {
	Url string `json:"url"`
}

type ReturnURL struct {
	Url string `json:"url"`
}

type Database struct {
	data []string
}

func ServiceStart(port string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/helloworld", helloWorld)

	s := &Database{
		data: make([]string, 0),
	}

	// POST /store {"url": "http://medium.com"}
	// {"url": "http://localhost:8080/id/abc123"}
	mux.Handle("/store", storeHandler(s))

	// GET /id/abc123
	// Status 302 location: http://medium.com
	mux.Handle("/id/", idHandler(s))

	mux.HandleFunc("/", hello)

	http.ListenAndServe(port, auth("user", "pass", logging(mux)))
}

func storeHandler(store *Database) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

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
		// Validation
		if storeUrl.Url == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !strings.HasPrefix(storeUrl.Url, "https://") {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		store.data = append(store.data, storeUrl.Url)

		// Create unique ID
		returnUrl := ReturnURL{Url: fmt.Sprintf("http://localhost:8080/id/%d", len(store.data))}

		// Return JSON payload
		if err := json.NewEncoder(w).Encode(returnUrl); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func idHandler(store *Database) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// Guard clause
			if r.Method != http.MethodGet {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}

			// Read abc123 from path
			path := r.URL.Path[4:]
			// Guard
			i, err := strconv.Atoi(path)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if i > len(store.data) {
				log.Printf("OUt of bounds, %d, %d", i, len(store.data))
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			// Guard
			url := store.data[i-1]

			// Set status and location header
			w.Header().Set("Location", url)
			w.WriteHeader(http.StatusFound)
		})
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

func auth(user, password string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		usr, pwd, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if usr != user {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		if pwd != password {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s", r.URL)
		next.ServeHTTP(w, r)
	})
}
