package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readFile(fileName string) ([]int, error) {
	var listOfElements []int
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lineToInt, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		listOfElements = append(listOfElements, lineToInt)
	}
	return listOfElements, nil
}

func fuelCalcul(mass int) int {
	neededFuel := (mass/3 - 2)
	if neededFuel > 0 {
		return neededFuel + fuelCalcul(neededFuel)
	}
	return 0
}

func main() {
	listOfElements, _ := readFile("./input.txt")
	var fuel int
	for _, mass := range listOfElements {
		fuel = fuel + fuelCalcul(mass)
	}
	fmt.Println(fuel)
}
