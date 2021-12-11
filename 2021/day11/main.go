package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//Puzzle 1
	octopuses := GetPuzzleInput()
	var flashesTotal int
	for i := 0; i < 100; i++ {
		flashesTotal += Step(octopuses)
	}
	fmt.Println("Puzzle 1 =", flashesTotal)

	//Puzzle 2
	octopuses = GetPuzzleInput()
	var index int
	for {
		Step(octopuses)
		if AllOctopusesFlash(octopuses) {
			break
		}
		index++
	}

	fmt.Println("Puzzle 2 =", index + 1)
}

func AllOctopusesFlash(octopuses [][]int) bool {
	for _, row := range octopuses {
		for _, octopus := range row {
			if octopus != 0 {
				return false
			}
		}
	}
	return true
}

func Step(octopuses [][]int) int {
	for i, row := range octopuses {
		for j, _ := range row {
			octopuses[i][j] += 1
		}
	}
	var flashes int
	for {
		result := FlashedOctopus(octopuses)
		if result == 0 {
			break
		}
		flashes += result
	}
	return flashes
}

func FlashedOctopus(octopuses [][]int) int {
	for i, row := range octopuses {
		for j, octopus := range row {
			if octopus > 9 {
				octopuses[i][j] = 0
				for x:= i-1; x <= i+1; x++ {
					for y := j-1; y <= j+1; y++ {
						if x < 0 || x >= len(octopuses) || y < 0 || y >= len(octopuses[0]) {
							continue
						}
						if octopuses[x][y] == 0 {
							continue
						}
						octopuses[x][y] += 1
					}
				}
				return 1
			}
		}
	}
	return 0
}

func GetPuzzleInput() [][]int {
	pwd, _ := os.Getwd()
	f, _ := os.Open(pwd + "/2021/day11/input.txt")

	scanner := bufio.NewScanner(f)
	var puzzleInput [][]int
	for scanner.Scan() {
		line := scanner.Text()
		var row []int
		for _, value := range line {
			num, _ := strconv.Atoi(string(value))
			row = append(row, num)
		}
		puzzleInput = append(puzzleInput, row)
	}
	return puzzleInput
}
