package main

import (
	"fmt"
	"strconv"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

type Tuple struct {
	Zero int
	One  int
}

func (t Tuple) gamma() string {
	if t.One > t.Zero {
		return "1"
	} else {
		return "0"
	}
}

func (t Tuple) epsilon() string {
	if t.One > t.Zero {
		return "0"
	} else {
		return "1"
	}
}

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. :", pt1(lines))
	fmt.Println("Part 2. :", pt2(lines))
}

func pt1(lines []string) int64 {
	acc := make(map[int]*Tuple)
	for _, l := range lines {
		for b, i := range l {
			if acc[b] == nil {
				acc[b] = &Tuple{}
			}
			if string(i) == "0" {
				acc[b].Zero++
			} else {
				acc[b].One++
			}
		}
	}
	var gammaRates, epsilonRates string
	for i := 0; i < len(acc); i++ {
		gammaRates = gammaRates + acc[i].gamma()
		epsilonRates = epsilonRates + acc[i].epsilon()
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
