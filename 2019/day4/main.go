package main

import (
	"fmt"
	"strconv"
)

func validatePassword(i int) bool {
	var similar bool
	s := strconv.Itoa(i)
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			return false
		}
		if s[i] == s[i+1] {
			similar = true
		}
	}
	return similar
}

func pt1(lowerBoundary, upperBoundary int) int {
	var count int
	for lowerBoundary <= upperBoundary {
		if validatePassword(lowerBoundary) {
			count++
		}
		lowerBoundary++
	}
	return count
}

func validatePasswordPt2(i int) bool {
	var similar bool
	var count int
	s := strconv.Itoa(i)
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			return false
		}
		if s[i] == s[i+1] {
			if (i == len(s)-2) && similar && (count%2 == 1) {
				return false
			}
			similar = true
			count++
		} else {
			if similar && (count%2 == 0) {
				return false
			}
			similar = false
			count = 0
		}
	}
	return true
}

func pt2(lowerBoundary, upperBoundary int) int {
	var count int
	for lowerBoundary <= upperBoundary {
		if validatePasswordPt2(lowerBoundary) {
			count++
		}
		lowerBoundary++
	}
	return count
}

func main() {
	fmt.Printf("Part 1. %d passwords within the range input meet criterias\n", pt1(356261, 846303))
	fmt.Printf("Part 2. %d passwords within the range input meet criterias\n", pt2(356261, 846303))
}
