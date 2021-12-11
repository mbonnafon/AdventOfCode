package main

import (
	"fmt"
	"sort"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

type Grid struct {
	helpers.Grid
}

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	grid := Grid{helpers.ParseGrid(lines)}
	fmt.Println("Part 1. :", pt1(grid))
	fmt.Println("Part 2. :", pt2(grid))
}

func pt1(grid Grid) int {
	var riskLevels int
	for i := 0; i < grid.Width; i++ {
		for j := 0; j < grid.Height; j++ {
			if grid.IsSmallestComparedToNeigh(i, j) {
				riskLevels += grid.Cells[i][j] + 1
			}
		}
	}
	return riskLevels
}

func pt2(grid Grid) int {
	var totalSize []int
	for i := 0; i < grid.Width; i++ {
		for j := 0; j < grid.Height; j++ {
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

// It's a DFS and mark nodes as visited with a 9
func (g *Grid) bassinRec(i, j int) (size int) {
	//up
	if (i > 0) && g.Cells[i-1][j] != 9 {
		g.Cells[i-1][j] = 9
		s := g.bassinRec(i-1, j)
		size += s + 1
	}
	//down
	if ((i + 1) < g.Height) && g.Cells[i+1][j] != 9 {
		g.Cells[i+1][j] = 9
		s := g.bassinRec(i+1, j)
		size += s + 1
	}
	//left
	if (j > 0) && g.Cells[i][j-1] != 9 {
		g.Cells[i][j-1] = 9
		s := g.bassinRec(i, j-1)
		size += s + 1
	}
	//right
	if ((j + 1) < g.Width) && g.Cells[i][j+1] != 9 {
		g.Cells[i][j+1] = 9
		s := g.bassinRec(i, j+1)
		size += s + 1
	}
	return size
}
