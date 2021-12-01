package main

import (
	"fmt"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

func pt1(lines []string) int {
	intLines := helpers.StringSliceToIntSlice(lines)
	var previous, counter int
	for k, current := range intLines {
		if k == 0 {
			previous = current
			continue
		}
		if current > previous {
			counter++
		}
		previous = current
	}
	return counter
}

func pt2(lines []string) int {
	intLines := helpers.StringSliceToIntSlice(lines)
	var previous, counter int
	for i := 2; i < len(intLines); i++ {
		current := intLines[i] + intLines[i-1] + intLines[i-2]
		if i == 2 {
			previous = intLines[i] + intLines[i-1] + intLines[i-2]
			continue
		}
		if current > previous {
			counter++
		}
		previous = current
	}
	return counter
}

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. :", pt1(lines))
	fmt.Println("Part 2. :", pt2(lines))
}
