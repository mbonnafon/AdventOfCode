package main

import (
	"fmt"
	"strings"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

type coord helpers.Coord

type Fold struct {
	order string
	size  int
}

type Grid struct {
	coords     []coord
	foldCoords []Fold
	Xmax       int
	Ymax       int
}

func (g *Grid) calcMax() {
	var maxX, maxY int
	for _, c := range g.coords {
		if c.X > maxX {
			maxX = c.X
		}
		if c.Y > maxY {
			maxY = c.Y
		}
	}
	g.Xmax = maxX
	g.Ymax = maxY
}

func newGrid(lines []string) Grid {
	g := Grid{}

	var i int
	for i = 0; i < len(lines); i++ {
		if lines[i] == "" {
			i++
			break
		}
		c := strings.Split(lines[i], ",")
		g.coords = append(g.coords, coord{X: helpers.ToInt(c[0]), Y: helpers.ToInt(c[1])})
	}

	var foldCoords []Fold
	for i = i; i < len(lines); i++ {
		var matched string
		r := strings.NewReader(lines[i])
		fmt.Fscanf(r, "fold along %s", &matched)
		m := strings.Split(matched, "=")
		foldCoords = append(foldCoords, Fold{order: m[0], size: helpers.ToInt(m[1])})
	}
	g.foldCoords = foldCoords
	g.calcMax()
	return g
}

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. :", pt1(newGrid(lines)))
	fmt.Println("Part 2. :", pt2(newGrid(lines)))
}

func pt1(g Grid) int {
	size := g.foldCoords[0].size
	if g.foldCoords[0].order == "x" {
		g.foldX(size)
	} else {
		g.foldY(size)
	}
	count := make(map[coord]bool)
	for _, v := range g.coords {
		count[v] = true
	}
	return len(count)
}

func pt2(g Grid) int {
	for _, f := range g.foldCoords {
		size := f.size
		if f.order == "x" {
			g.foldX(size)
		} else {
			g.foldY(size)
		}
	}
	display := make([][]bool, g.Ymax)
	for _, v := range g.coords {
		if display[v.Y] == nil {
			display[v.Y] = make([]bool, g.Xmax)
		}
		display[v.Y][v.X] = true
	}
	for i := 0; i < g.Ymax; i++ {
		for j := 0; j < g.Xmax; j++ {
			if display[i][j] == true {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	count := make(map[coord]bool)
	for _, v := range g.coords {
		count[v] = true
	}
	return len(count)
}

func (g *Grid) foldX(x int) {
	var newCoords []coord
	for _, c := range g.coords {
		if c.X > x {
			newX := x - (c.X - x)
			if newX < 0 {
				continue
			}
			newC := coord{
				X: x - (c.X - x),
				Y: c.Y,
			}
			newCoords = append(newCoords, newC)
			continue
		}
		newCoords = append(newCoords, c)
	}
	g.coords = newCoords
	g.Xmax = x
}

func (g *Grid) foldY(y int) {
	var newCoords []coord
	for _, c := range g.coords {
		if c.Y > y {
			newY := y - (c.Y - y)
			if newY < 0 {
				continue
			}
			newC := coord{
				X: c.X,
				Y: y - (c.Y - y),
			}
			newCoords = append(newCoords, newC)
			continue
		}
		newCoords = append(newCoords, c)
	}
	g.coords = newCoords
	g.Ymax = y
}
