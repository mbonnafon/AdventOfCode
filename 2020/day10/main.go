package main

import (
	"fmt"
	"sort"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

func main() {
	lines, _ := helpers.IntLines("./input.txt")
	lines = append(lines, 0)
	sort.Ints(lines)
	lines = append(lines, lines[len(lines)-1]+3)
	fmt.Println("Part 1. 1-jolt differences multiplied by the number of 3-jolt differences is :", pt1(lines))
	fmt.Printf("Part 2 dynamic programming. you can arrange the adapters %d distinct ways\n", pt2Dynamic(lines))
	fmt.Printf("Part 2 divide and conquer. you can arrange the adapters %d distinct ways\n", pt2DivideAndConquer(lines))
}

func pt1(lines []int) int {
	var oneJolt, threeJolt int
	for i := 0; i < len(lines)-1; i++ {
		if (lines[i+1] - lines[i]) == 1 {
			oneJolt++
		}
		if (lines[i+1] - lines[i]) == 3 {
			threeJolt++
		}
	}
	return oneJolt * threeJolt
}

// First way to solve the problem with dynamic programming
func pt2Dynamic(lines []int) int {
	m := make(map[int]int, len(lines))
	m[lines[len(lines)-1]] = 1
	for i := len(lines) - 2; i >= 0; i-- {
		m[lines[i]] = m[lines[i]+1] + m[lines[i]+2] + m[lines[i]+3]
	}
	return m[0]
}

// Second way to solve the problem with divide and conquer
// Counts the number of possibilities per sub-set.
// Each subset is separated by 3 jolt.
func pt2DivideAndConquer(lines []int) int {
	var pos int
	combinations := 1
	for i := 0; i < len(lines)-1; i++ {
		if lines[i+1]-lines[i] == 3 {
			combinations = combinations * countPossibilties(lines[pos:i+1], 0)
			pos = i + 1
			continue
		}
	}
	return combinations
}

func countPossibilties(jolts []int, curr int) int {
	var count int
	if curr == len(jolts)-1 {
		return 1
	}
	for i := 1; i <= 3; i++ {
		if curr < len(jolts)-i {
			c := jolts[curr+i] - jolts[curr]
			if c >= 1 && c <= 3 {
				count = count + countPossibilties(jolts, curr+i)
			}
		}
	}
	return count
}
