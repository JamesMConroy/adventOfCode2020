package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input")
	part1, part2 := 0, 0
	for _, s := range strings.Split(string(input), "\n\n") {
		part1 = part1 + NumYes(s)
		part2 = part2 + AllYes(s)
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func AllYes(input string) int {
	encountered := map[string]int{}
	numInGroup := strings.Count(input, "\n") + 1
	for _, v := range input {
		for _, s := range strings.Split(strings.TrimSpace(string(v)), "\n") {
			if string(s) != "" {
				if _, ok := encountered[string(s)]; ok == false {
					encountered[string(s)] = 0
				}
				encountered[string(s)] = encountered[string(s)] + 1
			}
		}
	}
	numAllYes := 0
	for _, number := range encountered {
		if number == numInGroup {
			numAllYes++
		}
	}
	fmt.Println(encountered)
	fmt.Println("Num in group", numInGroup)
	fmt.Println("Num All Yes", numAllYes)
	return numAllYes
}

func NumYes(input string) int {
	encountered := map[string]bool{}
	for _, v := range input {
		for _, s := range strings.Split(strings.TrimSpace(string(v)), "\n") {
			if encountered[string(s)] != true && string(s) != "\n" && string(s) != "" {
				encountered[string(s)] = true
			}
		}
	}
	return len(encountered)
}
