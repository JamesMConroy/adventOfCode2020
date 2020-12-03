package main

import "testing"

func TestPart1(t *testing.T) {
	file := "testInput"
	got := Part1(file)
	want := 514579

	if got != want {
		t.Errorf("got: %d, wanted: %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	file := "testInput"
	got := Part2(file)
	want := 241861950

	if got != want {
		t.Errorf("got: %d, wanted: %d", got, want)
	}
}
