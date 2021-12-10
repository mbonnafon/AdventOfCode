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
	fmt.Println("Part 1. :", pt1(parseGame(lines)))
	fmt.Println("Part 2. :", pt2(parseGame(lines)))
}

func pt1(g Game) int {
	return g.play(true)
}

func pt2(g Game) int {
	return g.play(false)
}

func parseGame(lines []string) Game {
	g := Game{}
	g.draws = func(line string) (draws []int) {
		for _, v := range strings.Split(line, ",") {
			draws = append(draws, helpers.ToInt(v))
		}
		return
	}(lines[0])

	// discard already processed draws
	for i := 1; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		b := Board{}
		b.grid = make([][]*Cell, 5)
		for j := 0; j < 5; j++ {
			b.grid[j] = func(line string) []*Cell {
				var cells []*Cell
				for _, c := range strings.Split(line, " ") {
					n, err := strconv.Atoi(string(c))
					if err != nil { //hack to discard spaces
						continue
					}
					cells = append(cells, &Cell{number: n, drawn: false})
				}
				return cells
			}(lines[i+j])
		}
		g.boards = append(g.boards, &b)
		i = i + 5
	}

	return g
}

func (g Game) play(firstIsWinner bool) (winnerScore int) {
	for _, draw := range g.draws {
		for _, board := range g.boards {
			if board.winner {
				continue
			}
			board.updateBoard(draw)
			if firstIsWinner && board.isWinner() {
				winnerScore = draw * board.countPoints()
				return winnerScore
			}
			if !firstIsWinner && board.isWinner() {
				winnerScore = draw * board.countPoints()
				board.winner = true
			}
		}
	}
	return
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

func (b Board) countPoints() (points int) {
	for _, x := range b.grid {
		for _, y := range x {
			if y.drawn == false {
				points += y.number
			}
		}
	}
	return
}
