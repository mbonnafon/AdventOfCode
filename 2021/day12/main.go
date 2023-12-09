package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

type Node struct {
	prev []*Node
	next []*Node
	key  string
}

func (n Node) isBigCave() bool {
	return unicode.IsUpper(rune(n.key[0]))
}

type List struct {
	head  *Node
	tail  *Node
	elems map[string]*Node
}

func newList(head, tail string) List {
	headNode := &Node{key: "start"}
	tailNode := &Node{key: "end"}
	elems := map[string]*Node{
		head: headNode,
		tail: tailNode,
	}
	return List{
		head:  headNode,
		tail:  tailNode,
		elems: elems,
	}
}

func (l *List) Insert(curr, next string) {
	var currNode *Node
	if l.elems[curr] == nil {
		currNode = &Node{key: curr}
		l.elems[curr] = currNode
	} else {
		currNode = l.elems[curr]
	}

	var nextNode *Node
	if l.elems[next] == nil {
		nextNode = &Node{key: next}
		l.elems[next] = nextNode
	} else {
		nextNode = l.elems[next]
	}

	currNode.next = append(currNode.next, nextNode)
	nextNode.prev = append(nextNode.prev, currNode)
}

func (l List) visualize(curr string) {
	node := l.elems[curr]
	for _, n := range node.next {
		fmt.Printf("%s => %s (%s)\n", curr, n.key, n.isBigCave())
		l.visualize(n.key)
	}
}

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. :", pt1(lines))
	fmt.Println("Part 2. :", pt2(lines))
}

func pt1(lines []string) int {
	return 0
	list := newList("start", "end")
	for _, v := range lines {
		splitted := strings.Split(v, "-")
		list.Insert(splitted[0], splitted[1])
	}

	path := make(map[string]int)
	for k, v := range list.elems {
		if !v.isBigCave() {
			path[k] = 1
		}
	}
	return list.doPath(*list.head, false, path)
}

func pt2(lines []string) int {
	list := newList("start", "end")
	for _, v := range lines {
		splitted := strings.Split(v, "-")
		list.Insert(splitted[0], splitted[1])
	}

	path := make(map[string]int)
	return list.doPath(*list.head, false, path)
}

func (l *List) doPath(current Node, alreadyVisitedTwice bool, visited map[string]int) (count int) {
	newlyVisited := make(map[string]int)
	for k, v := range visited {
		newlyVisited[k] = v
	}

	if !current.isBigCave() && newlyVisited[current.key] > 0 {
		alreadyVisitedTwice = true
	}
	newlyVisited[current.key]++

	nextNodes := make(map[*Node]bool)
	for _, n := range current.next {
		nextNodes[n] = true
	}
	for _, n := range current.prev {
		nextNodes[n] = true

	}

	for n := range nextNodes {
		if n.key == l.head.key {
			continue
		}
		if n.key == l.tail.key {
			count++
			continue
		}
		if !n.isBigCave() && alreadyVisitedTwice && newlyVisited[n.key] > 0 {
			continue
		}
		count += l.doPath(*n, alreadyVisitedTwice, newlyVisited)
	}

	return
}
