package main

import (
	"fmt"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

type trails map[Coord]*int

func NewTrails(lines []string) (trails, []Coord) {
	trails := make(trails)
	start := make([]Coord, 0)
	for i, line := range lines {
		for j, pos := range line {
			heihgt := helpers.ToInt(string(pos))
			c := Coord{i, j}
			if heihgt == 0 {
				start = append(start, c)
			}
			trails[c] = &heihgt
		}
	}
	return trails, start
}

func (t trails) GetLevel(c Coord) *int {
	return t[c]
}

func (t trails) NeighboorOf(c Coord) []Coord {
	neigh := make([]Coord, 0)
	if upper := (Coord{c.x, c.y - 1}); t[upper] != nil {
		neigh = append(neigh, upper)
	}
	if right := (Coord{c.x + 1, c.y}); t[right] != nil {
		neigh = append(neigh, right)
	}
	if bottom := (Coord{c.x, c.y + 1}); t[bottom] != nil {
		neigh = append(neigh, bottom)
	}
	if left := (Coord{c.x - 1, c.y}); t[left] != nil {
		neigh = append(neigh, left)
	}
	return neigh
}

type Coord struct {
	x int
	y int
}

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. :", pt1(lines))
	fmt.Println("Part 2. :", pt2(lines))
}

func (t trails) countTrails(start Coord) []Coord {
	startHeight := *t.GetLevel(start)
	if startHeight == 9 {
		return []Coord{start}
	}
	var count []Coord
	for _, coord := range t.NeighboorOf(start) {
		neighHeight := *t.GetLevel(coord)
		if startHeight+1 == neighHeight {
			count = append(count, t.countTrails(coord)...)
		}
	}
	return count
}

func pt1(lines []string) int {
	trails, start := NewTrails(lines)

	var totalScore int
	for _, trailStart := range start {
		uniqueEnd := make(map[Coord]bool)
		for _, endCoord := range trails.countTrails(trailStart) {
			uniqueEnd[endCoord] = true
		}
		totalScore += len(uniqueEnd)
	}
	return totalScore
}

func pt2(lines []string) int {
	trails, start := NewTrails(lines)

	var totalScore int
	for _, trailStart := range start {
		totalScore += len(trails.countTrails(trailStart))
	}
	return totalScore
}
