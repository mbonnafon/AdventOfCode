package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

type Coords struct {
	x, y int
}

func (c Coords) toString() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}

type Vent struct {
	x1, y1 int
	x2, y2 int
}

func (v Vent) isHorizontalOrVertical() bool {
	if (v.x1 == v.x2) || (v.y1 == v.y2) {
		return true
	}
	return false
}

func (v Vent) isDiagonal() bool {
	if helpers.AbsInt(v.x1-v.x2) == helpers.AbsInt(v.y1-v.y2) {
		return true
	}
	return false
}

func (v Vent) horizontalOrVerticalCoverage() []Coords {
	var coords []Coords
	if v.x1 == v.x2 {
		if v.y1 < v.y2 {
			for i := v.y1; i <= v.y2; i++ {
				coords = append(coords, Coords{x: v.x1, y: i})
			}
		} else {
			for i := v.y1; i >= v.y2; i-- {
				coords = append(coords, Coords{x: v.x1, y: i})
			}
		}
	} else {
		if v.x1 < v.x2 {
			for i := v.x1; i <= v.x2; i++ {
				coords = append(coords, Coords{x: i, y: v.y1})
			}
		} else {
			for i := v.x1; i >= v.x2; i-- {
				coords = append(coords, Coords{x: i, y: v.y1})
			}
		}

	}
	return coords
}

func (v Vent) diagonalCoverage() []Coords {
	var coords []Coords
	var xOperator, yOperator int
	if v.x1 > v.x2 {
		xOperator = -1
	} else {
		xOperator = 1
	}
	if v.y1 > v.y2 {
		yOperator = -1
	} else {
		yOperator = 1
	}

	dist := int(math.Abs(float64(v.x1 - v.x2)))
	for i := 0; i <= dist; i++ {
		x := v.x1 + (xOperator * i)
		y := v.y1 + (yOperator * i)
		coords = append(coords, Coords{x, y})
	}

	return coords
}

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	vents := parseVents(lines)
	fmt.Println("Part 1. :", pt1(vents))
	fmt.Println("Part 2. :", pt2(vents))
}

func pt1(vents []Vent) int {
	m := make(map[string]int)
	for _, v := range vents {
		if v.isHorizontalOrVertical() {
			for _, c := range v.horizontalOrVerticalCoverage() {
				m[c.toString()]++
			}
		}
	}
	return score(m)
}

func pt2(vents []Vent) int {
	m := make(map[string]int)
	for _, v := range vents {
		if v.isHorizontalOrVertical() {
			for _, c := range v.horizontalOrVerticalCoverage() {
				m[c.toString()]++
			}
		}
		if v.isDiagonal() {
			for _, c := range v.diagonalCoverage() {
				m[c.toString()]++
			}
		}
	}
	return score(m)
}

func parseVents(lines []string) []Vent {
	var vents []Vent
	for _, v := range lines {
		sVent := func(s string) []int {
			s = strings.ReplaceAll(s, " -> ", ",")
			coords := strings.Split(s, ",")
			var r []int
			for _, v := range coords {
				a, _ := strconv.Atoi(string(v))
				r = append(r, a)
			}
			return r
		}(v)
		vent := Vent{
			x1: sVent[0],
			y1: sVent[1],
			x2: sVent[2],
			y2: sVent[3],
		}
		vents = append(vents, vent)
	}
	return vents
}

func score(m map[string]int) (counter int) {
	for _, v := range m {
		if v > 1 {
			counter++
		}
	}
	return
}
