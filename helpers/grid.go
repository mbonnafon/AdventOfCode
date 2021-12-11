package helpers

import (
	"strings"
)

type Grid struct {
	Cells  [][]int
	Width  int
	Height int
}

type Coord struct {
	X, Y int
}

func (g *Grid) InGrid(i, j int) bool {
	return i >= 0 && i < g.Width && j >= 0 && j < g.Height
}

func (g *Grid) GetNeigh(i, j int) []Coord {
	if !g.InGrid(i, j) {
		return nil
	}

	neighbors := make([]Coord, 0)
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dy == 0 && dx == 0 {
				continue
			}

			p := Coord{i + dx, j + dy}
			if g.InGrid(p.X, p.Y) {
				neighbors = append(neighbors, p)
			}
		}
	}

	return neighbors
}

func (g Grid) IsSmallestComparedToNeigh(i, j int) bool {
	ref := g.Cells[i][j]
	//up
	if (i > 0) && g.Cells[i-1][j] <= ref {
		return false
	}
	//down
	if ((i + 1) < g.Height) && g.Cells[i+1][j] <= ref {
		return false
	}
	//left
	if (j > 0) && g.Cells[i][j-1] <= ref {
		return false
	}
	//right
	if ((j + 1) < g.Width) && g.Cells[i][j+1] <= ref {
		return false
	}
	return true
}

func ParseGrid(lines []string) Grid {
	var cells [][]int
	for _, v := range lines {
		lineInt := ToIntSlice(strings.Split(v, ""))
		cells = append(cells, lineInt)
	}
	return Grid{
		Cells:  cells,
		Height: len(cells),
		Width:  len(cells[0]),
	}
}
