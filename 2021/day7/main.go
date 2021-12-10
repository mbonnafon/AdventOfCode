package main

import (
	"fmt"
	"strings"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	crabsPos := helpers.ToIntSlice(strings.Split(lines[0], ","))
	fmt.Println("Part 1. :", pt1(crabsPos))
	fmt.Println("Part 2. :", pt2(crabsPos))
}

func pt1(crabsPos []int) (score int) {
	median := helpers.Median(crabsPos)
	for _, pos := range crabsPos {
		if pos == median {
			continue
		}
		if pos > median {
			score += (pos - median)
			continue
		}
		score += (median - pos)
	}
	return
}

func pt2(crabsPos []int) (score int) {
	mean := helpers.Mean(crabsPos)
	for _, pos := range crabsPos {
		if pos == mean {
			continue
		}
		if pos > mean {
			score += burnFuel(mean, pos)
			continue
		}
		score += burnFuel(pos, mean)
	}
	return
}

func burnFuel(start, end int) (score int) {
	acc := 1
	for i := start; i < end; i++ {
		score += acc
		acc++
	}
	return
}
