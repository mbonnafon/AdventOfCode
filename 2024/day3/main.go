package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. :", pt1(lines))
	fmt.Println("Part 2. :", pt2(lines))
}

func pt1(lines []string) int {
	var re = regexp.MustCompile(`mul\((\d+,\d+)\)`)
	var result int
	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			r := strings.Split(match[1], ",")
			result += helpers.ToInt(r[0]) * helpers.ToInt(r[1])
		}
	}
	return result
}

func pt2(lines []string) int {
	var re = regexp.MustCompile(`mul\((\d+,\d+)\)|do(?:n't)?\(\)`)
	var result int
	calculationEnabled := true
	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			switch match[0] {
			case "do()":
				calculationEnabled = true
			case "don't()":
				calculationEnabled = false
			default:
				if calculationEnabled {
					r := strings.Split(match[1], ",")
					result += helpers.ToInt(r[0]) * helpers.ToInt(r[1])
				}
			}
		}
	}
	return result
}
