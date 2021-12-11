package main

import (
	"fmt"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

type Coord struct {
	x, y int
}

type Grid struct {
	helpers.Grid
}

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. :", pt1(100, Grid{helpers.ParseGrid(lines)}))
	fmt.Println("Part 2. :", pt2(500, Grid{helpers.ParseGrid(lines)}))
}

func pt1(steps int, grid Grid) (count int) {
	for i := 0; i < steps; i++ {
		grid.step1()
		count += grid.step2()
	}
	return
}

func pt2(steps int, grid Grid) int {
	for i := 1; i < steps; i++ {
		grid.step1()
		count := grid.step2()
		if count == grid.Width*grid.Height {
			return i
		}
	}
	return 0
}

func (g *Grid) step1() {
	for i := 0; i < g.Height; i++ {
		for j := 0; j < g.Width; j++ {
			g.Cells[i][j]++
		}
	}
}

func (g *Grid) step2() (count int) {
	for i := 0; i < g.Height; i++ {
		for j := 0; j < g.Width; j++ {
			if g.Cells[i][j] > 9 {
				count += g.makeFlash(i, j)
			}
		}
	}
	return
}

func (g *Grid) makeFlash(i, j int) (count int) {
	count++
	g.Cells[i][j] = 0
	for _, c := range g.GetNeigh(i, j) {
		x, y := c.X, c.Y
		if g.Cells[x][y] == 0 {
			continue
		}
		g.Cells[x][y]++
		if g.Cells[x][y] > 9 {
			g.Cells[x][y] = 0
			count += g.makeFlash(x, y)
			continue
		}
	}
	return
}
