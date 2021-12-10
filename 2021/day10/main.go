package main

import (
	"fmt"
	"sort"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. :", pt1(lines))
	fmt.Println("Part 2. :", pt2(lines))
}

func pt1(lines []string) (score int) {
	for _, l := range lines {
		switch processCorruptedLine(l) {
		case ")":
			score += 3
		case "]":
			score += 57
		case "}":
			score += 1197
		case ">":
			score += 25137
		}
	}
	return
}

func pt2(lines []string) int {
	var incompleteLines []string
	for _, l := range lines {
		if processCorruptedLine(l) == "" {
			incompleteLines = append(incompleteLines, l)
		}
	}

	var scores []int
	for _, l := range incompleteLines {
		scores = append(scores, scoreIncompleteLine(l))
	}
	sort.Ints(scores)

	return helpers.Median(scores)
}

func processCorruptedLine(l string) string {
	var stack helpers.Stack
	for _, c := range l {
		symbol := string(c)
		if symbol == "(" || symbol == "{" || symbol == "[" || symbol == "<" {
			stack.Push(symbol)
			continue
		}
		switch symbol {
		case ")":
			if stack.Pop() != "(" {
				return symbol
			}
		case "}":
			if stack.Pop() != "{" {
				return symbol
			}
		case "]":
			if stack.Pop() != "[" {
				return symbol
			}
		case ">":
			if stack.Pop() != "<" {
				return symbol
			}
		}
	}
	return ""
}

func scoreIncompleteLine(line string) int {
	var stack helpers.Stack
	for _, c := range line {
		symbol := string(c)
		if symbol == "(" || symbol == "{" || symbol == "[" || symbol == "<" {
			stack.Push(symbol)
			continue
		}
		if symbol == stack.Pop() {
			continue
		}
	}

	var score int
	helpers.Reverse(stack)
	for _, c := range stack {
		score *= 5
		switch string(c) {
		case "(":
			score += 1
		case "{":
			score += 3
		case "[":
			score += 2
		case "<":
			score += 4
		}
	}
	return score
}
