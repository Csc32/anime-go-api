package main

import (
	"anime-go-api/greetings"
	"fmt"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
