package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. :", pt1(ParseCells(lines)))
	fmt.Println("Part 2. :", pt2(ParseCells(lines)))
}

func pt1(lines [][]int) int {
	g := NewGraph(lines)
	node := g.AStar()
	//printGraph(node)

	var score int
	for node.parent != nil {
		score += node.coord.weight
		node = node.parent
	}
	return score
}

func pt2(lines [][]int) int {
	fullLines := make([][]int, len(lines)*5)
	for i := range fullLines {
		fullLines[i] = make([]int, len(lines)*5)
	}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for iLines := 0; iLines < len(lines); iLines++ {
				for jLines := 0; jLines < len(lines); jLines++ {
					value := (lines[iLines][jLines] + i + j) % 9
					if value == 0 {
						value = 9
					}
					fullLines[(len(lines)*i)+iLines][(len(lines)*j)+jLines] = value
				}
			}
		}
	}
	g := NewGraph(fullLines)
	node := g.AStar()
	var score int
	for node.parent != nil {
		score += node.coord.weight
		node = node.parent
	}
	return score
}

func printGraph(node *Node) {
	grid := make([][]string, 5)
	for i := 0; i < 50; i++ {
		grid[i] = make([]string, 5)
		for j := 0; j < 50; j++ {
			grid[i][j] = " "
		}
	}
	for {
		if node == nil {
			break
		}
		grid[node.coord.x][node.coord.y] = strconv.Itoa(node.coord.weight)
		node = node.parent
	}
	for _, v := range grid {
		fmt.Println(v)
	}
}

/////////

type Coord struct {
	x, y   int
	weight int
}

// Graph contains all prerequisites for A*
type Graph struct {
	start       Node
	end         Node
	cells       [][]int
	width       int
	height      int
	openNodes   map[Node]bool
	closedNodes map[Coord]bool
}

// NewGraph build a graph
func NewGraph(cells [][]int) Graph {
	height := len(cells)
	width := len(cells[0])
	startNode := Node{coord: Coord{x: 0, y: 0, weight: cells[0][0]}}
	endNode := Node{coord: Coord{x: height - 1, y: width - 1, weight: cells[height-1][width-1]}}
	return Graph{
		start:       startNode,
		end:         endNode,
		cells:       cells,
		height:      len(cells),
		width:       len(cells[0]),
		openNodes:   map[Node]bool{startNode: true},
		closedNodes: make(map[Coord]bool),
	}
}

func ParseCells(lines []string) [][]int {
	var cells [][]int
	for _, v := range lines {
		lineInt := helpers.ToIntSlice(strings.Split(v, ""))
		cells = append(cells, lineInt)
	}
	return cells
}

func (g Graph) lowestNodeFromOpenNodes() *Node {
	var node *Node
	for n := range g.openNodes {
		if node == nil {
			node = ptrNode(n)
			continue
		}
		if n.F < node.F {
			node = ptrNode(n)
		}
		if n.F == node.F && n.H < node.H { // if F cost is the same, select the H cost with the lowest value
			node = ptrNode(n)
		}
	}
	return node
}

func (g Graph) nodeInOpenList(c Coord) *Node {
	for n := range g.openNodes {
		if n.coord == c {
			return &n
		}
	}
	return nil
}

// Node A* node
type Node struct {
	G      int // G cost (score of the node)
	H      int // H cost (distance from the end node) or heuristic
	F      int // F cost (sum of G and H)
	coord  Coord
	parent *Node
}

func (g Graph) newNode(coord Coord) Node {
	gCost := g.cells[coord.x][coord.y]
	hCost := manhattanDistance(coord, g.end.coord)
	return Node{
		G:     gCost,
		H:     hCost,
		F:     gCost + hCost,
		coord: coord,
	}
}

func manhattanDistance(a, b Coord) int {
	return int(math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)))
}

func ptrNode(n Node) *Node {
	return &n
}

//// GAME
func (g *Graph) AStar() *Node {
	for {
		current := g.lowestNodeFromOpenNodes()
		if current == nil {
			fmt.Println("NO NODES")
			return nil
		}
		delete(g.openNodes, *current)
		g.closedNodes[*&current.coord] = true

		if current.coord == g.end.coord {
			return current
		}

		for _, n := range g.getdAdjacentNotDiag(*current) {
			if g.closedNodes[n.coord] {
				continue
			}
			if nOpenList := g.nodeInOpenList(n.coord); nOpenList != nil {
				if n.F < nOpenList.F {
					nOpenList.F = n.F
					nOpenList.G = n.G
					nOpenList.H = n.H
					nOpenList.parent = current
				}
				continue
			}
			g.openNodes[n] = true
		}
	}
}

///////
func (g *Graph) InGrid(i, j int) bool {
	return i >= 0 && i < g.width && j >= 0 && j < g.height
}

func (g Graph) getdAdjacentNotDiag(n Node) []Node {
	if !g.InGrid(n.coord.x, n.coord.y) {
		return nil
	}

	i, j := n.coord.x, n.coord.y
	var neighbors []Node
	//up
	if i > 0 {
		node := g.newNode(Coord{x: i - 1, y: j, weight: g.cells[i-1][j]})
		node.G += n.G
		node.F = node.G + node.H
		node.parent = &n
		neighbors = append(neighbors, node)
	}
	//down
	if (i + 1) < g.height {
		node := g.newNode(Coord{x: i + 1, y: j, weight: g.cells[i+1][j]})
		node.G += n.G
		node.F = node.G + node.H
		node.parent = &n
		neighbors = append(neighbors, node)
	}
	//left
	if j > 0 {
		node := g.newNode(Coord{x: i, y: j - 1, weight: g.cells[i][j-1]})
		node.G += n.G
		node.F = node.G + node.H
		node.parent = &n
		neighbors = append(neighbors, node)
	}
	//right
	if (j + 1) < g.width {
		node := g.newNode(Coord{x: i, y: j + 1, weight: g.cells[i][j+1]})
		node.G += n.G
		node.F = node.G + node.H
		node.parent = &n
		neighbors = append(neighbors, node)
	}
	return neighbors
}
