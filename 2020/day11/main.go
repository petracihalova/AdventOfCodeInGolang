package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//Puzzle input
	pwd, _ := os.Getwd()
	f, _ := os.Open(pwd + "/day11/input.txt")

	scanner := bufio.NewScanner(f)
	var inputData [][]string
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		inputData = append(inputData, line)
	}

	//Puzzle 1
	seatLayout := inputData
	var newSeatLayout [][]string
	for {
		newSeatLayout = CreateNewState(seatLayout, 1)
		if LayoutsAreSame(seatLayout, newSeatLayout) {
			break
		} else {
			seatLayout = newSeatLayout
		}
	}

	var occupiedCount int
	for i := 0; i < len(newSeatLayout); i++ {
		for j := 0; j < len(newSeatLayout[0]); j++ {
			if newSeatLayout[i][j] == "#" {
				occupiedCount++
			}
		}
	}

	fmt.Println("Puzzle 1 =", occupiedCount)

	//Puzzle 2
	seatLayout = inputData
	for {
		newSeatLayout = CreateNewState(seatLayout, 2)
		if LayoutsAreSame(seatLayout, newSeatLayout) {
			break
		} else {
			seatLayout = newSeatLayout
		}
	}

	occupiedCount = 0
	for i := 0; i < len(newSeatLayout); i++ {
		for j := 0; j < len(newSeatLayout[0]); j++ {
			if newSeatLayout[i][j] == "#" {
				occupiedCount++
			}
		}
	}

	fmt.Println("Puzzle 2 =", occupiedCount)
}

func LayoutsAreSame(layout [][]string, layout2 [][]string) bool {
	areSame := true
	for i := 0; i < len(layout); i++ {
		for j := 0; j < len(layout[0]); j++ {
			if layout[i][j] != layout2[i][j] {
				areSame = false
				break
			}
		}
	}
	return areSame
}

func CreateNewState(layout [][]string, puzzle int) [][]string {
	var newLayout [][]string
	for i, row := range layout {
		var newRow []string
		for j := range row {
			newRow = append(newRow, NewState(layout, i, j, puzzle))
		}
		newLayout = append(newLayout, newRow)
	}
	return newLayout
}

func NewState(layout [][]string, i int, j int, puzzle int) string {
	old := layout[i][j]
	switch old {
	case ".":
		return old
	case "L":
		if puzzle == 1 && OccupiedSeatsPuzzle1(layout, i, j) == 0 {
			return "#"
		}
		if puzzle == 2 && OccupiedSeatsPuzzle2(layout, i, j) == 0 {
			return "#"
		}
		return old
	case "#":
		if puzzle == 1 && OccupiedSeatsPuzzle1(layout, i, j) >= 4 {
			return "L"
		}
		if puzzle == 2 && OccupiedSeatsPuzzle2(layout, i, j) >= 5 {
			return "L"
		}
		return old
	default:
		panic("unexpected character")
	}
}

type Direction struct {
	Name string
	Row int
	Column int
}

func OccupiedSeatsPuzzle2(layout [][]string, i int, j int) int {
	var count int
	directions := []Direction {
		{Name: "up", Row: -1, Column: 0},
		{Name: "upRight", Row: -1, Column: 1},
		{Name: "right", Row: 0, Column: 1},
		{Name: "downRight", Row: 1, Column: 1},
		{Name: "down", Row: 1, Column: 0},
		{Name: "downLeft", Row: 1, Column: -1},
		{Name: "left", Row: 0, Column: -1},
		{Name: "upLeft", Row: -1, Column: -1},
	}

	for _, dir := range directions {
		for rowIndex, colIndex := i+dir.Row, j+dir.Column;
		rowIndex >= 0 && rowIndex < len(layout) && colIndex >= 0 && colIndex < len(layout[0]);
		rowIndex, colIndex = rowIndex + dir.Row, colIndex + dir.Column {
			if rowIndex == i && colIndex == j {
				continue
			}
			if layout[rowIndex][colIndex] == "#" {
				count++
				break
			}
			if layout[rowIndex][colIndex] == "L" {
				break
			}
		}
	}
	return count
}

func OccupiedSeatsPuzzle1(layout [][]string, i int, j int) int {
	var count int
	for rowIndex := i-1; rowIndex <= (i+1); rowIndex++ {
		if rowIndex < 0 || rowIndex >= len(layout) {
			continue
		}
		for colIndex := j-1; colIndex <= (j+1); colIndex++ {
			if colIndex < 0 || colIndex >= len(layout[0]) {
				continue
			}
			if rowIndex == i && colIndex == j {
				continue
			}
			if layout[rowIndex][colIndex] == "#" {
				count++
			}
		}
	}
	return count
}
