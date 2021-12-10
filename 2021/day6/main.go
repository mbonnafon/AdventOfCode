package main

import (
	"fmt"
	"strings"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. :", pt1(80, strings.Split(lines[0], ",")))
	fmt.Println("Part 2. :", pt2(256, strings.Split(lines[0], ",")))
}

// This is the bruteforce solution
func pt1(days int, lines []string) int {
	lanternfishs := helpers.ToIntSlice(lines)
	for i := 0; i < days; i++ {
		var acc []int
		for i := 0; i < len(lanternfishs); i++ {
			if lanternfishs[i] == 0 {
				lanternfishs[i] = 6
				acc = append(acc, 8)
				continue
			}
			lanternfishs[i]--
		}
		lanternfishs = append(lanternfishs, acc...)
	}
	return len(lanternfishs)
}

// It solves the same problem but with a better complexity
func pt2(days int, lines []string) int {
	lanternfishs := make(map[int]int)
	for _, v := range helpers.ToIntSlice(lines) {
		lanternfishs[v]++
	}

	for i := 0; i < days; i++ {
		m := make(map[int]int)
		for k, v := range lanternfishs {
			if k == 0 {
				m[8] += v
				m[6] += v
				continue
			}
			m[k-1] += v
		}
		lanternfishs = m
	}

	var score int
	for _, v := range lanternfishs {
		score += v
	}

	return score
}
