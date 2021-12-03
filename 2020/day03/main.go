package main

import (
	"bufio"
	"fmt"
	"os"
)

func TreeCounter (puzzleMap []string, stepRight int, stepDown int) int {
	var treeCounter int
	tree := "#"
	colIndex := 0

	for rowIndex, row := range puzzleMap {
		if rowIndex % stepDown != 0 {
			continue
		}

		if string(row[colIndex]) == tree {
			treeCounter++
		}
		colIndex = (colIndex + stepRight) % len(row)
	}
	return treeCounter
}

type Direction struct {
	right int
	down int
}

func main() {
	//Puzzle input
	pwd, _ := os.Getwd()
	f, _ := os.Open(pwd + "/day03/input.txt")

	scanner := bufio.NewScanner(f)
	var puzzleInput []string
	for scanner.Scan() {
		puzzleInput = append(puzzleInput, scanner.Text())
	}

	//Puzzle 1
	result1 := TreeCounter(puzzleInput, 3, 1)
	fmt.Println("Puzzle 1 =", result1)

	//Puzzle 2
	directions := []Direction{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	result2 := 1
	for _, dir := range directions {
		result2 *= TreeCounter(puzzleInput, dir.right, dir.down)
	}
	fmt.Println("Puzzle 2 =", result2)

}