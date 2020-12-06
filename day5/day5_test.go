package main

import "testing"

func TestSeatID(t *testing.T) {

	t.Run("BFFFBBFRRR", func(t *testing.T) {
		want := 567
		input := "BFFFBBFRRR"
		got := SeatID(input)

		if want != got {
			t.Errorf("Seat: %s got: %d want: %d", input, got, want)
		}
	})
	t.Run("FFFBBBFRRR", func(t *testing.T) {
		want := 119
		input := "FFFBBBFRRR"
		got := SeatID(input)

		if want != got {
			t.Errorf("Seat: %s got: %d want: %d", input, got, want)
		}
	})
	t.Run("BBFFBBFRLL", func(t *testing.T) {
		want := 820
		input := "BBFFBBFRLL"
		got := SeatID(input)

		if want != got {
			t.Errorf("Seat: %s got: %d want: %d", input, got, want)
		}
	})
}
func TestFindSeat(t *testing.T) {
	seats := []int{101, 102, 104, 105, 106}
	want := 103
	got := FindSeat(seats)

	if got != want {
		t.Errorf("got: %d want: %d", got, want)
	}
}
