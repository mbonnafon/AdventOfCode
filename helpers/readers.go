package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func IntLines(fileName string) ([]int, error) {
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

func StringLines(fileName string) ([]string, error) {
	var listOfElements []string
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		listOfElements = append(listOfElements, scanner.Text())
	}
	return listOfElements, nil
}

func StringGroupOfLines(fileName string) ([]string, error) {
	var listOfElements []string
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var buff, l string
	for scanner.Scan() {
		l = scanner.Text()
		if l == "" {
			listOfElements = append(listOfElements, buff)
			buff = ""
		} else {
			if buff == "" {
				buff = l
			} else {
				buff = fmt.Sprintf("%s %s", buff, l)
			}
		}
	}
	if buff != "" {
		listOfElements = append(listOfElements, buff)
	}
	return listOfElements, nil
}
