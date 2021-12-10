package main

import (
	"fmt"
	"strings"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. :", pt1(lines))
	fmt.Println("Part 2. :", pt2(lines))
}

func pt1(commands []string) int {
	var horizontal, depth int
	for _, line := range commands {
		s := strings.Split(line, " ")
		direction := s[0]
		increment := helpers.ToInt(s[1])
		switch direction {
		case "forward":
			horizontal += increment
		case "down":
			depth += increment
		case "up":
			depth -= increment
		}
	}
	return horizontal * depth
}

func pt2(commands []string) int {
	var horizontal, depth, aim int
	for _, line := range commands {
		s := strings.Split(line, " ")
		direction := s[0]
		increment := helpers.ToInt(s[1])
		switch direction {
		case "down":
			aim += increment
		case "up":
			aim -= increment
		case "forward":
			horizontal += increment
			depth += aim * increment
		}
	}
	return horizontal * depth
}
