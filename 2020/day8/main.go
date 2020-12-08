package main

import (
	"fmt"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

type Command struct {
	operation string
	value     int
	executed  bool
}

func parseCommands(lines []string) []Command {
	var operation string
	var val int
	commands := make([]Command, 0, len(lines))
	for _, l := range lines {
		fmt.Sscanf(l, "%s %d", &operation, &val)
		commands = append(commands, Command{operation: operation, value: val})
	}
	return commands
}

func calcAccVal(commands []Command, allowDup bool) int {
	var accumulator int
	var i int
	for {
		if i >= len(commands) || (commands[i].executed && allowDup) {
			break
		}
		if commands[i].executed && !allowDup {
			return -1
		}

		switch commands[i].operation {
		case "acc":
			accumulator = accumulator + commands[i].value
			commands[i].executed = true
			i++
		case "jmp":
			commands[i].executed = true
			i = i + commands[i].value
			continue
		case "nop":
			i++
			continue
		}
	}
	return accumulator
}

func pt1(lines []string) int {
	commands := parseCommands(lines)
	return calcAccVal(commands, true)
}

func pt2(lines []string) int {
	var c = make([]Command, len(lines))
	commands := parseCommands(lines)
	for i, command := range commands {
		copy(c, commands)
		if command.operation == "jmp" {
			c[i].operation = "nop"
			if acc := calcAccVal(c, false); acc > 0 {
				return acc
			}
			continue
		}
		if command.operation == "nop" {
			copy(c, commands)
			c[i].operation = "jmp"
			if acc := calcAccVal(c, false); acc > 0 {
				return acc
			}
			continue
		}
	}
	return 0
}

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. The accumulator value is:", pt1(lines))
	fmt.Println("Part 2. The accumulator value after the program terminates is:", pt2(lines))
}
