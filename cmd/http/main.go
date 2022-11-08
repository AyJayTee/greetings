package main

import (
	"flag"
	"log"

	"github.com/AyJayTee/greetings/internal/server"
)

func main() {
	log.Println("Running...")
	port := flag.String("p", ":8080", "port")
	flag.Parse()

	server.ServiceStart(*port)
}
