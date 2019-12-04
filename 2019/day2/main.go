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

func main() {
	listInputInt := readFileToIntList("./input2.txt")
	fmt.Println(listInputInt)
}
