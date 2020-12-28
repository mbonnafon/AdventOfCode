package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

var rmask = regexp.MustCompile(`mask = ([10X]+)`)
var rmem = regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. the sum of all values left in memory after it completes is:", pt1(lines))
	fmt.Println("Part 2. :", pt2(lines))
}

type memory struct {
	index int64
	value int64
}

type computerSystem struct {
	mask string
	mems []memory
}

func pt1(lines []string) int {
	computers := parseComputers(lines)
	memory := make(map[int64]int64)
	for _, c := range computers {
		c.compute(memory, 1)
	}
	var sum int
	for _, v := range memory {
		sum += int(v)
	}
	return sum
}

func pt2(lines []string) int {
	computers := parseComputers(lines)
	memory := make(map[int64]int64)
	for _, c := range computers {
		c.compute(memory, 2)
	}
	var sum int
	for _, v := range memory {
		sum += int(v)
	}
	return sum
}

func parseComputers(linesOfComputers []string) []computerSystem {
	var computers []computerSystem
	var computer computerSystem
	for i, c := range linesOfComputers {
		if mask := rmask.FindStringSubmatch(c); mask != nil {
			if i > 0 {
				computers = append(computers, computer)
				computer = computerSystem{}
			}
			computer.mask = mask[1]
		}
		if mems := rmem.FindStringSubmatch(c); mems != nil {
			computer.mems = append(computer.mems, createMemory(mems[1], mems[2]))
		}
	}
	computers = append(computers, computer)
	return computers
}

func createMemory(index, value string) memory {
	i, _ := strconv.ParseInt(index, 10, 64)
	v, _ := strconv.ParseInt(value, 10, 64)
	return memory{i, v}
}

func (c computerSystem) compute(memory map[int64]int64, version int) {
	for _, m := range c.mems {
		if version == 1 {
			b := []byte(fmt.Sprintf("%036b", m.value))
			for i := len(c.mask) - 1; i >= 0; i-- {
				if c.mask[i] != 'X' {
					b[i] = c.mask[i]
				}
			}
			bint, _ := strconv.ParseInt(string(b), 2, 64)
			memory[m.index] = bint

		} else {
			writeFloatingMem(memory, c.mask, m.index, m.value)
		}
	}
}

func writeFloatingMem(memory map[int64]int64, mask string, address, value int64) {
	result := []byte(fmt.Sprintf("%036b", address))
	var count int
	for i := len(mask) - 1; i >= 0; i-- {
		if mask[i] == 'X' {
			result[i] = byte('X')
		}
		if mask[i] == '1' {
			result[i] = '1'
		}
		if result[i] == '1' {
			count += powInt(2, len(mask)-i-1)
		}
	}
	//result is now the floating address
	var floatingIndex []int
	for i := len(result) - 1; i >= 0; i-- {
		if result[i] == 'X' {
			if len(floatingIndex) == 0 {
				floatingIndex = append(floatingIndex, count, count+powInt(2, len(mask)-i-1))
				continue
			}
			var buffIndex []int
			for _, index := range floatingIndex {
				buffIndex = append(buffIndex, index+powInt(2, len(mask)-i-1))
			}
			floatingIndex = append(floatingIndex, buffIndex...)
		}
	}
	//add value to address
	for _, v := range floatingIndex {
		memory[int64(v)] = value
	}
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
