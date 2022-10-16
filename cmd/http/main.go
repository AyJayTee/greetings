package main

import (
	"flag"

	"github.com/AyJayTee/greetings/internal/server"
)

func main() {

	port := flag.String("p", ":8080", "port")
	flag.Parse()

	server.ServiceStart(*port)
}
