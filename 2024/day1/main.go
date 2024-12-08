package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. :", pt1(lines))
	fmt.Println("Part 2. :", pt2(lines))
}

func parse(lines []string) ([]int, []int) {
	left := []int{}
	right := []int{}
	for _, line := range lines {
		s := strings.Split(line, "   ")
		left = append(left, helpers.ToInt(s[0]))
		right = append(right, helpers.ToInt(s[1]))
	}
	return left, right
}

func pt1(lines []string) int {
	left, right := parse(lines)
	sort.Ints(left)
	sort.Ints(right)

	var result int
	for i := 0; i < len(left); i++ {
		result += helpers.AbsInt(left[i] - right[i])
	}
	return result
}

func pt2(lines []string) int {
	left, right := parse(lines)

	numberOccurence := make(map[int]int)
	for _, numb := range right {
		numberOccurence[numb]++
	}

	var score int
	for _, numb := range left {
		score += numb * numberOccurence[numb]
	}
	return score
}
