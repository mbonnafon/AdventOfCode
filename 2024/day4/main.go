package main

import (
	"fmt"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. :", pt1(lines))
	fmt.Println("Part 2. :", pt2(lines))
}

const target = "XMAS"

func pt1(lines []string) int {
	var count int
	horizontalSize := len(lines)
	verticalSize := len(lines[0])
	for i := 0; i < horizontalSize; i++ {
		for j := 0; j < verticalSize; j++ {
			if !(lines[i][j] == 'X') {
				continue
			}

			// check horizontal left
			if i >= 3 {
				if (lines[i-1][j] == 'M') && (lines[i-2][j] == 'A') && (lines[i-3][j] == 'S') {
					count++
				}
			}
			// check horizontal right
			if i < horizontalSize-3 {
				if (lines[i+1][j] == 'M') && (lines[i+2][j] == 'A') && (lines[i+3][j] == 'S') {
					count++
				}
			}
			// check vertical upper
			if j >= 3 {
				if (lines[i][j-1] == 'M') && (lines[i][j-2] == 'A') && (lines[i][j-3] == 'S') {
					count++
				}
			}
			// check vertical lower
			if j < verticalSize-3 {
				if (lines[i][j+1] == 'M') && (lines[i][j+2] == 'A') && (lines[i][j+3] == 'S') {
					count++
				}
			}
			// check diagonal upper left
			if j >= 3 && i >= 3 {
				if (lines[i-1][j-1] == 'M') && (lines[i-2][j-2] == 'A') && (lines[i-3][j-3] == 'S') {
					count++
				}
			}
			// check diagonal upper right
			if j >= 3 && i < horizontalSize-3 {
				if (lines[i+1][j-1] == 'M') && (lines[i+2][j-2] == 'A') && (lines[i+3][j-3] == 'S') {
					count++
				}
			}
			// check diagonal lower left
			if j < verticalSize-3 && i >= 3 {
				if (lines[i-1][j+1] == 'M') && (lines[i-2][j+2] == 'A') && (lines[i-3][j+3] == 'S') {
					count++
				}
			}
			// check diagonal lower right
			if j < verticalSize-3 && i < horizontalSize-3 {
				if (lines[i+1][j+1] == 'M') && (lines[i+2][j+2] == 'A') && (lines[i+3][j+3] == 'S') {
					count++
				}
			}
		}
	}

	return count
}

func pt2(lines []string) int {

	var count int
	horizontalSize := len(lines)
	verticalSize := len(lines[0])
	for i := 0; i < horizontalSize; i++ {
		for j := 0; j < verticalSize; j++ {
			if !(lines[i][j] == 'A') {
				continue
			}

			if i == 0 || j == 0 || i == horizontalSize-1 || j == verticalSize-1 {
				continue
			}

			// check M at top
			if (lines[i-1][j-1] == 'M') && (lines[i+1][j-1] == 'M') && (lines[i-1][j+1] == 'S') && (lines[i+1][j+1] == 'S') {
				count++
			}
			// check M at left
			if (lines[i-1][j-1] == 'M') && (lines[i-1][j+1] == 'M') && (lines[i+1][j-1] == 'S') && (lines[i+1][j+1] == 'S') {
				count++
			}
			// check M at bottom
			if (lines[i-1][j+1] == 'M') && (lines[i+1][j+1] == 'M') && (lines[i-1][j-1] == 'S') && (lines[i+1][j-1] == 'S') {
				count++
			}
			// check M at right
			if (lines[i+1][j-1] == 'M') && (lines[i+1][j+1] == 'M') && (lines[i-1][j-1] == 'S') && (lines[i-1][j+1] == 'S') {
				count++
			}
		}
	}

	return count
}
