package main

import (
	"fmt"
	"strconv"
)

func part1(lowerBoundary, upperBoundary int) int {
	var counter int
	for i := lowerBoundary; i <= upperBoundary; i++ {
		istring := strconv.Itoa(i)
		similiar := false
		for j := 0; j < 5; j++ {
			if istring[j:j+1] == istring[j+1:j+2] {
				similiar = true
			}
			if istring[j:j+1] > istring[j+1:j+2] {
				j = 5
			} else if (j == 4) && (similiar) {
				counter++
			}
		}
	}
	return counter
}

// Part2 z
func Part2(lowerBoundary, upperBoundary int) int {
	var counter int
	for i := lowerBoundary; i <= upperBoundary; i++ {
		istring := strconv.Itoa(i)
		mapOfNumbers := make(map[string]int)
		similiar := false
		for j := 0; j < 5; j++ {
			val1 := istring[j : j+1]
			val2 := istring[j+1 : j+2]

			if val1 > val2 {
				j = 5
			}

			if val1 == val2 {
				mapOfNumbers[val1]++
			} else if mapOfNumbers[val1]%2 != 0 {
				delete(mapOfNumbers, val1)
				similiar = true
			}
			if j == 4 {
				if mapOfNumbers[val1]%2 != 0 {
					delete(mapOfNumbers, val1)
					similiar = true
				}
				if len(mapOfNumbers) == 0 && similiar {
					counter++
				}
			}
			//fmt.Println(j, "-", val1, val2, mapOfNumbers)
		}
		//fmt.Println("_____________")
	}
	return counter
}

func main() {
	fmt.Println(part1(356261, 846303))
	fmt.Println(Part2(356261, 846303))
}
