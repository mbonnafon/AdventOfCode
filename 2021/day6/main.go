package main

import (
	"fmt"
	"strings"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	lanternfishs := helpers.ToIntSlice(strings.Split(lines[0], ","))
	fmt.Println("Part 1. :", pt1(80, lanternfishs))
	fmt.Println("Part 2. :", pt2(256, lanternfishs))
}

// This is the bruteforce solution
func pt1(days int, lanternfishs []int) int {
	for i := 0; i < days; i++ {
		var acc []int
		for j := range lanternfishs {
			if lanternfishs[j] == 0 {
				lanternfishs[j] = 6
				acc = append(acc, 8)
				continue
			}
			lanternfishs[j]--
		}
		lanternfishs = append(lanternfishs, acc...)
	}
	return len(lanternfishs)
}

// It solves the same problem but with a better complexity
func pt2(days int, lines []int) int {
	lanternfishs := make(map[int]int)
	for _, v := range lines {
		lanternfishs[v]++
	}

	for i := 0; i < days; i++ {
		childs := make(map[int]int)
		for timer, amount := range lanternfishs {
			if timer == 0 {
				childs[8] += amount
				childs[6] += amount
				continue
			}
			childs[timer-1] += amount
		}
		lanternfishs = childs
	}

	var score int
	for _, amount := range lanternfishs {
		score += amount
	}

	return score
}
