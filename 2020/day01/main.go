package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


func main() {
	//Puzzle input
	pwd, _ := os.Getwd()
	f, _ := os.Open(pwd + "/2020/day01/input.txt")

	scanner := bufio.NewScanner(f)
	var numbers []int
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, i)
	}
	//Puzzle1
	var resultPuzzle1 int

	for _, num1 := range numbers {
		for _, num2 := range numbers {
			if num1 + num2 == 2020 {
				resultPuzzle1 = num1 * num2
				break
			}
		}
	}

	fmt.Println("Result puzzle 1 =", resultPuzzle1)

	//Puzzle1
	var resultPuzzle2 int

	for _, num1 := range numbers {
		for _, num2 := range numbers {
			for _, num3 := range numbers {
				if num1 + num2 + num3 == 2020 {
					resultPuzzle2 = num1 * num2 * num3
					break
				}
			}
		}
	}

	fmt.Println("Result puzzle 2 =", resultPuzzle2)

}
