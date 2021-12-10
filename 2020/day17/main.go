package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//Puzzle input
	puzzleInput := GetPuzzleInput()
	grid := [][][]string{puzzleInput}

	//Puzzle 1
	result1 := Puzzle1(grid)
	fmt.Println("Puzzle 1 =", result1)


	//Puzzle 2
	grid4D := [][][][]string{grid}
	result2 := Puzzle2(grid4D)
	fmt.Println("Puzzle2 =", result2)

}

func GetPuzzleInput() [][]string {
	pwd, _ := os.Getwd()
	f, _ := os.Open(pwd + "/2020/day17/input.txt")

	scanner := bufio.NewScanner(f)
	var puzzleInput [][]string
	for scanner.Scan() {
		puzzleInput = append(puzzleInput, strings.Split(scanner.Text(), ""))
	}

	return puzzleInput
}
