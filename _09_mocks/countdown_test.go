package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}
	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf(" obteve '%s', queria '%s'", got, want)
	}

	want2 := 3
	if spySleeper.Calls != want2 {
		t.Errorf("não houve chamadas suficientes do sleeper, obteve %d esperado %d", spySleeper.Calls, want2)
	}
}
