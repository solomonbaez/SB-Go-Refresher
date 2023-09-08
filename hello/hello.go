package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	var message string
	message = greetings.Hello("Goobab")

	fmt.Println(message)
}
