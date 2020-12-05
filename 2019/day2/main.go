package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFileToIntList(fileName string) []int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	line, err := ioutil.ReadAll(file)
	lineSplitted := strings.Split(string(line), ",")
	var intList []int
	for _, v := range lineSplitted {
		vToint, _ := strconv.Atoi(v)
		intList = append(intList, vToint)
	}
	return intList
}

//Part1
func processList(listInputInt []int) int {
	listCopy := make([]int, len(listInputInt))
	copy(listCopy, listInputInt)
	for i := 0; i <= len(listCopy); i += 4 {
		index1 := listCopy[i+1]
		index2 := listCopy[i+2]
		index3 := listCopy[i+3]

		if listCopy[i] == 1 {
			listCopy[index3] = listCopy[index1] + listCopy[index2]
		} else if listCopy[i] == 2 {
			listCopy[index3] = listCopy[index1] * listCopy[index2]
		} else if listCopy[i] == 99 {
			return listCopy[0]
		}
	}
	return 0
}

//Part2
func bruteforceList(listInputInt []int) int {
	var listCopy []int
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			listCopy = make([]int, len(listInputInt))
			copy(listCopy, listInputInt)
			listCopy[1] = i
			listCopy[2] = j
			if processList(listCopy) == 19690720 {
				return 100*i + j
			}
		}
	}
	return 0
}

func pt1(lines []int) int {
	return 0
}
func pt2(lines []int) int {
	return 0
}

func main() {
	// TODO
	// lines, _ := helpers.IntLines("./input.txt")
	// fmt.Println("Part 1. value left at position 0 after the program halts is:", pt1(lines))
	// fmt.Println("Part 2. :", pt2(lines))
	listInputInt := readFileToIntList("./input2.txt")
	fmt.Println(processList(listInputInt))
	fmt.Println(bruteforceList(listInputInt))
}
