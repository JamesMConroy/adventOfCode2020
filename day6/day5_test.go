package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	input, _ := ioutil.ReadFile("testInput")
	want := 11
	got := 0
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n\n") {
		got = got + NumYes(string(s))
	}

	if got != want {
		t.Errorf("got: %d want: %d", got, want)
	}
}
func TestPart2(t *testing.T) {
	input, _ := ioutil.ReadFile("testInput")
	want := 6
	got := 0
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n\n") {
		got = got + AllYes(string(s))
	}

	if got != want {
		t.Errorf("got: %d want: %d", got, want)
	}
}
