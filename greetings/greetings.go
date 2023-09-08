package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	var message string
	message = fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v welcome!",
		"Ooo, wasn't expecting you, %v...",
		"Greetings, %v!",
	}

	return formats[rand.Intn(len(formats))]
}
