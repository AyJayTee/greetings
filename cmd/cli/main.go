package main

import (
	"fmt"
	"os"

	"github.com/AyJayTee/greetings"
)

func main() {
	name := "World"

	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	fmt.Println(greetings.Hello(name))
}
