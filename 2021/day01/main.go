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
	f, _ := os.Open(pwd + "/day01/input.txt")

	scanner := bufio.NewScanner(f)
	var puzzleInput []int
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		puzzleInput = append(puzzleInput, num)
	}

	//Puzzle 1
	result := CountMeasurementIsLarger(puzzleInput)
	fmt.Println("Puzzle 1 =", result)

	//Puzzle 2
	sumsOfThree := CreateSumsOfThree(puzzleInput)
	result = CountMeasurementIsLarger(sumsOfThree)
	fmt.Println("Puzzle 2 =", result)
}

func CreateSumsOfThree(numbers []int) []int {
	var newNumbers []int
	end := len(numbers) - 2
	for i, v := range numbers[:end] {
		num := v + numbers[i+1] + numbers[i+2]
		newNumbers = append(newNumbers, num)
	}
	return newNumbers
}

func CountMeasurementIsLarger(numbers []int) int {
	var isLarger int
	for i, v := range numbers {
		if i == 0 {
			continue
		}
		if v > numbers[i-1] {
			isLarger++
		}
	}
	return isLarger
}
