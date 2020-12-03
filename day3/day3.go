package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const tree = "#"

func main() {
	input, _ := ioutil.ReadFile("input")
	treeMap := make([]string, 0)
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		treeMap = append(treeMap, s)

	}
	fmt.Println("Number of Trees: ", Part1(treeMap, 1, 1))
	fmt.Println("Number of Trees: ", Part1(treeMap, 3, 1))
	fmt.Println("Number of Trees: ", Part1(treeMap, 5, 1))
	fmt.Println("Number of Trees: ", Part1(treeMap, 7, 1))
	fmt.Println("Number of Trees: ", Part1(treeMap, 1, 2))
}

func Part1(input []string, right int, down int) int {
	numTrees := 0
	width := len(input[1])
	height := len(input)
	x := 0
	for y := 0; y < height; y = y + down {
		if string(input[y][x]) == tree {
			numTrees++
		}
		//fmt.Println(x, y, string(input[y]), numTrees)
		x = x + right
		if x > width-1 {
			x = x - width
		}
	}
	return numTrees
}
