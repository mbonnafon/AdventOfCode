package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	grid := func([]string) Grid {
		var grid Grid
		for _, i := range lines {
			grid = append(grid, strings.Split(i, ""))
		}
		return grid
	}(lines)
	fmt.Println("Part 1. :", pt1(grid))
	fmt.Println("Part 2. :", pt2(grid))
}

type Grid [][]string

func pt1(grid Grid) int {
	var riskLevels int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			ref := grid[i][j]
			//up
			if (i > 0) && grid[i-1][j] <= ref {
				continue
			}
			//down
			if ((i + 1) < len(grid)) && grid[i+1][j] <= ref {
				continue
			}
			//left
			if (j > 0) && grid[i][j-1] <= ref {
				continue
			}
			//right
			if ((j + 1) < len(grid[i])) && grid[i][j+1] <= ref {
				continue
			}
			s, _ := strconv.Atoi(ref)
			riskLevels += (s + 1)
		}
	}
	return riskLevels
}

func pt2(grid Grid) int {
	var values []int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			s := grid.bassinRec(i, j)
			if s != 0 {
				values = append(values, s)
			}
		}
	}
	return func(values []int) int {
		sort.Ints(values)
		count := 1
		for _, v := range values[len(values)-3:] {
			count *= v
		}
		return count
	}(values)
}

// It's a DFS and mark nodes as visited with a 9
func (g *Grid) bassinRec(i, j int) (size int) {
	//up
	if (i > 0) && (*g)[i-1][j] != "9" {
		(*g)[i-1][j] = "9"
		s := g.bassinRec(i-1, j)
		size += s + 1
	}
	//down
	if ((i + 1) < len(*g)) && (*g)[i+1][j] != "9" {
		(*g)[i+1][j] = "9"
		s := g.bassinRec(i+1, j)
		size += s + 1
	}
	//left
	if (j > 0) && (*g)[i][j-1] != "9" {
		(*g)[i][j-1] = "9"
		s := g.bassinRec(i, j-1)
		size += s + 1
	}
	//right
	if ((j + 1) < len((*g)[i])) && (*g)[i][j+1] != "9" {
		(*g)[i][j+1] = "9"
		s := g.bassinRec(i, j+1)
		size += s + 1
	}
	return size
}
