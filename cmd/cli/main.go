package main

import (
	"fmt"
	"os"

	"github.com/AyJayTee/greetings/internal/memory"
)

func main() {

	db := memory.NewDatabase([]string{"https://google.com", "https://medium.com"})

	name := "World"

	if len(os.Args) > 2 {
		name = os.Args[1]
	}

	switch name {
	case "get":
		url, err := db.FetchUrl(os.Args[2])
		if err != nil {
			panic(err)
		}
		fmt.Println(url)
	case "set":
		id, err := db.AddUrl(os.Args[2])
		if err != nil {
			panic(err)
		}
		fmt.Println(id)
	}
}
