package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Rogery"
	want := regexp.MustCompile(`\b` + name + `\b`)
	greeting, err := Hello(name)
	if !want.MatchString(greeting) || err != nil {
		t.Fatalf(`Hello(\"Gladys\") = %q, %v, want match for %#q, nil`, greeting, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	name := ""
	greeting, err := Hello(name)
	if greeting != "" || err == nil {
		t.Fatalf("Exepect to throw an error but got (%v) and return and empty message", err)
	}
}
