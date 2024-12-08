package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. :", pt1(lines))
	fmt.Println("Part 2. :", pt2(lines))
}

type PageOrdering struct {
	before int
	after  int
}

func parse(lines []string) ([]PageOrdering, []string) {
	var lineBreakFound bool
	pageOrdering := []PageOrdering{}
	updates := []string{}
	for _, line := range lines {
		if line == "" {
			lineBreakFound = true
			continue
		}
		if !lineBreakFound {
			s := strings.Split(line, "|")
			pageOrdering = append(pageOrdering, PageOrdering{helpers.ToInt(s[0]), helpers.ToInt(s[1])})
		} else {
			updates = append(updates, line)
		}
	}
	return pageOrdering, updates
}

func findUpdatesFor(expected string, pageOrderings []PageOrdering, updates []string) [][]int {
	var target [][]int

	for _, line := range updates {
		numToIndex := make(map[int]*int)
		updateSequence := helpers.ToIntSlice(strings.Split(line, ","))
		for pos, num := range updateSequence {
			numToIndex[num] = &pos
		}

		valid := true
		for _, pageOrdering := range pageOrderings {
			if numToIndex[pageOrdering.before] == nil || numToIndex[pageOrdering.after] == nil {
				continue
			}
			if *numToIndex[pageOrdering.before] > *numToIndex[pageOrdering.after] {
				valid = false
				break
			}
		}
		if valid {
			if expected == "valid" {
				target = append(target, updateSequence)
			}
		} else {
			if expected == "invalid" {
				target = append(target, updateSequence)
			}
		}
	}

	return target
}

func pt1(lines []string) int {
	pageOrderings, updates := parse(lines)
	validLines := findUpdatesFor("valid", pageOrderings, updates)

	var count int
	for _, line := range validLines {
		count += line[len(line)/2]
	}

	return count
}

func pt2(lines []string) int {
	pageOrderings, updates := parse(lines)
	invalidUpdates := findUpdatesFor("invalid", pageOrderings, updates)

	shouldBeAfter := make(map[int][]int)
	for _, pageOrder := range pageOrderings {
		shouldBeAfter[pageOrder.after] = append(shouldBeAfter[pageOrder.after], pageOrder.before)
	}

	var count int
	for _, update := range invalidUpdates {
		ordered := order(0, update, shouldBeAfter)
		count += ordered[len(ordered)/2]
	}
	return count
}

func order(initial int, updateSequence []int, shouldBeAfter map[int][]int) []int {
	for i := initial; i < len(updateSequence); i++ {
		pos := updateSequence[i]
		if contains(updateSequence[i+1:], shouldBeAfter[pos]) {
			s := slices.Clone(updateSequence)
			return order(i, append(append(s[:i], s[i+1:]...), updateSequence[i]), shouldBeAfter)
		}
	}
	return updateSequence
}

func contains(src []int, target []int) bool {
	for _, i := range src {
		for _, j := range target {
			if i == j {
				return true
			}
		}
	}
	return false
}
