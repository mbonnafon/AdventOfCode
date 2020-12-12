package main

import (
	"fmt"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

func main() {
	lines, _ := helpers.IntLines("./input.txt")
	preamble := 5
	fmt.Printf("Part 1. the first number which is not the sum of two of the %d numbers before it is %d\n", preamble, pt1(lines, preamble))
	fmt.Println("Part 2. the encryption weakness in your XMAS-encrypted list of numbers is:", pt2(lines, pt1(lines, preamble)))
}

func pt1(lines []int, preamble int) int {
	for k, v := range lines {
		minterval := IntIntervalToMap(lines[k : k+preamble])
		x1, x2 := TwoSum(minterval, lines[k+preamble])
		if x1 == 0 {
			return x2
		}
		if minterval[x1] > 1 {
			minterval[x1]--
			continue
		}
		delete(minterval, x1)
		minterval[v]++
	}
	return -1
}

func pt2(lines []int, target int) int {
	for index, v := range lines {
		var sum, highest int
		smallest := v
		for _, currentVal := range lines[index:] {
			sum += currentVal
			if currentVal > highest {
				highest = currentVal
			}
			if currentVal < smallest {
				smallest = currentVal
			}
			if sum == target {
				return highest + smallest
			}
			if sum > target {
				break
			}
		}
	}
	return 0
}

func IntIntervalToMap(lines []int) map[int]int {
	m := make(map[int]int)
	for _, v := range lines {
		m[v]++
	}
	return m
}

// TwoSum looks for two existing numbers that can be summed up to target
func TwoSum(numbers map[int]int, target int) (int, int) {
	var x2 int
	for x1 := range numbers {
		x2 = target - x1
		if numbers[x2] != 0 {
			if numbers[x2] == x2 && numbers[x2] > 1 {
				return x2, numbers[x2]
			}
			return x2, numbers[x2]
		}
	}
	return 0, target
}
