package main

import (
	"fmt"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

func seatFromBoardingPass(pass string) int {
	var columnPlace, rowPlace int

	low, high := 0, 127
	for _, c := range pass[:len(pass)-3] {
		columnPlace = low + high
		if string(c) == "F" {
			columnPlace = columnPlace >> 1
			high = columnPlace
		} else if string(c) == "B" {
			columnPlace = columnPlace>>1 + 1
			low = columnPlace
		}
	}
	low, high = 0, 7
	for _, c := range pass[len(pass)-3:] {
		rowPlace = low + high
		if string(c) == "L" {
			rowPlace = rowPlace >> 1
			high = rowPlace
		} else if string(c) == "R" {
			rowPlace = rowPlace>>1 + 1
			low = rowPlace
		}
	}
	return columnPlace*8 + rowPlace
}

func pt1(lines []string) int {
	var high int
	for _, l := range lines {
		seatID := seatFromBoardingPass(l)
		if seatID > high {
			high = seatID
		}
	}
	return high
}

func pt2(lines []string) int {
	var yourSeat int
	seats := make(map[int]bool, len(lines))

	for _, l := range lines {
		seatID := seatFromBoardingPass(l)
		seats[seatID] = true
	}

	for seat := range seats {
		if seats[seat+1] && !seats[seat+2] && seats[seat+3] {
			yourSeat = seat + 2
		}
	}
	return yourSeat
}

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. highest seat ID on a boarding pass is:", pt1(lines))
	fmt.Println("Part 2. ID of your seat is:", pt2(lines))
}
