package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	file := "testInput"
	data, _ := ioutil.ReadFile(file)
	lines := strings.Split(string(data), "\n")
	fmt.Println(Part1(lines))
	fmt.Println(Part2(lines))
}

func Part1(file []string) int {
	correct := 0
	for _, line := range file {
		line = strings.TrimSpace(line)
		var min, max int
		var char byte
		var passWord string
		fmt.Sscanf(line, "%v-%v %q: %v", &min, &max, &char, &passWord)

		numC := strings.Count(passWord, string(char))
		if numC >= min && numC <= max {
			correct++
		}
	}
	return correct
}

func Part2(file []string) int {
	return 0
}
