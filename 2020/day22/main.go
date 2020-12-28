package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

type player struct {
	number       int
	deck         Queue
	previousDeck []string
}

func (p *player) logGame() {
	p.previousDeck = append(p.previousDeck, computeHashForList(p.deck))
}

func (p *player) isCyclic() bool {
	hash := computeHashForList(p.deck)
	for _, d := range p.previousDeck {
		if hash == d {
			return true
		}
	}
	return false
}

func (p player) score() int {
	var count int
	deckLen := len(p.deck)
	for i := 0; i < deckLen; i++ {
		count += p.deck.Draw() * (deckLen - i)
	}
	return count
}

type Queue []int

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

func (q *Queue) Draw() int {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}

func (q Queue) String() string {
	var s string
	for _, c := range q {
		s = fmt.Sprintf("%s %d,", s, c)
	}
	return strings.TrimSuffix(s, ",")
}

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. the winning player's score is:", pt1(lines))
	fmt.Println("Part 2. the winning player's score is:", pt2(lines))
}

func pt1(lines []string) int {
	p1, p2 := parseGame(lines)
	return playGame(1, false, p1, p2).score()
}

func pt2(lines []string) int {
	p1, p2 := parseGame(lines)
	return playGame(1, true, p1, p2).score()
}

func parseGame(lines []string) (player, player) {
	var p1Parsed bool
	var p1, p2 player
	p1.number = 1
	p2.number = 2
	for _, l := range lines {
		if l == "Player 2:" {
			p1Parsed = true
		}
		i, err := strconv.Atoi(l)
		if err != nil {
			continue
		}
		if !p1Parsed {
			p1.deck.Push(i)
		} else {
			p2.deck.Push(i)
		}
	}
	return p1, p2
}

func playGame(game int, recursive bool, p1, p2 player) player {
	var round int
	for !p1.deck.IsEmpty() && !p2.deck.IsEmpty() {
		round++
		// check for cyclic round
		if p1.isCyclic() && p2.isCyclic() {
			return p1
		}
		p1.logGame()
		p2.logGame()
		winner, c1, c2 := playRound(game, recursive, &p1, &p2)
		winner.deck.Push(c1)
		winner.deck.Push(c2)
	}
	var winner player
	if p1.deck.IsEmpty() {
		winner = p2
	} else {
		winner = p1
	}
	return winner
}

func playRound(game int, recursive bool, p1, p2 *player) (*player, int, int) {
	// play normal game
	p1Card, p2Card := p1.deck.Draw(), p2.deck.Draw()
	if recursive && (len(p1.deck) >= p1Card) && (len(p2.deck) >= p2Card) {
		p1Copy, p2Copy := *p1, *p2
		p1Copy.deck = make([]int, len(p1.deck))
		p2Copy.deck = make([]int, len(p2.deck))
		copy(p1Copy.deck, p1.deck)
		copy(p2Copy.deck, p2.deck)
		p1Copy.deck, p2Copy.deck = p1Copy.deck[:p1Card], p2Copy.deck[:p2Card]
		winner := playGame(game+1, recursive, p1Copy, p2Copy)
		if winner.number == 1 {
			return p1, p1Card, p2Card
		}
		return p2, p2Card, p1Card
	}
	if p1Card > p2Card {
		return p1, p1Card, p2Card
	}
	return p2, p2Card, p1Card
}

func computeHashForList(list []int) string {
	var s string
	for _, i := range list {
		s = fmt.Sprintf("%s%s", s, strconv.Itoa(i))
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
