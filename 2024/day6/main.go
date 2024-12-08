package main

import (
	"fmt"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

const Visited = 'X'
const Obstacle = '#'

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. :", pt1(lines))
	fmt.Println("Part 2. :", pt2(lines))
}

type Boundaries struct {
	xMin int
	xMax int
	yMin int
	yMax int
}

type Game struct {
	Map        [][]rune
	boundaries Boundaries
	character  *Pos
	start      Pos
}

func (g Game) offTheMap(pos Pos) bool {
	return (pos.x < g.boundaries.xMin || pos.x > g.boundaries.xMax-1 || pos.y < g.boundaries.yMin || pos.y > g.boundaries.yMax-1)
}

func findStartPos(lines []string) *Pos {
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			if lines[i][j] == '^' {
				return &Pos{x: i, y: j, direction: UP}
			}
		}
	}
	return nil
}

func NewGame(lines []string) *Game {
	character := findStartPos(lines)
	plan := make([][]rune, len(lines))
	for i, line := range lines {
		plan[i] = []rune(line)
	}
	boundaries := Boundaries{xMin: 0, xMax: len(lines[0]), yMin: 0, yMax: len(lines)}
	game := &Game{Map: plan, boundaries: boundaries, character: character, start: *character}
	game.markAsVisited(character)
	return game
}

func (g Game) print() {
	fmt.Println("-------------")
	for _, l := range g.Map {
		fmt.Println(string(l))
	}
	fmt.Println("-------------")
}

func (g *Game) play() {
	//fmt.Println("Start at:", *g.character)
	for {
		//g.print()
		nextMove := g.character.nextMove()
		if g.offTheMap(nextMove) {
			g.markAsVisited(g.character)
			break
		}
		if g.foundObstacle(nextMove) {
			g.character.turn()
			continue
		}
		g.markAsVisited(g.character)
		g.move(nextMove)
	}
}

func (g Game) foundObstacle(pos Pos) bool {
	return (g.Map[pos.x][pos.y] == Obstacle)
}

func (g *Game) countVisited() int {
	var count int
	for i := 0; i < len(g.Map); i++ {
		for j := 0; j < len(g.Map[i]); j++ {
			if g.Map[i][j] == Visited {
				count++
			}
		}
	}
	return count
}

func (g *Game) markAsVisited(pos *Pos) {
	g.Map[pos.x][pos.y] = Visited
}

func (g *Game) markAsObstacle(pos *Pos) {
	g.Map[pos.x][pos.y] = Obstacle
}

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

type Pos struct {
	x         int
	y         int
	direction Direction
}

func (p *Game) move(newPos Pos) {
	p.character.x = newPos.x
	p.character.y = newPos.y
}

func (p Pos) nextMove() Pos {
	switch p.direction {
	case UP:
		p.x--
	case DOWN:
		p.x++
	case LEFT:
		p.y--
	case RIGHT:
		p.y++
	}
	return p
}

func (p *Pos) turn() {
	switch p.direction {
	case UP:
		p.direction = RIGHT
	case DOWN:
		p.direction = LEFT
	case LEFT:
		p.direction = UP
	case RIGHT:
		p.direction = DOWN
	}
}

func pt1(lines []string) int {
	game := NewGame(lines)
	game.play()
	return game.countVisited()
}

func pt2(lines []string) int {
	g := NewGame(lines)
	visited := make(map[Pos]bool)
	for {
		nextMove := g.character.nextMove()
		if g.offTheMap(nextMove) {
			g.markAsVisited(g.character)
			break
		}
		if g.foundObstacle(nextMove) {
			g.character.turn()
			continue
		}
		g.markAsVisited(g.character)
		g.move(nextMove)
		visited[nextMove] = true
	}

	obstructions := make(map[Pos]bool)
	for v := range visited {
		g = NewGame(lines)

		// simulate an obstacle for the next move
		nextMoveForObstacle := v.nextMove()
		if g.offTheMap(nextMoveForObstacle) {
			continue
		}
		g.Map[nextMoveForObstacle.x][nextMoveForObstacle.y] = Obstacle
		seen := make(map[Pos]bool)
		for {
			nextMove := g.character.nextMove()
			if seen[nextMove] {
				obstructions[Pos{nextMoveForObstacle.x, nextMoveForObstacle.y, 0}] = true
				break
			}
			seen[nextMove] = true
			if g.offTheMap(nextMove) {
				g.markAsVisited(g.character)
				break
			}
			if g.foundObstacle(nextMove) {
				g.character.turn()
				continue
			}
			g.markAsVisited(g.character)
			g.move(nextMove)
		}
	}

	for v := range visited {
		g = NewGame(lines)

		// simulate an obstacle for the current move
		if g.offTheMap(v) {
			continue
		}
		g.Map[v.x][v.y] = Obstacle
		seen := make(map[Pos]bool)
		for {
			nextMove := g.character.nextMove()
			if seen[nextMove] {
				obstructions[Pos{v.x, v.y, 0}] = true
				break
			}
			seen[nextMove] = true
			if g.offTheMap(nextMove) {
				g.markAsVisited(g.character)
				break
			}
			if g.foundObstacle(nextMove) {
				g.character.turn()
				continue
			}
			g.markAsVisited(g.character)
			g.move(nextMove)
		}
	}

	return len(obstructions)
}
