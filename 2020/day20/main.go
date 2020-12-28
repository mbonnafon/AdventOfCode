package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

type Tile struct {
	number  int
	content [][]string
	top     string
	right   string
	left    string
	bottom  string
}

// SetEdges will set edges values from content value
func (t *Tile) SetEdges() {
	var left, right []string
	for _, l := range t.content {
		left = append(left, string(l[0]))
		right = append(right, string(l[len(l)-1]))
	}
	t.top = strings.Join(t.content[0], " ")
	t.right = strings.Join(right, " ")
	t.left = strings.Join(left, " ")
	t.bottom = strings.Join(t.content[len(t.content)-1], " ")
}

// WithoutEdges returns tile content without edges
func (t Tile) WithoutEdges() [][]string {
	content := make([][]string, len(t.content)-2)
	for i := 1; i < len(t.content)-1; i++ {
		if content[i-1] == nil {
			content[i-1] = make([]string, len(t.content)-2)
		}
		for j := 1; j < len(t.content)-1; j++ {
			content[i-1][j-1] = t.content[i][j]
		}
	}
	return content
}

func (t *Tile) Rotate() {
	helpers.RotateString(t.content)
	t.SetEdges()
}

func (t *Tile) FlipHorizontally() {
	helpers.FlipStringHorizontally(t.content)
	t.SetEdges()
}

func (t *Tile) FlipVertically() {
	helpers.FlipStringVertically(t.content)
	t.SetEdges()
}

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. the product of the four corner IDs tiles is:", pt1(lines))
	fmt.Printf("Part 2. %d # are not part of a sea monster\n", pt2(lines))
}

func pt1(lines []string) int {
	var tiles []Tile
	// parse all tiles
	for i := 0; i < len(lines); i = i + 12 {
		tile := parseTile(lines[i : i+12])
		tiles = append(tiles, tile)
	}

	corners := findCorners(tiles)

	count := 1
	for _, v := range corners {
		count *= v
	}
	return count
}

func pt2(lines []string) int {
	var tiles []Tile
	// parse all tiles
	for i := 0; i < len(lines); i = i + 12 {
		tile := parseTile(lines[i : i+12])
		tiles = append(tiles, tile)
	}

	image := buildImage(tiles)
	img := imageWithoutEdges(image)
	return countSeaMonster(img)
}

func parseTile(lines []string) Tile {
	var number int
	fmt.Sscanf(lines[0], "Tile %d:", &number)

	var content [][]string
	for _, l := range lines[1 : len(lines)-1] {
		var line []string
		for _, c := range l {
			line = append(line, string(c))
		}
		content = append(content, line)
	}
	tile := Tile{
		number:  number,
		content: content,
	}
	tile.SetEdges()
	return tile
}

func findCorners(tiles []Tile) []int {
	adjTiles := findAdjacentTiles(tiles)

	countMatchingEdges := make(map[int]int)
	for _, tiles := range adjTiles {
		for _, t := range tiles {
			countMatchingEdges[t]++
		}
	}

	// consider tile as a corner if exaclty 4 edges are matching (as we reverse 2*2)
	var corners []int
	for k, v := range countMatchingEdges {
		if v == 4 {
			corners = append(corners, k)
		}
	}

	return corners
}

func findAdjacentTiles(tiles []Tile) map[string][]int {
	adjacentTiles := make(map[string][]int)
	for _, t := range tiles {
		adjacentTiles[t.top] = append(adjacentTiles[t.top], t.number)
		adjacentTiles[Reverse(t.top)] = append(adjacentTiles[Reverse(t.top)], t.number)

		adjacentTiles[t.right] = append(adjacentTiles[t.right], t.number)
		adjacentTiles[Reverse(t.right)] = append(adjacentTiles[Reverse(t.right)], t.number)

		adjacentTiles[t.left] = append(adjacentTiles[t.left], t.number)
		adjacentTiles[Reverse(t.left)] = append(adjacentTiles[Reverse(t.left)], t.number)

		adjacentTiles[t.bottom] = append(adjacentTiles[t.bottom], t.number)
		adjacentTiles[Reverse(t.bottom)] = append(adjacentTiles[Reverse(t.bottom)], t.number)
	}

	for k, v := range adjacentTiles {
		if len(v) < 2 {
			delete(adjacentTiles, k)
		}
	}

	return adjacentTiles
}

