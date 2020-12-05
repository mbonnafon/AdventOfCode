package main

import (
	"fmt"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

func fuelCalcul(mass int) int {
	neededFuel := (mass/3 - 2)
	if neededFuel > 0 {
		return neededFuel + fuelCalcul(neededFuel)
	}
	return 0
}

func pt1(lines []int) int {
	var fuel int
	for _, mass := range lines {
		fuel = fuel + (mass/3 - 2)
	}
	return fuel
}

func pt2(lines []int) int {
	var fuel int
	for _, mass := range lines {
		fuel = fuel + fuelCalcul(mass)
	}
	return fuel
}

func main() {
	lines, _ := helpers.IntLines("./input.txt")
	fmt.Println("Part 1. sum of the fuel requirements is:", pt1(lines))
	fmt.Println("Part 2. sum of the fuel requirements is:", pt2(lines))
}
