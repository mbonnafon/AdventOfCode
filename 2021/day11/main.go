package main

import (
	"fmt"
	"strings"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

type Coord struct {
	x, y int
}

type Grid struct {
	cells   [][]int
	flashed [][]bool
	width   int
	height  int
}

func (g *Grid) inGrid(i, j int) bool {
	return i >= 0 && i < g.width && j >= 0 && j < g.height
}

func (g *Grid) clearFlashed() {
	g.flashed = make([][]bool, g.height)
	for i := 0; i < g.width; i++ {
		g.flashed[i] = make([]bool, g.width)
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
	var dumboOctopuses [][]int
	for _, v := range lines {
		lineInt := helpers.ToIntSlice(strings.Split(v, ""))
		dumboOctopuses = append(dumboOctopuses, lineInt)
	}
	g := Grid{
		cells:  dumboOctopuses,
		height: len(dumboOctopuses),
		width:  len(dumboOctopuses[0]),
	}
	g.clearFlashed()
	fmt.Println("Part 1. :", pt1(100, g))
	fmt.Println("Part 2. :", pt2(500, g))
}

func pt1(steps int, dumboOctopuses Grid) (count int) {
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
	for i := 0; i < g.height; i++ {
		for j := 0; j < g.width; j++ {
			g.cells[i][j]++
		}
	}
}

func (g *Grid) step2() (count int) {
	for i := 0; i < g.height; i++ {
		for j := 0; j < g.width; j++ {
			if g.cells[i][j] > 9 {
				count += g.makeFlash(i, j)
			}
		}
	}
	return
}

func (g *Grid) makeFlash(i, j int) (count int) {
	count++
	g.cells[i][j] = 0
	g.flashed[i][j] = true
	for _, c := range g.getNeigh(i, j) {
		if g.cells[c.x][c.y] == 0 {
			continue
		}
		g.cells[c.x][c.y]++
		if g.cells[c.x][c.y] > 9 {
			g.cells[c.x][c.y] = 0
			count += g.makeFlash(c.x, c.y)
			continue
		}
	}
	return
}

func (g *Grid) getNeigh(i, j int) []Coord {
	if !g.inGrid(i, j) {
		return nil
	}

	neighbors := make([]Coord, 0)
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dy == 0 && dx == 0 {
				continue
			}

			p := Coord{i + dx, j + dy}
			if g.inGrid(p.x, p.y) {
				neighbors = append(neighbors, p)
			}
		}
	}

	return neighbors
}
