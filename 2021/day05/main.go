package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)
type Vent struct {
	x1		  int
	y1 		  int
	x2 		  int
	y2 		  int
	isDiag 	  bool
	isVert    bool
	isHorizon bool
}

type Position struct {
	x	  int
	y	  int
	count int
}

func main() {
	//Puzzle input
	input := GetPuzzleInput()

	//Puzzle 1
	vents := GetListVents(input)
	var positions []Position
	for _, vent := range vents {
		if vent.isDiag {
			continue
		}
		coordinatesList := GetListCoordinates(vent)
		for _, coor := range coordinatesList {
			x, y := coor[0], coor[1]
			if index, ok := FindCoorInPositionsList(positions, x, y); ok {
				positions[index].count += 1
			} else {
				positions = append(positions, Position{x, y, 1})
			}
		}
	}
	var result int
	for _, pos := range positions {
		if pos.count > 1 {
			result++
		}
	}
	fmt.Println("Puzzle 1 =", result)


	//Puzzle 2
	vents = GetListVents(input)
	positions = []Position{}
	for _, vent := range vents {
		coordinatesList := GetListCoordinates(vent)
		for _, coor := range coordinatesList {
			x, y := coor[0], coor[1]
			if index, ok := FindCoorInPositionsList(positions, x, y); ok {
				positions[index].count += 1
			} else {
				positions = append(positions, Position{x, y, 1})
			}
		}
	}
	result = 0
	for _, pos := range positions {
		if pos.count > 1 {
			result++
		}
	}
	fmt.Println("Puzzle 2 =", result)
}

func FindCoorInPositionsList(positions []Position, x int, y int) (int, bool) {
	for i, pos := range positions {
		if pos.x == x && pos.y == y {
			return i, true
		}
	}
	return 0, false
}

func GetListCoordinates(vent Vent) [][]int {
	var result [][]int
	stepX, stepY := 0, 0
	end := 0
	switch {
	case vent.isHorizon:
		stepX, stepY = 1, 0
		if vent.x1 > vent.x2 {
			stepX = -1
		}
		end = int(math.Abs(float64(vent.x1) - float64(vent.x2)))

	case vent.isVert:
		stepX, stepY = 0, 1
		if vent.y1 > vent.y2 {
			stepY = -1
		}
		end = int(math.Abs(float64(vent.y1) - float64(vent.y2)))

	case vent.isDiag:
		stepX, stepY = 1, 1
		if vent.x1 > vent.x2 {
			stepX = -1
		}
		if vent.y1 > vent.y2 {
			stepY = -1
		}
		end = int(math.Abs(float64(vent.x1) - float64(vent.x2)))
	}

	for multiplicator := 0; multiplicator <= end; multiplicator++ {
		var coor []int
		newCoorX := vent.x1 + multiplicator * stepX
		newCoorY := vent.y1 + multiplicator * stepY
		coor = append(coor, newCoorX)
		coor = append(coor, newCoorY)
		result = append(result, coor)
	}
	return result

}

func GetListVents(input []string) []Vent {
	var vents []Vent
	for _, row := range input {
		row = strings.Replace(row, " -> ", ",", -1)
		coor := strings.Split(row, ",")
		x1, _ := strconv.Atoi(coor[0])
		y1, _ := strconv.Atoi(coor[1])
		x2, _ := strconv.Atoi(coor[2])
		y2, _ := strconv.Atoi(coor[3])
		isDiag := math.Abs(float64(x1) - float64(x2)) == math.Abs(float64(y1) - float64(y2))
		isVert := x1 == x2
		isHorizon := y1 == y2
		vent := Vent{x1, y1, x2, y2, isDiag, isVert, isHorizon}
		vents = append(vents, vent)
	}
	return vents
}

func GetPuzzleInput() []string {
	pwd, _ := os.Getwd()
	f, _ := os.Open(pwd + "/2021/day05/input.txt")

	scanner := bufio.NewScanner(f)
	var puzzleInput []string
	for scanner.Scan() {
		puzzleInput = append(puzzleInput, scanner.Text())
	}
	return puzzleInput
}
