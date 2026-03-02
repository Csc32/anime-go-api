package main

import (
	"anime-go-api/greetings"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Carlos", "Jose"}

	// Request greetings messages for the names

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	for _, message := range messages {
		fmt.Println(message)
	}
}
