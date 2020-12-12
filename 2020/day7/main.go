package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

type Bag struct {
	color string
	count int
}

func main() {
	const target = "shiny gold"
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Printf("Part 1. One %s can hold %d bags colors\n", target, pt1(lines, target))
	fmt.Printf("Part 2. %d individual bags are required inside your single %s\n", pt2(lines, target), target)
}

func pt1(lines []string, target string) int {
	var count int
	bags := parseBags(lines)
	for bag := range bags {
		if canHoldTarget(bags, bag, target) {
			count++
			continue
		}
	}
	return count
}

func pt2(lines []string, target string) int {
	bags := parseBags(lines)
	return countHoldableTarget(bags, target, target) - 1
}

func parseBags(lines []string) map[string][]Bag {
	bagsContent := make(map[string][]Bag)
	for _, l := range lines {
		trimmed := strings.TrimSuffix(l, ".")
		s := strings.Split(trimmed, " bags contain ")
		if s[1] == "no other bags" {
			bagsContent[s[0]] = nil
			continue
		}

		for _, bag := range strings.Split(s[1], ", ") {
			bagsContent[s[0]] = append(bagsContent[s[0]], bagFromString(bag))
		}
	}
	return bagsContent
}

func canHoldTarget(bags map[string][]Bag, current, target string) bool {
	for _, bag := range bags[current] {
		if bag.color == target && bag.count >= 1 {
			return true
		}
		if canHoldTarget(bags, bag.color, target) {
			return true
		}
	}
	return false
}

func bagFromString(s string) Bag {
	b := strings.Split(s, " ")
	i, _ := strconv.Atoi(b[0])
	return Bag{
		color: strings.Join(b[1:len(b)-1], " "),
		count: i,
	}
}

func countHoldableTarget(bags map[string][]Bag, current, target string) int {
	count := 1
	for _, bag := range bags[current] {
		count = count + bag.count*countHoldableTarget(bags, bag.color, target)
	}
	return count
}
