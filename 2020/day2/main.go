package main

import (
	"fmt"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

var (
	min, max int
	letter   rune
	password string
)

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Printf("Part1. %d passwords are valid\n", pt1(lines))
	fmt.Printf("Part2. %d passwords are valid\n", pt2(lines))
}

func pt1(lines []string) int {
	var count int
	for _, l := range lines {
		fmt.Sscanf(l, "%d-%d %c: %s", &min, &max, &letter, &password)
		letterOccurence := make(map[string]int)
		for _, c := range password {
			letterOccurence[string(c)] = letterOccurence[string(c)] + 1
		}
		if (letterOccurence[string(letter)] >= min) && (letterOccurence[string(letter)] <= max) {
			count++
		}
	}
	return count
}

func pt2(lines []string) int {
	var count int
	for _, l := range lines {
		fmt.Sscanf(l, "%d-%d %c: %s", &min, &max, &letter, &password)
		l1, l2, validLetter := password[min-1], password[max-1], byte(letter)
		if l1 != l2 && (l1 == validLetter || l2 == validLetter) {
			count++
		}
	}
	return count
}
