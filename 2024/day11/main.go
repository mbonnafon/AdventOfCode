package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

type Cache map[CacheKey]*int

type CacheKey struct {
	value int
	epoch int
}

func (c Cache) countBlink(currentEpoch, targetEpoch int, grave int) int {
	if currentEpoch == targetEpoch {
		return 1
	}

	cacheKey := CacheKey{grave, currentEpoch}
	gravesCount := c[cacheKey]
	if gravesCount != nil {
		return *gravesCount
	}

	graveAfterBlink := replace(strconv.Itoa(grave))

	if graveAfterBlink.right == nil {
		result := c.countBlink(currentEpoch+1, targetEpoch, graveAfterBlink.left)
		c[CacheKey{graveAfterBlink.left, currentEpoch + 1}] = &result
		return result
	}

	leftResult := c.countBlink(currentEpoch+1, targetEpoch, graveAfterBlink.left)
	c[CacheKey{graveAfterBlink.left, currentEpoch + 1}] = &leftResult
	rightResult := c.countBlink(currentEpoch+1, targetEpoch, *graveAfterBlink.right)
	c[CacheKey{*graveAfterBlink.right, currentEpoch + 1}] = &rightResult
	return leftResult + rightResult
}

type Grave struct {
	left  int
	right *int
}

func replace(grave string) *Grave {
	if grave == "0" {
		return &Grave{1, nil}
	}

	if len(grave)%2 == 0 { //even
		middle := len(grave) / 2
		left, right := (helpers.ToInt(grave[:middle])), (helpers.ToInt(grave[middle:]))
		return &Grave{left, &right}
	}
	return &Grave{(helpers.ToInt(grave) * 2024), nil}
}

func blink(graves []string) []string {
	var newGraveArrangement []string
	for _, grave := range graves {
		grave := replace(grave)
		newGraveArrangement = append(newGraveArrangement, strconv.Itoa(grave.left))
		if grave.right != nil {
			newGraveArrangement = append(newGraveArrangement, strconv.Itoa(*grave.right))
		}
	}
	return newGraveArrangement
}

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. :", pt1(lines))
	fmt.Println("Part 2. :", pt2(lines))
}

func pt1(lines []string) int {
	graves := strings.Split(lines[0], " ")
	for i := 0; i < 25; i++ {
		graves = blink(graves)
	}
	return len(graves)
}

func pt2(lines []string) int {
	graves := strings.Split(lines[0], " ")
	cache := make(Cache)
	blinkCount := 75

	var count int
	for _, grave := range graves {
		count += cache.countBlink(0, blinkCount, helpers.ToInt(grave))
	}
	return count
}
