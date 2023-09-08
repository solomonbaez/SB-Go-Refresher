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

// messages type -> map with string keys, string values
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		var message string
		var err error

		message, err = Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v welcome!",
		"Ooo, wasn't expecting you, %v...",
		"Greetings, %v!",
	}

	return formats[rand.Intn(len(formats))]
}
