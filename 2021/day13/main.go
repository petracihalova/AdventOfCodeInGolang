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
	coordinates, foldInstructions := GetPuzzleInput()

	//Puzzle 1
	var puzzle1Result int
	for i, item := range foldInstructions {
		rawInstruction := strings.Split(item, "=")
		direction := rawInstruction[0]
		value, _ := strconv.Atoi(rawInstruction[1])
		var index int
		if direction == "x" {
			index = 0
		} else {
			index = 1
		}
		var temp [][]int
		for _, coor := range coordinates {
			if coor[index] < value {
				temp = append(temp, coor)
			} else if coor[index] > value {
				newIndex := value - (coor[index] - value)
				x, y := coor[0], coor[1]
				if index == 0 {
					x = newIndex
				} else {
					y = newIndex
				}
				temp = append(temp, []int{x, y})
			}
		}
		coordinates = Unique(temp)

		if i == 0 {
			puzzle1Result = len(coordinates)
		}
	}
	fmt.Println("Puzzle 1 =", puzzle1Result)

	//Puzzle 2
	fmt.Println("Puzzle 2 =")
	PrintCoordinates(coordinates)
}

func Unique(coordinates [][]int) [][]int {
	var temp [][]int
	for _, coor := range coordinates {
		if !CoorInList(coor[0], coor[1], temp) {
			temp = append(temp, coor)
		}
	}
	return temp
}

func PrintCoordinates(coordinates [][]int) {
	var sizeX, sizeY int
	for _, coor := range coordinates {
		x, y := coor[0], coor[1]
		if x > sizeX {
			sizeX = x
		}
		if y > sizeY {
			sizeY = y
		}
	}
	for j := 0; j <= sizeY; j++ {
		for i := 0; i <= sizeX; i++ {
			if CoorInList(i, j, coordinates) {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func CoorInList(i int, j int, coordinates [][]int) bool {
	for _, coor := range coordinates {
		if coor[0] == i && coor[1] == j {
			return true
		}
	}
	return false
}

func GetPuzzleInput() ([][]int, []string) {
	pwd, _ := os.Getwd()
	f, _ := os.Open(pwd + "/2021/day13/input.txt")

	scanner := bufio.NewScanner(f)
	var coordinates [][]int
	var foldInstructions []string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "fold along") {
			foldInstructions = append(foldInstructions, strings.Replace(line, "fold along ", "", 1))
		} else if strings.Contains(line, ",") {
			numbers := strings.Split(line, ",")
			num1, _ := strconv.Atoi(numbers[0])
			num2, _ := strconv.Atoi(numbers[1])
			coordinates = append(coordinates, []int{num1, num2})
		}
	}
	return coordinates, foldInstructions
}
