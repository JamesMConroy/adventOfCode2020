package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

const file = "testIn"

func TestPart1(t *testing.T) {
	fileb, _ := ioutil.ReadFile(file)
	input := make([]string, 0)
	for _, s := range strings.Split(strings.TrimSpace(string(fileb)), "\n") {
		input = append(input, s)

	}
	t.Run("Right 1, Down 1", func(t *testing.T) {
		got := Part1(input, 1, 1)
		want := 2

		if got != want {
			t.Errorf("Got %d, want: %d)", got, want)
		}
	})
	t.Run("Right 3, Down 1", func(t *testing.T) {
		got := Part1(input, 3, 1)
		want := 7

		if got != want {
			t.Errorf("Got %d, want: %d)", got, want)
		}
	})
	t.Run("Right 5, Down 1", func(t *testing.T) {
		got := Part1(input, 5, 1)
		want := 3

		if got != want {
			t.Errorf("Got %d, want: %d)", got, want)
		}
	})
	t.Run("Right 7, Down 1", func(t *testing.T) {
		got := Part1(input, 7, 1)
		want := 4

		if got != want {
			t.Errorf("Got %d, want: %d)", got, want)
		}
	})
	t.Run("Right 1, Down 2", func(t *testing.T) {
		got := Part1(input, 1, 2)
		want := 2

		if got != want {
			t.Errorf("Got %d, want: %d)", got, want)
		}
	})
}
