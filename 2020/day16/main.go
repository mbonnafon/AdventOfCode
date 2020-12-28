package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

type rules map[string][]Interval

type ticket []int

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. your ticket scanning error rate:", pt1(lines))
	fmt.Println("Part 2. if you multiply the six fields on your ticket that start with the word departure:", pt2(lines))
}

func pt1(lines []string) int {
	rules, _, otherTickets := parseInput(lines)

	var intervals []Interval
	for _, v := range rules {
		intervals = append(intervals, v...)
	}
	interval := merge(intervals)
	invalid := findInvalidValues(otherTickets, interval)

	var count int
	for _, v := range invalid {
		count += v
	}
	return count
}

func pt2(lines []string) int {
	rules, myTicket, otherTickets := parseInput(lines)

	var intervals []Interval
	for _, v := range rules {
		intervals = append(intervals, v...)
	}
	interval := merge(intervals)

	otherTickets = removeInvalidTickets(append(otherTickets, myTicket), interval)
	corr := findTicketCorrespondance(rules, otherTickets)
	return calcMyTicket(myTicket, corr)
}

func parseInput(lines []string) (rules, ticket, []ticket) {
	var i, min, max, minBis, maxBis int
	var f1, f2 string

	rules := make(rules)
	for lines[i] != "" {
		_, err := fmt.Sscanf(lines[i], "%s %s %d-%d or %d-%d", &f1, &f2, &min, &max, &minBis, &maxBis)
		f1 += f2
		if err != nil {
			fmt.Sscanf(lines[i], "%s %d-%d or %d-%d", &f1, &min, &max, &minBis, &maxBis)
		}
		rules[strings.TrimSuffix(f1, ":")] = []Interval{{min, max}, {minBis, maxBis}}
		i++
	}

	myTicket := helpers.StringSliceToIntSlice(strings.Split(lines[i+2], ","))
	i += 5

	var otherTickets []ticket
	for i < len(lines) {
		otherTickets = append(otherTickets, helpers.StringSliceToIntSlice(strings.Split(lines[i], ",")))
		i++
	}
	return rules, myTicket, otherTickets
}

func findInvalidValues(tickets []ticket, intervals []Interval) []int {
	var invalid []int
	for _, t := range tickets {
		for _, n := range t {
			var valid bool
			for _, interval := range intervals {
				if helpers.IntInInterval(interval.min, interval.max, n) {
					valid = true
					break
				}
			}
			if !valid {
				invalid = append(invalid, n)
			}
		}
	}
	return invalid
}

func removeInvalidTickets(tickets []ticket, intervals []Interval) []ticket {
	var validTickets []ticket
	for _, ticket := range tickets {
		valid := true
		for _, field := range ticket {
			var fieldValidity bool
			for _, interval := range intervals {
				if helpers.IntInInterval(interval.min, interval.max, field) {
					fieldValidity = true
					break
				}
			}
			if !fieldValidity {
				valid = false
			}
		}
		if valid {
			validTickets = append(validTickets, ticket)
		}
	}
	return validTickets
}

func findTicketCorrespondance(rules rules, tickets []ticket) map[string]*int {
	columnLen := len(tickets[0])
	ruleValidForColumns := make(map[string][]int)

	for name, interval := range rules {
		for i := 0; i < columnLen; i++ {
			var valid int
			for j := 0; j < len(tickets); j++ {
				if helpers.IntInInterval(interval[0].min, interval[0].max, tickets[j][i]) {
					valid++
				}
				if helpers.IntInInterval(interval[1].min, interval[1].max, tickets[j][i]) {
					valid++
				}
			}
			if valid == len(tickets) {
				ruleValidForColumns[name] = append(ruleValidForColumns[name], i)
			}
		}
	}

	alreadyFound := make(map[int]bool)
	ticketCorresp := make(map[string]*int)
	rulesLen := len(rules)
	count := 1
	for len(ticketCorresp) <= rulesLen-1 {
		for name, interval := range ruleValidForColumns {
			if len(interval) == count {
				var currentInterval int
				for _, i := range interval {
					if !alreadyFound[i] {
						currentInterval = i
						alreadyFound[currentInterval] = true
						break
					}
				}
				ticketCorresp[name] = &currentInterval
				delete(rules, name)
				count++
			}
		}
	}
	return ticketCorresp
}

func calcMyTicket(myTicket ticket, corr map[string]*int) int {
	count := 1
	count *= myTicket[*corr["departuredate"]]
	count *= myTicket[*corr["departurelocation"]]
	count *= myTicket[*corr["departureplatform"]]
	count *= myTicket[*corr["departurestation"]]
	count *= myTicket[*corr["departuretime"]]
	count *= myTicket[*corr["departuretrack"]]
	return count
}

type Interval struct {
	min, max int
}

func merge(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].min < intervals[j].min
	})

	var r []Interval
	start, end := intervals[0].min, intervals[0].max
	for _, v := range intervals {
		if v.min <= end {
			if v.max > end {
				end = v.max
			}
		} else {
			r = append(r, Interval{min: start, max: end})
			start = v.min
			end = v.max
		}
	}
	r = append(r, Interval{min: start, max: end})

	return r
}
