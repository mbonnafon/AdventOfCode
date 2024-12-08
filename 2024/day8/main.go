package main

import (
	"fmt"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

type Coord struct {
	x int
	y int
}

func (c Coord) Extend(direction Coord) Coord {
	return Coord{c.x + direction.x, c.y + direction.y}
}

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. :", pt1(lines))
	fmt.Println("Part 2. :", pt2(lines))
}

type Vector struct {
	A         Coord
	B         Coord
	direction Coord
}

func NewVector(a, b Coord) Vector {
	coords := Coord{b.x - a.x, b.y - a.y}
	return Vector{a, b, coords}
}

func (v Vector) Extend() Coord {
	return Coord{v.B.x + v.direction.x, v.B.y + v.direction.y}
}

func parseAtennas(lines []string) (map[rune][]Coord, map[Coord]bool) {
	antennas := make(map[rune][]Coord)
	inMap := make(map[Coord]bool)
	for i, line := range lines {
		for j, letter := range line {
			inMap[Coord{i, j}] = true
			if letter == '.' {
				continue
			}
			antennas[letter] = append(antennas[letter], Coord{i, j})
		}
	}
	return antennas, inMap
}

func pt1(lines []string) int {
	antennas, inMap := parseAtennas(lines)
	antinodes := make(map[Coord]bool)
	for _, coords := range antennas {
		for _, a := range coords {
			for _, b := range coords {
				if a == b {
					continue
				}
				if pos := NewVector(a, b).Extend(); inMap[pos] {
					antinodes[pos] = true
				}
			}
		}
	}
	return len(antinodes)
}

func pt2(lines []string) int {
	antennas, inMap := parseAtennas(lines)
	antinodes := make(map[Coord]bool)
	for _, coords := range antennas {
		for _, a := range coords {
			for _, b := range coords {
				if a == b {
					continue
				}
				vector := NewVector(a, b)
				for node := vector.B; inMap[node]; node = node.Extend(vector.direction) {
					antinodes[node] = true
				}
			}
		}
	}
	return len(antinodes)
}
