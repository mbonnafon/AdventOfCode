package main

import (
	"fmt"
	"strings"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	crabsPos := helpers.StringSliceToIntSlice(strings.Split(lines[0], ","))
	fmt.Println("Part 1. :", pt1(crabsPos))
	fmt.Println("Part 2. :", pt2(crabsPos))
}

func pt1(crabsPos []int) int {
	var score int
	median := helpers.Median(crabsPos)
	for _, v := range crabsPos {
		if v == median {
			continue
		}
		if v > median {
			score += (v - median)
			continue
		}
		score += (median - v)
	}
	return score
}

func pt2(crabsPos []int) int {
	var score int
	mean := helpers.Mean(crabsPos)
	for _, v := range crabsPos {
		if v == mean {
			continue
		}
		if v > mean {
			score += rec(mean, v)
			continue
		}
		score += rec(v, mean)
	}
	return score
}

func rec(start, end int) int {
	var score int
	acc := 1
	for i := start; i < end; i++ {
		score += acc
		acc++
	}
	return score
}
