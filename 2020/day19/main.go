package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Printf("Part 1. %d messages completely match rule 0\n", pt1(lines))
	fmt.Printf("Part 2. after updating rules 8 and 11, %d messages completely match rule\n", pt2(lines))
}

func pt1(lines []string) int {
	rules, messages := parseInput(lines)
	regPattern := "^" + createRegexPattern(0, rules) + "$"
	r := regexp.MustCompile(regPattern)

	var count int
	for _, m := range messages {
		if r.MatchString(m) {
			count++
		}
	}
	return count
}

func pt2(lines []string) int {
	rules, messages := parseInput(lines)
	rules[8] = "42 +"

	// emulate a recursive pattern for rule 11
	rulePattern := "42 31"
	rules[11] = rulePattern
	for i := 0; i < 3; i++ {
		rulePattern = "42 " + rulePattern + " 31"
		rules[11] += " | " + rulePattern
	}

	regPattern := "^" + createRegexPattern(0, rules) + "$"
	r := regexp.MustCompile(regPattern)

	var count int
	for _, m := range messages {
		if r.MatchString(m) {
			count++
		}
	}
	return count
}

func parseInput(lines []string) (map[int]string, []string) {
	var i int

	rules := make(map[int]string)
	for lines[i] != "" {
		s := strings.Split(lines[i], ":")
		idx, _ := strconv.Atoi(s[0])

		s[1] = strings.TrimPrefix(s[1], " ")
		s[1] = strings.TrimPrefix(s[1], "\"")

		rules[idx] = strings.TrimSuffix(s[1], "\"")
		i++
	}

	i++ // remove empty line

	var messages []string
	for i < len(lines) {
		messages = append(messages, lines[i])
		i++
	}

	return rules, messages
}

func createRegexPattern(i int, rules map[int]string) string {
	if rules[i] == "a" || rules[i] == "b" {
		return rules[i]
	}

	regex := "("
	for _, r := range strings.Split(rules[i], " ") {
		if r == "|" {
			regex += "|"
			continue
		}
		if r == "+" {
			regex += "+"
			continue
		}
		idx, _ := strconv.Atoi(r)
		regex += createRegexPattern(idx, rules)
	}

	return regex + ")"
}
