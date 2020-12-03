package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	input := "testInput"
	data, _ := ioutil.ReadFile(input)
	lines := strings.Split(string(data), "\n")
	got := Part1(lines)
	want := 2

	if got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := "testInput"
	data, _ := ioutil.ReadFile(input)
	lines := strings.Split(string(data), "\n")
	got := Part2(lines)
	want := 1

	if got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}
