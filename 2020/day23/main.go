package main

import (
	"fmt"
	"strconv"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

type node struct {
	val  int
	next *node
}

func main() {
	fmt.Println("Part 1. labels on the cups after cup 1 are:", pt1("326519478"))
	fmt.Println("Part 2. product for the labels of the two cups that will end up immediately clockwise of cup 1:", pt2("326519478"))
}

// Use circular linked list
func pt1(input string) string {
	cups := setGame(input, false)
	head, labelPtr := createCircularLinkedList(cups)
	playMoves(100, head, helpers.Min(cups), helpers.MaxS(cups), labelPtr)

	//Starting after the cup labeled 1
	var finalOrder string
	pos := labelPtr[1]
	for {
		pos = pos.next
		if pos == labelPtr[1] {
			break
		}
		finalOrder = fmt.Sprintf("%s%d", finalOrder, pos.val)
	}
	return finalOrder
}

func pt2(input string) int {
	cups := setGame(input, true)
	head, labelPtr := createCircularLinkedList(cups)
	playMoves(10000000, head, helpers.Min(cups), helpers.MaxS(cups), labelPtr)
	return labelPtr[1].next.val * labelPtr[1].next.next.val
}

func setGame(input string, part2 bool) []int {
	var circle []int
	for _, s := range input {
		label, _ := strconv.Atoi(string(s))
		circle = append(circle, label)
	}
	if part2 {
		for i := helpers.MaxS(circle) + 1; i <= 1000000; i++ {
			circle = append(circle, i)
		}
	}
	return circle
}

func createCircularLinkedList(cups []int) (*node, map[int]*node) {
	var head, tail, p *node
	labelPtr := make(map[int]*node)
	for i, label := range cups {
		// nodes
		n := &node{val: label}
		if i == 0 {
			head = n
			p = head
		}
		if i == len(cups)-1 {
			tail = n
		}
		p.next = n
		p = n
		labelPtr[label] = n
	}
	tail.next = head
	labelPtr[tail.val] = tail
	labelPtr[head.val] = head

	return head, labelPtr
}

func playMoves(moves int, currentCup *node, min, max int, labelPtr map[int]*node) {
	for i := 1; i <= moves; i++ {
		pickUpList := listPickUp(currentCup)
		last := lastPickUp(currentCup)
		dest := findDest(currentCup.val, min, max, pickUpList, labelPtr)

		// uncomment to display game move
		// fmt.Printf("-- move %d --\n", i)
		// fmt.Printf("Cups: (%d) ", currentCup.val)
		// displayList(currentCup)
		// fmt.Println("Pick up: ", pickUpList, last.val)
		// fmt.Println("Destination ", dest.val)
		// fmt.Println("")

		cn := currentCup.next
		ln := last.next
		dn := dest.next
		currentCup.next = ln
		last.next = dn
		dest.next = cn
		currentCup = currentCup.next
	}
}

func listPickUp(currentCup *node) map[int]bool {
	ptr := currentCup.next
	m := make(map[int]bool)
	for i := 0; i < 3; i++ {
		m[ptr.val] = true
		ptr = ptr.next
	}
	return m
}

func findDest(currentCup, min, max int, pickUpList map[int]bool, labelPtr map[int]*node) *node {
	currentCup--
	for currentCup >= min {
		if !pickUpList[currentCup] {
			return labelPtr[currentCup]
		}
		currentCup--
	}
	currentCup = max
	for currentCup >= min {
		if !pickUpList[currentCup] {
			return labelPtr[currentCup]
		}
		currentCup--
	}
	return nil
}

func lastPickUp(currentCup *node) *node {
	ptr := currentCup.next
	for i := 1; i < 3; i++ {
		ptr = ptr.next
	}
	return ptr
}

func displayList(head *node) {
	var count int
	n := head
	for {
		fmt.Printf("%d ", n.val)
		n = n.next
		if n == head {
			break
		}
		count++
	}
	fmt.Println("")
}
