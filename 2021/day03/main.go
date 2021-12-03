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
	f, _ := os.Open(pwd + "/day03/input.txt")

	scanner := bufio.NewScanner(f)
	var puzzleInput []int
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		puzzleInput = append(puzzleInput, num)
	}
	fmt.Println(puzzleInput)

	//Puzzle 1


	//Puzzle 2
	
}
