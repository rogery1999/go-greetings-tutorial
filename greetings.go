package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("Empty name")
	}
	return fmt.Sprintf(getRandomFormat(), name), nil
}

func HelloNames(names ...string) (map[string]string, error) {
	if names == nil || len(names) == 0 {
		return nil, errors.New("there are no valid information")
	}

	greetings := make(map[string]string)
	for _, name := range names {
		_, found := greetings[name]
		switch {
		case found:
			return nil, errors.New(fmt.Sprintf("A duplicated name was found: %v", name))
		case name == "":
			return nil, errors.New("empty name")
		}
		greetings[name] = fmt.Sprintf(getRandomFormat(), name)
	}
	return greetings, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getRandomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
