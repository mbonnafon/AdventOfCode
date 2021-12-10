package main

import (
	"fmt"
	"sort"
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

func pt1(grid Grid) int {
	var riskLevels int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid.isSmallestComparedToNeigh(i, j) {
				riskLevels += (helpers.ToInt(grid[i][j]) + 1)
			}
		}
	}
	return riskLevels
}

func pt2(grid Grid) int {
	var totalSize []int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if size := grid.bassinRec(i, j); size != 0 {
				totalSize = append(totalSize, size)
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
	}(totalSize)
}

type Grid [][]string

func (g Grid) isSmallestComparedToNeigh(i, j int) bool {
	ref := g[i][j]
	//up
	if (i > 0) && g[i-1][j] <= ref {
		return false
	}
	//down
	if ((i + 1) < len(g)) && g[i+1][j] <= ref {
		return false
	}
	//left
	if (j > 0) && g[i][j-1] <= ref {
		return false
	}
	//right
	if ((j + 1) < len(g[i])) && g[i][j+1] <= ref {
		return false
	}
	return true
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
