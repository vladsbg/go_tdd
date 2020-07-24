package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Olá, Chris"

	if got != want {
		t.Errorf(" obteve '%s', queria '%s'", got, want)
	}
}
