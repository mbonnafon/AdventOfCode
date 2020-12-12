package main

import (
	"fmt"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Printf("Part 1. %d seats end up occupied\n", pt1(lines))
	fmt.Printf("Part 2. %d seats end up occupied\n", pt2(lines))
}

func pt1(lines []string) int {
	var i int
	for {
		i++
		seats := populateSeats(countAdjacentOccupied, lines, 4)
		if seats == nil {
			return occupiedSeats(lines)
		}
		lines = seats
	}
	return 0
}

func pt2(lines []string) int {
	var i int
	for {
		i++
		seats := populateSeats(countVisibleOccupied, lines, 5)
		if seats == nil {
			return occupiedSeats(lines)
		}
		lines = seats
	}
	return 0
}

func populateSeats(countOccupiedSeats func([]string, int, int) int, seats []string, tolerance int) []string {
	var changes bool
	var columnBuff []string
	for irow, row := range seats {
		var rowBuff string
		for icolumn := range row {
			switch string(seats[irow][icolumn]) {
			case "L":
				if countOccupiedSeats(seats, irow, icolumn) == 0 {
					rowBuff = rowBuff + "#"
					changes = true
				} else {
					rowBuff = rowBuff + "L"
				}
			case "#":
				if countOccupiedSeats(seats, irow, icolumn) >= tolerance {
					rowBuff = rowBuff + "L"
					changes = true
				} else {
					rowBuff = rowBuff + "#"
				}
			default:
				rowBuff = rowBuff + string(seats[irow][icolumn])
			}
		}
		//fmt.Println(rowBuff)
		columnBuff = append(columnBuff, rowBuff)
	}
	if !changes {
		return nil
	}
	return columnBuff
}

func occupiedSeats(seats []string) int {
	var count int
	for irow, row := range seats {
		for icolumn := range row {
			if string(seats[irow][icolumn]) == "#" {
				count++
			}
		}
	}
	return count
}

func countAdjacentOccupied(lines []string, r, c int) int {
	var count int
	rowLen, colLen := len(lines)-1, len(lines[r])-1

	// up/down/left/right
	if r > 0 && string(lines[r-1][c]) == "#" {
		count++
	}
	if r < rowLen && string(lines[r+1][c]) == "#" {
		count++
	}
	if c > 0 && string(lines[r][c-1]) == "#" {
		count++
	}
	if c < colLen && string(lines[r][c+1]) == "#" {
		count++
	}

	// diagonals: upLeft/upRight/downLeft/downRight
	if r > 0 && c > 0 && string(lines[r-1][c-1]) == "#" {
		count++
	}
	if r > 0 && c < colLen && string(lines[r-1][c+1]) == "#" {
		count++
	}
	if r < rowLen && c > 0 && string(lines[r+1][c-1]) == "#" {

		count++
	}
	if r < rowLen && c < colLen && string(lines[r+1][c+1]) == "#" {
		count++
	}

	return count
}

func countVisibleOccupied(lines []string, r, c int) int {
	var count int
	rowLen, colLen := len(lines), len(lines[r])

	// up/down/left/right
	for ir := r - 1; ir >= 0; ir-- {
		if string(lines[ir][c]) == "#" {
			count++
			break
		}
		if string(lines[ir][c]) == "L" {
			break
		}
	}
	for ir := r + 1; ir < rowLen; ir++ {
		if string(lines[ir][c]) == "#" {
			count++
			break
		}
		if string(lines[ir][c]) == "L" {
			break
		}
	}
	for ic := c - 1; ic >= 0; ic-- {
		if string(lines[r][ic]) == "#" {
			count++
			break
		}
		if string(lines[r][ic]) == "L" {
			break
		}
	}
	for ic := c + 1; ic < colLen; ic++ {
		if string(lines[r][ic]) == "#" {
			count++
			break
		}
		if string(lines[r][ic]) == "L" {
			break
		}
	}

	// diagonals: upLeft/upRight/downLeft/downRight
	for ir, ic := r-1, c-1; ir >= 0 && ic >= 0; ir, ic = ir-1, ic-1 {
		if string(lines[ir][ic]) == "#" {
			count++
			break
		}
		if string(lines[ir][ic]) == "L" {
			break
		}
	}
	for ir, ic := r-1, c+1; ir >= 0 && ic < colLen; ir, ic = ir-1, ic+1 {
		if string(lines[ir][ic]) == "#" {
			count++
			break
		}
		if string(lines[ir][ic]) == "L" {
			break
		}
	}
	for ir, ic := r+1, c-1; ir < rowLen && ic >= 0; ir, ic = ir+1, ic-1 {
		if string(lines[ir][ic]) == "#" {
			count++
			break
		}
		if string(lines[ir][ic]) == "L" {
			break
		}
	}
	for ir, ic := r+1, c+1; ir < rowLen && ic < colLen; ir, ic = ir+1, ic+1 {
		if string(lines[ir][ic]) == "#" {
			count++
			break
		}
		if string(lines[ir][ic]) == "L" {
			break
		}
	}

	return count
}
