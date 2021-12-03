package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	//Puzzle input
	pwd, _ := os.Getwd()
	f, _ := os.Open(pwd + "/day09/input.txt")

	scanner := bufio.NewScanner(f)
	var puzzleInput []int
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		puzzleInput = append(puzzleInput, num)
	}

	//Puzzle 1
	preambleLength := 25
	var result int
	result = FindFirstWrong(puzzleInput, preambleLength)
	fmt.Println("Puzzle 1 =", result)

	//Puzzle 1
	invalidNumber := result
	var finalNumbers []int
	finalNumbers = FindContiguosSet(puzzleInput, invalidNumber)
	sort.Ints(finalNumbers)
	fmt.Println("Puzzle 2 =", finalNumbers[0] + finalNumbers[len(finalNumbers)-1])

}

func FindContiguosSet(input []int, invalidNumber int) []int {
	for i := 0; i < len(input) - 1; i++ {
		var sum int
		for j := i; j < len(input); j++ {
			sum += input[j]
			if sum == invalidNumber {
				return input[i:j+1]
			}
			if sum > invalidNumber {
				break
			}
		}
	}
	return nil
}

func FindFirstWrong(input []int, preambleLength int) int {
	var num int
	for i := preambleLength; i < len(input); i++ {
		var wrong bool
		var preamble []int
		for j := 0; j < preambleLength; j++ {
			preamble = append(preamble, input[i-1-j])
		}
		num = input[i]
		wrong = IsWrong(preamble, num)
		if wrong {
			break
		}
	}
	return num
}

func IsWrong(preamble []int, num int) bool {
	result := true
	for i := 0; i < len(preamble); i++ {
		for j := 0; j < len(preamble); j++ {
			if i == j {
				continue
			}
			if preamble[i] + preamble[j] == num {
				result = false
			}
		}
	}
	return result
}
