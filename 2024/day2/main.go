package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. :", pt1(lines))
	fmt.Println("Part 2. :", pt2(lines))
}

func isSafe(levels []int) bool {
	ascending := levels[0] < levels[1]

	// start at 1 as we have the state of the 0 index in previousLevel
	for i := 1; i < len(levels); i++ {
		previousLevel, currentLevel := levels[i-1], levels[i]

		if ascending && previousLevel > currentLevel {
			return false
		}
		if !ascending && previousLevel < currentLevel {
			return false
		}

		absoluteDiff := helpers.AbsInt(previousLevel - currentLevel)
		if absoluteDiff > 3 || absoluteDiff < 1 {
			return false
		}
	}
	return true
}

func remove(s []string, idx int) []string {
	return append(s[:idx], s[idx+1:]...)
}

func pt1(lines []string) int {
	var safeCount int
	for _, report := range lines {
		levels := strings.Split(report, " ")
		if isSafe(helpers.ToIntSlice(levels)) {
			safeCount++
		}
	}
	return safeCount
}

func pt2(lines []string) int {
	var safeCount int
	for _, report := range lines {
		levels := strings.Split(report, " ")
		for i := range levels {
			levelWithoutOneError := remove(slices.Clone(levels), i)
			if isSafe(helpers.ToIntSlice(levelWithoutOneError)) {
				safeCount++
				break
			}
		}
	}
	return safeCount
}
