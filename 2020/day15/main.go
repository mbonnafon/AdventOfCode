package main

import (
	"fmt"
)

func main() {
	input := []int{16, 1, 0, 18, 12, 14, 19}
	fmt.Println("Part 1. :", game(input, 2020))
	fmt.Println("Part 2. :", game(input, 30000000))
}

func game(startingNumbers []int, iterations int) int {
	var previous int
	lastSpoken := make(map[int]int)
	spokenBefore := make(map[int]int)

	for turn, number := range startingNumbers {
		spokenBefore[number] = turn + 1
		previous = number
	}
	for i := len(startingNumbers) + 1; i <= iterations; i++ {
		if lastSpoken[previous] != 0 && spokenBefore[previous] != 0 {
			previous = lastSpoken[previous] - spokenBefore[previous]
			if lastSpoken[previous] != 0 {
				spokenBefore[previous] = lastSpoken[previous]
			}
			lastSpoken[previous] = i
		} else if lastSpoken[previous] != 0 {
			previous = 0
			if lastSpoken[previous] != 0 {
				spokenBefore[previous] = lastSpoken[previous]
			}
			lastSpoken[previous] = i
		} else {
			previous = 0
			lastSpoken[previous] = i
		}
	}
	return previous
}
