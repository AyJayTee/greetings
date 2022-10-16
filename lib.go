package greetings

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
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
