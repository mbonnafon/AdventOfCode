package main

import (
	"fmt"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. encryption key the is handshake trying to establish is:", pt1(3418282, 8719412))
}

func pt1(cardPublicKey, doorPublicKey int) int {
	//var count int //, cardLoopSize, remainder int
	cardLoopSize := loopSize(cardPublicKey)
	doorLoopSize := loopSize(doorPublicKey)
	encryptionCard := encryption(doorPublicKey, cardLoopSize)
	encryptionDoor := encryption(cardPublicKey, doorLoopSize)
	return encryptionDoor
}

func loopSize(target int) int {
	var loopSize int
	value := 1
	subjectNumber := 7
	for value != target {
		value *= subjectNumber
		value %= 20201227
		loopSize++
	}
	return loopSize
}

func encryption(subjectNumber, loopSize int) int {
	value := 1
	for i := 0; i < loopSize; i++ {
		value *= subjectNumber
		value %= 20201227
	}
	return value
}
