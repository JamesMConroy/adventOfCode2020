package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input")
	maxSeatID := 0
	seats := make([]int, 0)
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		seatID := SeatID(s)
		seats = append(seats, seatID)
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}
	fmt.Println("Highest Seat ID: ", maxSeatID)

	sort.Ints(seats)
	fmt.Println(FindSeat(seats))
}

func FindSeat(seats []int) int {
	mySeat := 0
	for i := 1; i < len(seats); i++ {
		if seats[i]-1 != seats[i-1] {
			mySeat = seats[i] - 1
		}
	}
	return mySeat
}

func SeatID(seat string) int {
	maxRow := 127
	minRow := 0
	Row := 63
	maxCol := 7
	minCol := 0
	Col := 4
	for i, s := range strings.Split(string(seat), "") {
		if i < 6 {
			if s == "B" {
				minRow = (minRow+maxRow)/2 + 1
			} else {
				maxRow = (minRow + maxRow) / 2
			}
		}
		if i == 6 {
			if s == "B" {
				Row = maxRow
			} else {
				Row = minRow
			}
		}
		if i == 7 || i == 8 {
			if s == "R" {
				minCol = (minCol+maxCol)/2 + 1
			} else {
				maxCol = (minCol + maxCol) / 2
			}
		}
		if i == 9 {
			if s == "R" {
				Col = maxCol
			} else {
				Col = minCol
			}
		}
	}
	seatID := (Row * 8) + Col
	return seatID
}
