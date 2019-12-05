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
	fmt.Println(len(listInputInt))
	for i := 0; i <= len(listCopy); i += 4 {
		fmt.Println(i, listCopy[i])
		index1 := listCopy[i+1]
		index2 := listCopy[i+2]
		index3 := listCopy[i+3]

		switch listCopy[i] {
		case 1:
			listCopy[index3] = listCopy[index1] + listCopy[index2]
		case 2:
			listCopy[index3] = listCopy[index1] * listCopy[index2]
		case 99:
			return listCopy[0]
		case 9999:
			fmt.Println("here")
		}
	}
	return 0
}


func main() {
	listInputInt := readFileToIntList("./input.txt")
	fmt.Println(processList(listInputInt))
	
}
