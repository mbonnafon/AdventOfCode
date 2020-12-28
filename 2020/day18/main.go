package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

// http://www.cs.nthu.edu.tw/~wkhon/ds/ds10/tutorial/tutorial2.pdf
func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. the sum of the resulting values is:", pt1(lines))
	fmt.Println("Part 2.  the sum of the resulting values with different precedence levels is:", pt2(lines))
}

// Use the Shunting Yard Algorithm
func pt1(lines []string) int {
	var count int
	operandPriority := map[string]int{
		"+": 1,
		"*": 1,
	}
	for _, l := range lines {
		expression := sanitizeExpression(l)
		exp := parse(operandPriority, expression)
		count += calculate(exp)
	}
	return count
}

func pt2(lines []string) int {
	var count int
	operandPriority := map[string]int{
		"+": 2,
		"*": 1,
	}
	for _, l := range lines {
		expression := sanitizeExpression(l)
		exp := parse(operandPriority, expression)
		count += calculate(exp)
	}
	return count
}

func sanitizeExpression(line string) []string {
	line = strings.ReplaceAll(line, "(", "( ")
	line = strings.ReplaceAll(line, ")", " )")
	return strings.Split(line, " ")
}

// Transform Infix to Postfix
func parse(opPrio map[string]int, expression []string) []string {
	var operatorStack helpers.Stack
	var outputQueue []string
	for _, tok := range expression {
		if helpers.IsInt(tok) {
			outputQueue = append(outputQueue, tok)
		} else {
			switch tok {
			case "(":
				operatorStack.Push("(")
			case ")":
				for {
					op := operatorStack.Pop()
					if op == "(" {
						break
					}
					outputQueue = append(outputQueue, op)
				}
			default:
				if !operatorStack.IsEmpty() {
					op := operatorStack.Pop()
					if op == "(" {
						operatorStack.Push("(")
					} else if opPrio[tok] > opPrio[op] {
						operatorStack.Push(op)
					} else {
						outputQueue = append(outputQueue, op)
						for {
							if operatorStack.IsEmpty() {
								break
							}
							op := operatorStack.Pop()
							if opPrio[tok] > opPrio[op] {
								operatorStack.Push(op)
								break
							}
							outputQueue = append(outputQueue, op)
						}
					}
				}
				operatorStack.Push(tok)
			}
		}
	}
	for !operatorStack.IsEmpty() {
		outputQueue = append(outputQueue, operatorStack.Pop())
	}
	return outputQueue
}

// From Postfix to Answer
func calculate(expression []string) int {
	var stack helpers.IntStack
	for _, tok := range expression {
		i, err := strconv.Atoi(tok)
		if err != nil {
			stack.Push(evaluateOperator(tok, stack.Pop(), stack.Pop()))
		} else {
			stack.Push(i)
		}
	}
	return stack.Pop()
}

func evaluateOperator(oper string, a, b int) int {
	switch oper {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		return 0
	}
}
