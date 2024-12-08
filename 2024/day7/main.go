package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. :", pt1(lines))
	fmt.Println("Part 2. :", pt2(lines))
}

func generateAllOperatorCombinations(operators []string, numbers int) [][]string {
	var combinations [][]string

	var generate func(current []string)
	generate = func(current []string) {
		if len(current) == numbers {
			combination := make([]string, numbers)
			copy(combination, current)
			combinations = append(combinations, combination)
			return
		}
		for _, operator := range operators {
			current = append(current, operator)
			generate(current)
			current = current[:len(current)-1]
		}
	}

	generate([]string{})

	return combinations
}

func computeIsValid(target int, computations [][]string, numbers []int) bool {
	for _, computation := range computations {
		result := numbers[0]
		for i, operator := range computation {
			switch operator {
			case "*":
				result *= numbers[i+1]
			case "+":
				result += numbers[i+1]
			case "||":
				result = helpers.ToInt(strconv.Itoa(result) + strconv.Itoa(numbers[i+1]))
			}
		}
		if result == target {
			return true
		}
	}
	return false
}

func findCalibration(lines []string, operators []string) int {
	var result int
	for _, line := range lines {
		s := strings.Split(line, ": ")
		expectedResult := helpers.ToInt(s[0])
		numbers := helpers.ToIntSlice(strings.Split(s[1], " "))

		combinations := generateAllOperatorCombinations(operators, len(numbers)-1)
		if computeIsValid(expectedResult, combinations, numbers) {
			result += expectedResult
		}
	}
	return result
}

func pt1(lines []string) int {
	var operators = []string{"+", "*"}
	return findCalibration(lines, operators)
}

func pt2(lines []string) int {
	var operators = []string{"+", "*", "||"}
	return findCalibration(lines, operators)
}
