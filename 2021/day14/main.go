package main

import (
	"fmt"
	"strings"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	initial := lines[0]
	transform := make(map[string]string)
	for i := 2; i < len(lines); i++ {
		v := strings.Split(lines[i], " -> ")
		transform[v[0]] = v[1]
	}
	fmt.Println("Part 1. :", pt1(initial, transform))
	fmt.Println("Part 2. :", pt2(initial, transform))
}

func pt1(initial string, transform map[string]string) int {
	for i := 0; i < 10; i++ {
		initial = transformS(initial, transform)
	}
	return max(initial) - min(initial)
}

func pt2(initial string, transform map[string]string) int {
	occurences := make(map[string]int)
	for i := 0; i < len(initial)-1; i++ {
		v := string(initial[i]) + string(initial[i+1])
		occurences[v]++
	}
	for i := 0; i < 40; i++ {
		occurences = transformN(occurences, transform)
	}
	letters := make(map[string]int)
	for k, v := range occurences {
		letters[string(k[0])] += v
	}
	letters[string(initial[len(initial)-1])]++

	var max, min int
	for _, v := range letters {
		if v > max {
			max = v
		}
		if min == 0 {
			min = v
		}
		if v < min {
			min = v
		}
	}
	return max - min
}

func transformN(occurences map[string]int, transform map[string]string) map[string]int {
	new := make(map[string]int)
	for k, v := range occurences {
		v1 := string(k[0]) + transform[k]
		v2 := transform[k] + string(k[1])
		new[v1] += v
		new[v2] += v
	}
	return new
}

func transformS(initial string, transform map[string]string) string {
	var new string
	max := len(initial)
	for i := 0; i < max; i++ {
		if i == max-1 {
			new = new + string(initial[i])
			break
		}
		v := string(initial[i]) + string(initial[i+1])
		new = new + string(initial[i]) + transform[v]
	}
	return new
}

func max(s string) (max int) {
	occurences := make(map[string]int)
	for _, v := range s {
		occurences[string(v)]++
	}
	for _, v := range occurences {
		if v > max {
			max = v
		}
	}
	return
}

func min(s string) (min int) {
	occurences := make(map[string]int)
	for _, v := range s {
		occurences[string(v)]++
	}
	fmt.Println(occurences)
	for _, v := range occurences {
		if min == 0 {
			min = v
		}
		if v < min {
			min = v
		}
	}
	return
}
