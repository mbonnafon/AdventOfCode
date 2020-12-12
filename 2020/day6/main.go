package main

import (
	"fmt"
	"strings"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

func main() {
	lines, _ := helpers.StringGroupOfLines("./input.txt")
	fmt.Println("Part 1. :", pt1(lines))
	fmt.Println("Part 2. :", pt2(lines))
}

func pt1(lines []string) int {
	var count int
	for _, l := range lines {
		count = count + countGroupAnswer(l)
	}
	return count
}

func pt2(lines []string) int {
	var count int
	for _, l := range lines {
		count = count + countEveryoneAnswer(l)
	}
	return count
}

func countGroupAnswer(s string) int {
	m := make(map[rune]int)
	for _, p := range strings.Split(s, " ") {
		for _, c := range p {
			m[c]++
		}
	}
	return len(m)
}

func countEveryoneAnswer(s string) int {
	var count int
	personAnswer := strings.Split(s, " ")
	m := make(map[rune]int)
	for _, p := range personAnswer {
		for _, c := range p {
			m[c]++
		}
	}
	for _, v := range m {
		if v == len(personAnswer) {
			count++
		}
	}
	return count
}
