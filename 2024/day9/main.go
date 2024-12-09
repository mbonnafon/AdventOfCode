package main

import (
	"fmt"
	"strings"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. :", pt1(lines))
	fmt.Println("Part 2. :", pt2(lines))
}

type Filesystem []Block

func NewFilesystem(seq []int) Filesystem {
	var filesystem Filesystem
	for i, nbObj := range seq {
		content := make([]int, 0, nbObj)
		block := Block{content: content}
		if i%2 == 0 { // file
			id := i / 2
			for range nbObj {
				block.content = append(block.content, id)
			}
		}
		filesystem = append(filesystem, block)
	}
	return filesystem
}

func (fs Filesystem) Print() {
	for _, b := range fs {
		for i := range cap(b.content) {
			if i < len(b.content) {
				fmt.Print(b.content[i])
			} else {
				fmt.Print(".")
			}
		}
	}
	fmt.Println("")
}

func (fs Filesystem) FindFreeSpaceFor(size int) *int {
	for idx, f := range fs {
		if f.RemainingSpace() >= size {
			return &idx
		}
	}
	return nil
}

func (fs Filesystem) Checksum() int {
	var checksum int
	var counter int
	for _, b := range fs {
		for i := range cap(b.content) {
			if i < len(b.content) {
				checksum += b.content[i] * counter
			}
			counter++
		}
	}
	return checksum
}

type Block struct {
	content []int
}

func (b Block) Size() int {
	return len(b.content)
}

func (b Block) RemainingSpace() int {
	return cap(b.content) - len(b.content)
}

func (b Block) Full() bool {
	return len(b.content) == cap(b.content)
}

func (b Block) Empty() bool {
	return len(b.content) == 0
}

func (b *Block) PopLast() int {
	lastIdx := len(b.content) - 1
	last := b.content[lastIdx]
	b.content = b.content[:lastIdx]
	return last
}

func (b *Block) PopBlock() []int {
	block := b.content
	b.content = b.content[:0]
	return block
}

func (b *Block) Append(i int) {
	b.content = append(b.content, i)
}

func (b *Block) InsertBlock(block []int) {
	b.content = append(b.content, block...)
}

func pt1(lines []string) int {
	seq := helpers.ToIntSlice(strings.Split(lines[0], ""))
	filesystem := NewFilesystem(seq)

	i, j := 0, len(filesystem)-1
	for i != j {
		if filesystem[i].Full() {
			i++
			continue
		}
		if filesystem[j].Empty() {
			j--
			continue
		}
		last := filesystem[j].PopLast()
		filesystem[i].Append(last)
	}

	return filesystem.Checksum()
}

func pt2(lines []string) int {
	seq := helpers.ToIntSlice(strings.Split(lines[0], ""))
	filesystem := NewFilesystem(seq)

	//filesystem.Print()
	for i := len(filesystem) - 1; i > 1; i-- {
		if filesystem[i].Full() {
			if pos := filesystem.FindFreeSpaceFor(filesystem[i].Size()); pos != nil && *pos < i { // consider that we only want to compact to the left
				block := filesystem[i].PopBlock()
				filesystem[*pos].InsertBlock(block)
			}
			continue
		}
	}
	//filesystem.Print()

	return filesystem.Checksum()
}
