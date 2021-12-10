package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. :", pt1(lines))
	fmt.Println("Part 2. :", pt2(lines))
}

func pt1(lines []string) int64 {
	var diagnosticReport [][]string
	for _, v := range lines {
		diagnosticReport = append(diagnosticReport, strings.Split(v, ""))
	}

	var gammaRates, epsilonRates string
	for _, binary := range helpers.Transpose(diagnosticReport) {
		if max(binary) == "0" {
			gammaRates = gammaRates + "0"
			epsilonRates = epsilonRates + "1"
			continue
		}
		gammaRates = gammaRates + "1"
		epsilonRates = epsilonRates + "0"
	}

	gammaRate, _ := strconv.ParseInt(gammaRates, 2, 64)
	epsilonRate, _ := strconv.ParseInt(epsilonRates, 2, 64)
	return gammaRate * epsilonRate
}

func pt2(lines []string) int64 {
	oxygen := make([]string, len(lines))
	CO2Scrubber := make([]string, len(lines))
	copy(oxygen, lines)
	copy(CO2Scrubber, lines)

	for i := 0; i < len(lines[0]); i++ {
		var matcher string
		switch one, zero := find(oxygen, i); {
		case zero > one:
			matcher = "1"
		case zero < one:
			matcher = "0"
		default:
			matcher = "1"
		}
		oxygen = extract(oxygen, i, matcher)
		if len(oxygen) == 1 {
			break
		}
	}

	for i := 0; i < len(lines[0]); i++ {
		var matcher string
		switch one, zero := find(oxygen, i); {
		case zero < one:
			matcher = "1"
		case zero > one:
			matcher = "0"
		default:
			matcher = "0"
		}
		CO2Scrubber = extract(CO2Scrubber, i, matcher)
		if len(CO2Scrubber) == 1 {
			break
		}
	}

	x, _ := strconv.ParseInt(oxygen[0], 2, 64)
	y, _ := strconv.ParseInt(CO2Scrubber[0], 2, 64)
	return x * y
}

func max(binary []string) string {
	var zeroCounter, oneCounter int
	for _, s := range binary {
		if s == "0" {
			zeroCounter++
			continue
		}
		oneCounter++
	}
	if zeroCounter > oneCounter {
		return "0"
	}
	return "1"
}

func find(lines []string, pos int) (int, int) {
	var one, zero int
	for _, l := range lines {
		if string(l[pos]) == "0" {
			zero++
		} else {
			one++
		}
	}
	return zero, one
}

func extract(lines []string, pos int, matcher string) []string {
	var newLines []string
	for _, l := range lines {
		li := l
		if string(l[pos]) == matcher {
			newLines = append(newLines, li)
		}
	}
	return newLines
}