func buildImage(tiles []Tile) [][]Tile {
	mTiles := make(map[int]*Tile)
	for _, t := range tiles {
		mTiles[t.number] = tPtr(t)
	}
	adjTiles := findAdjacentTiles(tiles)
	corners := findCorners(tiles)

	imageDimension := int(math.Sqrt(float64(len(tiles))))
	image := make([][]Tile, imageDimension)
	for i := 0; i < imageDimension; i++ {
		if len(image[i]) == 0 {
			image[i] = make([]Tile, imageDimension)
		}
		for j := 0; j < imageDimension; j++ {
			// add first tile
			if i == 0 && j == 0 {
				corner := alignFirstTile(mTiles[corners[0]], adjTiles)
				t := mTiles[corner]
				adjTiles[t.right], adjTiles[t.bottom] = removeTile(t.number, adjTiles[t.right]), removeTile(t.number, adjTiles[t.bottom])
				image[0][0] = *mTiles[corner]
				continue
			}
			// if first row, find adjacent tile from previous one
			if i == 0 {
				tID := adjTiles[image[0][j-1].right][0]
				t := mTiles[tID]
				alignToTargetLeft(t, image[0][j-1].right)
				adjTiles[t.right], adjTiles[t.bottom] = removeTile(t.number, adjTiles[t.right]), removeTile(t.number, adjTiles[t.bottom])
				image[0][j] = *t
				continue
			}

			// if others row, find adjacent tile from upper one
			tID := adjTiles[image[i-1][j].bottom][0]
			t := mTiles[tID]
			alignToTargetTop(t, image[i-1][j].bottom)
			adjTiles[t.right], adjTiles[t.bottom] = removeTile(t.number, adjTiles[t.right]), removeTile(t.number, adjTiles[t.bottom])
			image[i][j] = *t
		}
	}
	return image
}

func removeTile(t int, s []int) []int {
	var clean []int
	for _, v := range s {
		if v != t {
			clean = append(clean, v)
		}
	}
	return clean
}

func alignFirstTile(t *Tile, adj map[string][]int) int {
	for i := 0; i < 4; i++ {
		if len(adj[t.right]) == 2 && len(adj[t.bottom]) == 2 {
			return t.number
		}

		t.FlipVertically()
		if len(adj[t.right]) == 2 && len(adj[t.bottom]) == 2 {
			return t.number
		}
		t.FlipVertically()

		t.FlipHorizontally()
		if len(adj[t.right]) == 2 && len(adj[t.bottom]) == 2 {
			return t.number
		}
		t.FlipHorizontally()

		t.Rotate()
	}
	return 0
}

func alignToTargetTop(t *Tile, s string) {
	for i := 0; i < 4; i++ {
		if s == t.top {
			return
		}

		t.FlipVertically()
		if s == t.top {
			return
		}
		t.FlipVertically()

		t.FlipHorizontally()
		if s == t.top {
			return
		}
		t.FlipHorizontally()

		t.Rotate()
	}

}

func alignToTargetLeft(t *Tile, s string) {
	for i := 0; i < 4; i++ {
		if s == t.left {
			return
		}

		t.FlipVertically()
		if s == t.left {
			return
		}
		t.FlipVertically()

		t.FlipHorizontally()
		if s == t.left {
			return
		}
		t.FlipHorizontally()

		t.Rotate()
	}
}

func imageWithoutEdges(tiles [][]Tile) []string {
	edgesToRemove := 2
	imgSize := len(tiles) * (len(tiles[0][0].content) - edgesToRemove)
	img := make([]string, imgSize)
	for i := 0; i < len(tiles); i++ {
		for j := 0; j < len(tiles); j++ {
			t := tiles[i][j].WithoutEdges()
			offset := (i * len(t))
			for k := 0; k < len(t); k++ {
				if img[k+offset] != "" {
					img[k+offset] += " "
				}
				img[k+offset] += strings.Join(t[k], " ")
			}
		}
	}
	return img
}

// Based on https://github.com/Peter554/adventofcode/blob/master/2020/day20/main.go#L262
func countSeaMonster(img []string) int {
	var roughCount int
	var re = regexp.MustCompile(fmt.Sprintf(`#[#\.\n]{%[1]d}#[#\.]{4}##[#\.]{4}##[#\.]{4}###[#\.\n]{%[1]d}#[#\.]{2}#[#\.]{2}#[#\.]{2}#[#\.]{2}#[#\.]{2}#`, len(img)-18))
	seaMonsterVal := 15

	image := make([][]string, len(img))
	for i, l := range img {
		roughCount += strings.Count(l, "#")
		image[i] = strings.Split(l, " ")
	}

	var flatImg string
	for i := 0; i < 4; i++ {
		flatImg = matrixToFlat(image)
		if nb := len(re.FindStringIndex(flatImg)); nb > 0 {
			return roughCount - (test(flatImg, re) * seaMonsterVal)
		}

		helpers.FlipStringVertically(image)
		flatImg = matrixToFlat(image)
		if nb := len(re.FindStringIndex(flatImg)); nb > 0 {
			return roughCount - (test(flatImg, re) * seaMonsterVal)
		}
		helpers.FlipStringVertically(image)

		helpers.FlipStringHorizontally(image)
		flatImg = matrixToFlat(image)
		if nb := len(re.FindStringIndex(flatImg)); nb > 0 {
			return roughCount - (test(flatImg, re) * seaMonsterVal)
		}
		helpers.FlipStringHorizontally(image)

		helpers.RotateString(image)
	}
	fmt.Println(re.FindStringIndex(flatImg))
	return roughCount
}

func matrixToFlat(m [][]string) string {
	var flat string
	for _, l := range m {
		flat += strings.Join(l, "") + "\n"
	}
	return flat
}

func tPtr(t Tile) *Tile {
	return &t
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func test(s string, re *regexp.Regexp) int {
	var count int
	for {
		match := re.FindStringIndex(s)
		if match == nil {
			break
		}
		count++
		s = s[match[0]+1:]
	}
	return count
}
