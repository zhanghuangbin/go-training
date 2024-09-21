package greetings

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name is empty")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	if names == nil {
		return nil, errors.New("names is empty")
	}
	result := make(map[string]string, len(names))
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("create %v error, %v", name, err))
		}
		result[name] = message
	}
	return result, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome",
		"Nice to see you, %v",
		"How are you, %v?",
	}

	return formats[rand.IntN(len(formats))]
}
