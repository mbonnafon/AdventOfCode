package main

import (
	"fmt"
	"strconv"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"N", "E", "S", "W"}[d]
}

type position struct {
	x int
	y int
}

type ship struct {
	position position
	waypoint position
	facing   Direction
}

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. the Manhattan distance between sum of the absolute values of its east/west position and its north/south position is:", pt1(lines))
	fmt.Println("Part 2. the Manhattan distance between sum of the absolute values of its east/west position and its north/south position is:", pt2(lines))
}

func pt1(lines []string) int {
	ship := ship{
		position: position{},
		facing:   East,
	}
	for _, p := range lines {
		ship.move(p)
	}
	return helpers.AbsInt(ship.position.x) + helpers.AbsInt(ship.position.y)
}

func pt2(lines []string) int {
	ship := ship{
		position: position{},
		waypoint: position{
			x: 10,
			y: 1,
		},
		facing: East,
	}
	for _, p := range lines {
		ship.movePt2(p)
	}
	return helpers.AbsInt(ship.position.x) + helpers.AbsInt(ship.position.y)
}

func (s *ship) move(pos string) {
	positioner := string(pos[0])
	distance, _ := strconv.Atoi(pos[1:])
	switch positioner {
	case "N":
		s.position.y += distance
	case "S":
		s.position.y -= distance
	case "E":
		s.position.x += distance
	case "W":
		s.position.x -= distance
	case "L":
		s.rotate(positioner, distance)
	case "R":
		s.rotate(positioner, distance)
	case "F":
		s.move(fmt.Sprintf("%s%d", s.facing.String(), distance))
	}
}

func (s *ship) rotate(direction string, angle int) {
	for angle >= 90 {
		switch s.facing {
		case North:
			if direction == "L" {
				s.facing = West
			} else if direction == "R" {
				s.facing = East
			}
		case East:
			if direction == "L" {
				s.facing = North
			} else if direction == "R" {
				s.facing = South
			}
		case South:
			if direction == "L" {
				s.facing = East
			} else if direction == "R" {
				s.facing = West
			}
		case West:
			if direction == "L" {
				s.facing = South
			} else if direction == "R" {
				s.facing = North
			}
		}
		angle = angle - 90
	}
}

func (s *ship) movePt2(pos string) {
	positioner := string(pos[0])
	distance, _ := strconv.Atoi(pos[1:])
	switch positioner {
	case "N":
		s.waypoint.y += distance
	case "S":
		s.waypoint.y -= distance
	case "E":
		s.waypoint.x += distance
	case "W":
		s.waypoint.x -= distance
	case "L":
		s.rotatePt2(positioner, distance)
	case "R":
		s.rotatePt2(positioner, distance)
	case "F":
		s.position.x += s.waypoint.x * distance
		s.position.y += s.waypoint.y * distance

	}
}

// Just set origin as ship position and use the basic 90° carthesian rotation
// https://math.stackexchange.com/questions/1330161/how-to-rotate-points-through-90-degree
func (s *ship) rotatePt2(direction string, angle int) {
	for angle >= 90 {
		x := s.waypoint.x
		y := s.waypoint.y
		if direction == "L" {
			s.waypoint.x = (0 - y)
			s.waypoint.y = x

		} else {
			s.waypoint.x = y
			s.waypoint.y = (0 - x)
		}
		angle = angle - 90
	}
}
