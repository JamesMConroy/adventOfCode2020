package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const thresh = 2020

func main() {
	input := "input"
	fmt.Println("Part 1: ", Part1(input))
	fmt.Println("Part 2: ", Part2(input))
}

func Part1(input string) int {
	data, err := ioutil.ReadFile(input)
	if err != nil {
		fmt.Println("File reading error", err)
		return -1
	}
	lines := strings.Split(string(data), "\n")
	var nums = []int{}
	for _, i := range lines {
		j, _ := strconv.Atoi(i)
		nums = append(nums, j)
	}
	nums = nums[:len(nums)-1]
	for _, num := range nums {
		diff := thresh - num
		for _, num2 := range nums {
			if num2 <= diff {
				total := num * num2
				return total
			}
		}
	}
	return -1
}

func Part2(input string) int {
	data, err := ioutil.ReadFile(input)
	if err != nil {
		fmt.Println("File reading error", err)
		return -1
	}

	lines := strings.Split(string(data), "\n")
	var nums = []int{}

	for _, i := range lines {
		j, _ := strconv.Atoi(i)
		nums = append(nums, j)
	}
	nums = nums[:len(nums)-1]

	for _, num := range nums {
		if num < thresh {
			diff := thresh - num
			for _, num2 := range nums {
				if num2 < diff {
					diff2 := diff - num2
					for _, num3 := range nums {
						if num3 == diff2 {
							total := num * num2 * num3
							return total
						}
					}
				}
			}
		}
	}
	return -1
}
