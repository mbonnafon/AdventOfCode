package main

import (
	"fmt"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

type Slope struct {
	x int
	y int
}

func (s Slope) countTrees(lines []string) int {
	var count, x, y int
	xlen := len(lines[0])
	for y < (len(lines) - 1) {
		if (x + s.x) >= xlen {
			x = x - xlen
		}
		x, y = x+s.x, y+s.y
		if string(lines[y][x]) == "#" {
			count++
		}
	}
	return count
}

func pt1(lines []string) int {
	slope := Slope{
		x: 3,
		y: 1,
	}
	return slope.countTrees(lines)
}

func pt2(lines []string) int {
	slopes := []Slope{
		{
			x: 1,
			y: 1,
		},
		{
			x: 3,
			y: 1,
		},
		{
			x: 5,
			y: 1,
		},
		{
			x: 7,
			y: 1,
		},
		{
			x: 1,
			y: 2,
		},
	}
	count := 1
	for _, s := range slopes {
		count = count * s.countTrees(lines)
	}
	return count
}

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Printf("Part1. I encountered %d trees\n", pt1(lines))
	fmt.Printf("Part2. I encountered %d trees\n", pt2(lines))
}
