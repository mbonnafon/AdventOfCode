package main

import (
	"fmt"
	"slices"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. :", pt1(lines))
	fmt.Println("Part 2. :", pt2(lines))
}

type elf int
type calories int

func parseFoodForElves(lines []string) map[elf]calories {
	foodCarrier := make(map[elf]calories)
	elfNumber := elf(1)
	for _, l := range lines {
		if l == "" {
			elfNumber++
			continue
		}
		foodCarrier[elfNumber] += calories(helpers.ToInt(l))
	}
	return foodCarrier
}

func pt1(lines []string) int {
	foodCarrier := parseFoodForElves(lines)
	var max calories
	for _, calories := range foodCarrier {
		if calories > max {
			max = calories
		}
	}
	return int(max)
}

func pt2(lines []string) int {
	foodCarrier := parseFoodForElves(lines)
	food := make([]int, 0)
	for _, calorie := range foodCarrier {
		food = append(food, int(calorie))
	}

	slices.Sort(food)
	slices.Reverse(food)

	return food[0] + food[1] + food[2]
}
