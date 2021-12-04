package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

var space = regexp.MustCompile(`\s+`)

type Game struct {
	draws  []int
	boards []*Board
}

type Board struct {
	grid   [][]*Cell
	winner bool
}

type Cell struct {
	number int
	drawn  bool
}

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. :", pt1(lines))
	fmt.Println("Part 2. :", pt2(lines))
}

func pt1(lines []string) int {
	g := parseGame(lines)
	return g.play(true)
}

func pt2(lines []string) int {
	g := parseGame(lines)
	return g.play(false)
}

func parseGame(lines []string) Game {
	g := Game{}
	g.draws = func(line string) []int {
		var s []int
		for _, v := range strings.Split(line, ",") {
			n, _ := strconv.Atoi(string(v))
			s = append(s, n)
		}
		return s
	}(lines[0])

	for i := 1; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		b := Board{}
		b.grid = make([][]*Cell, 5)
		for j := 0; j < 5; j++ {
			b.grid[j] = func(line string) []*Cell {
				var s []*Cell
				line = space.ReplaceAllString(line, " ")
				for _, v := range strings.Split(line, " ") {
					n, err := strconv.Atoi(string(v))
					if err != nil {
						continue
					}
					s = append(s, &Cell{number: n, drawn: false})
				}
				return s
			}(lines[i+j])
		}
		g.boards = append(g.boards, &b)
		i = i + 5
	}

	return g
}

func (g Game) play(first bool) int {
	var winnerScore int
	for _, numb := range g.draws {
		for _, b := range g.boards {
			if b.winner {
				continue
			}
			b.updateBoard(numb)
			if first && b.isWinner() {
				winnerScore = numb * b.countPoints()
				return winnerScore
			}
			if !first && b.isWinner() {
				winnerScore = numb * b.countPoints()
				b.winner = true
			}
		}
	}
	return winnerScore
}

func (b Board) updateBoard(number int) {
	for _, x := range b.grid {
		for _, y := range x {
			if y.number == number {
				y.drawn = true
			}
		}
	}
}

func (b Board) isWinner() bool {
	for i := 0; i < 5; i++ {
		// rows
		if b.grid[i][0].drawn && b.grid[i][1].drawn && b.grid[i][2].drawn && b.grid[i][3].drawn && b.grid[i][4].drawn {
			return true
		}
		// columns
		if b.grid[0][i].drawn && b.grid[1][i].drawn && b.grid[2][i].drawn && b.grid[3][i].drawn && b.grid[4][i].drawn {
			return true
		}
	}
	return false
}

func (b Board) countPoints() int {
	var points int
	for _, x := range b.grid {
		for _, y := range x {
			if y.drawn == false {
				points += y.number
			}
		}
	}
	return points
}
