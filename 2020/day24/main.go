package main

import (
	"fmt"
	"strings"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

var directions = map[string]tile{
	"e":  {1, -1, 0},
	"se": {1, 0, -1},
	"sw": {0, 1, -1},
	"w":  {-1, 1, 0},
	"nw": {-1, 0, 1},
	"ne": {0, -1, 1},
}

type tile struct {
	x, y, z int
}

func (t tile) neighboor() []tile {
	var neighboor []tile
	for dir := range directions {
		nb := t
		nb.Move(dir)
		neighboor = append(neighboor, nb)
	}
	return neighboor
}

// https://catlikecoding.com/unity/tutorials/hex-map/part-1/
// https://www.redblobgames.com/grids/hexagons/#neighbors
func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Printf("Part 1. %d tiles are left with the black side up\n", pt1(lines))
	fmt.Printf("Part 2. %d tiles will be black after 100 days\n", pt2(lines))
}

func pt1(lines []string) int {
	return len(flip(lines))
}

func pt2(lines []string) int {
	blackTiles := flip(lines)

	for i := 0; i < 100; i++ {
		adjacentsTiles := make(map[tile]int)
		for tile := range blackTiles {
			for _, neighboor := range tile.neighboor() {
				adjacentsTiles[neighboor]++
			}
		}

		newTiles := make(map[tile]struct{})
		for pos, adjacentBlackTiles := range adjacentsTiles {
			if _, isBlack := blackTiles[pos]; (isBlack && adjacentBlackTiles == 1) || adjacentBlackTiles == 2 {
				newTiles[pos] = struct{}{}
			}
		}

		blackTiles = newTiles
	}

	return len(blackTiles)
}

func flip(lines []string) map[tile]struct{} {
	tiles := make(map[tile]struct{})
	for _, l := range lines {
		tile := moveHex(l)
		if _, ok := tiles[tile]; ok {
			delete(tiles, tile)
			continue
		}
		tiles[tile] = struct{}{}
	}
	return tiles
}

func moveHex(move string) tile {
	var tile tile
	var dir string
	for i := 0; i < len(move); i++ {
		dir = string(move[i])
		if !strings.HasPrefix(dir, "e") && !strings.HasPrefix(dir, "w") {
			i++
			dir = dir + string(move[i])
		}
		tile.Move(dir)
	}
	return tile
}

func (t *tile) Move(dir string) {
	t.x += directions[dir].x
	t.y += directions[dir].y
	t.z += directions[dir].z
}
