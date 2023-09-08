package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	var message string
	var err error

	message, err = greetings.Hello("Goobab")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
