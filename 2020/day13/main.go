package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. the ID of the earliest bus you can take to the airport multiplied by the number of minutes you'll need to wait for that bus is:", pt1(lines))
	fmt.Println("Part 2. the earliest timestamp such that all of the listed bus IDs depart at offsets matching their positions in the list is:", pt2(lines))
}

func pt1(lines []string) int {
	departure, _ := strconv.Atoi(lines[0])
	s := strings.Split(lines[1], ",")
	var busID, busSchedule int
	for _, v := range s {
		if v == "x" {
			continue
		}
		currBusID, _ := strconv.Atoi(v)
		busArrival := (departure/currBusID)*currBusID + currBusID
		if (busArrival > departure && busArrival < busSchedule) || busSchedule == 0 {
			busID = currBusID
			busSchedule = busArrival
		}
	}
	return (busSchedule - departure) * busID
}

type bus struct {
	position int
	number   int
}

func pt2(lines []string) int {
	var buses []bus
	s := strings.Split(lines[1], ",")
	for index, value := range s {
		if value == "x" {
			continue
		}
		v, _ := strconv.Atoi(value)
		buses = append(buses, bus{index, v})
	}

	var current int
	increment := buses[0].number
	for _, v := range buses[1:] {
		current = commonModulo(current, increment, buses[0], v)
		increment *= v.number
	}
	return current
}

func commonModulo(current, increment int, b1, b2 bus) int {
	for {
		current += increment
		if (current+(b2.position-b1.position))%b2.number == 0 {
			return current
		}
	}
	return 0
}
