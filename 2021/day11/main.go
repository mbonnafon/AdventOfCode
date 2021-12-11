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
	flashed [][]bool
}

func (g *Grid) clearFlashed() {
	g.flashed = make([][]bool, g.Height)
	for i := 0; i < g.Width; i++ {
		g.flashed[i] = make([]bool, g.Width)
	}
}

func (g *Grid) allFlashed() bool {
	for _, i := range g.flashed {
		for _, j := range i {
			if j == false {
				return false
			}
		}
	}
	return true
}

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. :", pt1(100, Grid{helpers.ParseGrid(lines), [][]bool{}}))
	fmt.Println("Part 2. :", pt2(500, Grid{helpers.ParseGrid(lines), [][]bool{}}))
}

func pt1(steps int, dumboOctopuses Grid) (count int) {
	dumboOctopuses.clearFlashed()
	for i := 0; i < steps; i++ {
		dumboOctopuses.step1()
		count += dumboOctopuses.step2()
	}
	return
}

func pt2(steps int, dumboOctopuses Grid) int {
	for i := 1; i < steps; i++ {
		dumboOctopuses.clearFlashed()
		dumboOctopuses.step1()
		dumboOctopuses.step2()
		if dumboOctopuses.allFlashed() {
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
	g.flashed[i][j] = true
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
