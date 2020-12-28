package main

import (
	"fmt"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Printf("Part 1. %d cubes are left in the active state after the sixth cycle\n", pt1(lines))
	fmt.Printf("Part 2. %d cubes are left in the active state after the sixth cycle\n", pt2(lines))
}

func pt1(lines []string) int {
	activatedCube := make(map[helpers.Vector3D]struct{})
	for x, l := range lines {
		for y, c := range l {
			if string(c) == "#" {
				activatedCube[helpers.Vector3D{x, y, 0}] = struct{}{}
			}
		}
	}
	for i := 0; i < 6; i++ {
		adjacentCube := make(map[helpers.Vector3D]int)
		for cube := range activatedCube {
			for _, neighboor := range cube.Neighboor() {
				adjacentCube[neighboor]++
			}
		}

		newTiles := make(map[helpers.Vector3D]struct{})
		for pos, adjacentActivatedCube := range adjacentCube {
			if _, isActivated := activatedCube[pos]; (isActivated && adjacentActivatedCube == 2) || adjacentActivatedCube == 3 {
				newTiles[pos] = struct{}{}
			}
		}

		activatedCube = newTiles
	}
	return len(activatedCube)
}

func pt2(lines []string) int {
	activatedCube := make(map[helpers.Vector4D]struct{})
	for x, l := range lines {
		for y, c := range l {
			if string(c) == "#" {
				activatedCube[helpers.Vector4D{x, y, 0, 0}] = struct{}{}
			}
		}
	}
	for i := 0; i < 6; i++ {
		adjacentCube := make(map[helpers.Vector4D]int)
		for cube := range activatedCube {
			for _, neighboor := range cube.Neighboor() {
				adjacentCube[neighboor]++
			}
		}

		newTiles := make(map[helpers.Vector4D]struct{})
		for pos, adjacentActivatedCube := range adjacentCube {
			if _, isActivated := activatedCube[pos]; (isActivated && adjacentActivatedCube == 2) || adjacentActivatedCube == 3 {
				newTiles[pos] = struct{}{}
			}
		}

		activatedCube = newTiles
	}
	return len(activatedCube)
}
