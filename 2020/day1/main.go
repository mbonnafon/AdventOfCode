package main

import (
	"fmt"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

func pt1(target int, lines []int) int {
	var existingValues = make(map[int]bool)
	for _, v := range lines {
		if existingValues[target-v] {
			return v * (target - v)
		}
		existingValues[v] = true
	}
	return 0
}

func pt2(target int, lines []int) int {
	var existingValues = make(map[int]bool)
	for _, v := range lines {
		existingValues[v] = true
	}
	for _, v1 := range lines {
		for _, v2 := range lines {
			v3 := target - v1 - v2
			if existingValues[v3] {
				return v1 * v2 * v3
			}
		}
	}
	return 0
}

func main() {
	const target = 2020
	lines, _ := helpers.IntLines("./input.txt")
	fmt.Printf("Part 1. the product of the two entries that sum to %d is: %d\n", target, pt1(target, lines))
	fmt.Printf("Part 2. the product of the three entries that sum %d is: %d\n", target, pt2(target, lines))
}
