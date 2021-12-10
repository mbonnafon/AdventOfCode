package main

import (
	"fmt"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. :", pt1(helpers.ToIntSlice(lines)))
	fmt.Println("Part 2. :", pt2(helpers.ToIntSlice(lines)))
}

func pt1(lines []int) (counter int) {
	for i := 1; i < len(lines); i++ {
		if lines[i] > lines[i-1] {
			counter++
		}
	}
	return
}

func pt2(lines []int) (counter int) {
	previous := lines[0] + lines[1] + lines[2]
	for i := 2; i < len(lines); i++ {
		current := lines[i] + lines[i-1] + lines[i-2]
		if current > previous {
			counter++
		}
		previous = current
	}
	return
}
