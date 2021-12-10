package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Puzzle input
	input := GetPuzzleInput()

	//Puzzle 1
	var sum int
	for _, expr := range input {
		result := EvaluateExpression(expr, GetResultSamePrecedence)
		sum += result
	}
	fmt.Println("Puzzle 1 =", sum)

	//Puzzle 2
	sum = 0
	for _, expr := range input {
		result := EvaluateExpression(expr, GetResultAdditionBeforeMultiplication)
		sum += result
	}
	fmt.Println("Puzzle 2 =", sum)
}

func EvaluateExpression(expr string, getResult func(expression string) int) int {
	for {
		if strings.Count(expr, "(") == 0 {
			return getResult(expr)
		}
		var newExpression string
		for _, item := range expr {
			if string(item) == "(" {
				newExpression = ""
			} else if string(item) == ")" {
				var result int
				result = getResult(newExpression)
				expr = strings.Replace(expr, "(" + newExpression + ")", strconv.Itoa(result), 1)
				break
			} else {
				newExpression += string(item)
			}
		}
	}
}

func GetResultAdditionBeforeMultiplication(expression string) int {
	var result int
	for strings.Contains(expression, "+") {
		expressionList := strings.Split(expression, " ")
		for i, item := range expressionList {
			if item == "+" {
				num1, _ := strconv.Atoi(expressionList[i-1])
				num2, _ := strconv.Atoi(expressionList[i+1])
				sum := num1 + num2
				oldString := expressionList[i-1] + " + " + expressionList[i+1]
				expression = strings.Replace(expression, oldString, strconv.Itoa(sum), 1)
				break
			}
		}
	}
	if strings.Contains(expression, "*") {
		for i, item := range strings.Split(expression, " ") {
			if item == "*" {
				continue
			}
			num, _ := strconv.Atoi(item)
			if i == 0 {
				result = num
			} else {
				result *= num
			}
		}
	} else {
		result, _ = strconv.Atoi(expression)
	}
	return result
}

func GetResultSamePrecedence(expression string) int {
	var result int
	var operation string
	for i, item := range strings.Split(expression, " ") {
		switch item {
		case "+", "*":
			operation = item
		default:
			num, _ := strconv.Atoi(item)
			if i == 0 {
				result = num
			} else if operation == "+" {
				result += num
			} else if operation == "*" {
				result *= num
			}
		}
	}
	return result
}

func GetPuzzleInput() []string {
	pwd, _ := os.Getwd()
	f, _ := os.Open(pwd + "/2020/day18/input.txt")

	scanner := bufio.NewScanner(f)
	var puzzleInput []string
	for scanner.Scan() {
		line := scanner.Text()
		puzzleInput = append(puzzleInput, line)
	}
	return puzzleInput
}