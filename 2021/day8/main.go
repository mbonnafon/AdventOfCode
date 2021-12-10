package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

type Entry struct {
	SignalPatterns []string
	DigitOutputs   []string
}

func (e Entry) isPresent(s string) bool {
	for _, digit := range e.SignalPatterns {
		if s == digit {
			return true
		}
	}
	return false
}

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. :", pt1(lines))
	fmt.Println("Part 2. :", pt2(lines))
}

func pt1(lines []string) int {
	counter := make(map[int]int)
	for _, entry := range parseEntries(lines) {
		for _, digit := range entry.DigitOutputs {
			if !entry.isPresent(digit) {
				continue
			}
			switch len(digit) {
			case 2:
				counter[2]++
			case 3:
				counter[3]++
			case 4:
				counter[4]++
			case 7:
				counter[7]++
			}
		}
	}
	return func(c map[int]int) (counter int) {
		for _, v := range c {
			counter += v
		}
		return
	}(counter)
}

func pt2(lines []string) int {
	var score int
	for _, entry := range parseEntries(lines) {
		wires := make(map[int][]string)
		for _, digit := range entry.SignalPatterns {
			wires[len(digit)] = append(wires[len(digit)], digit)
		}
		digitCorrespondance := computeDigitCorrespondance(wires)

		var outputValue string
		for _, v := range entry.DigitOutputs {
			outputValue = fmt.Sprintf("%s%d", outputValue, digitCorrespondance[v])
		}
		score += func(s string) int { i, _ := strconv.Atoi(s); return i }(outputValue)
	}
	return score
}

func computeDigitCorrespondance(wires map[int][]string) map[string]int {
	one, four := wires[2][0], wires[4][0]

	digitCorrespondance := make(map[string]int)
	digitCorrespondance[one] = 1
	digitCorrespondance[four] = 4
	digitCorrespondance[wires[3][0]] = 7
	digitCorrespondance[wires[7][0]] = 8
	for _, v := range wires[5] {
		if isThree(v, one) {
			digitCorrespondance[v] = 3
			continue
		}
		if isFive(v, one, four) {
			digitCorrespondance[v] = 5
			continue
		}
		digitCorrespondance[v] = 2
	}
	for _, v := range wires[6] {
		if isSix(v, one) {
			digitCorrespondance[v] = 6
			continue
		}
		if isNine(v, one, four) {
			digitCorrespondance[v] = 9
			continue
		}
		digitCorrespondance[v] = 0
	}
	return digitCorrespondance
}

func isThree(v, one string) bool {
	for _, c := range one {
		if !strings.Contains(v, string(c)) {
			return false
		}
	}
	return true
}

func isFive(v, one, four string) bool {
	if strings.Contains(v, string(one[0])) && strings.Contains(v, string(one[1])) {
		return false
	}
	for _, c := range four {
		if string(c) == string(one[0]) || string(c) == string(one[1]) {
			continue
		}
		if !strings.Contains(v, string(c)) {
			return false
		}
	}
	return true
}

func isSix(v, one string) bool {
	if strings.Contains(v, string(one[0])) && strings.Contains(v, string(one[1])) {
		return false
	}
	return true
}

func isNine(v, one, four string) bool {
	if !strings.Contains(v, string(one[0])) || !strings.Contains(v, string(one[1])) {
		return false
	}
	for _, c := range four {
		if !strings.Contains(v, string(c)) {
			return false
		}
	}
	return true
}

func parseEntries(lines []string) []Entry {
	var entries []Entry
	for _, l := range lines {
		entry := Entry{}
		s := strings.Split(l, "|")
		signal, output := strings.Split(strings.Trim(s[0], " "), " "), strings.Split(strings.Trim(s[1], " "), " ")
		for _, v := range signal {
			entry.SignalPatterns = append(entry.SignalPatterns, helpers.SortString(v))
		}
		for _, v := range output {
			entry.DigitOutputs = append(entry.DigitOutputs, helpers.SortString(v))
		}
		entries = append(entries, entry)
	}
	return entries
}
