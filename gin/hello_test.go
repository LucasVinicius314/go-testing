package main

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "marcos"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name, 11)

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("marcos") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}
