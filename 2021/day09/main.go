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
	points := GetPuzzleInput()

	//Puzzle 1
	var riskLevel int
	var lowPoints [][]int
	for i, row := range points {
		for j, value := range row {
			if IsLowPoint(points, i, j) {
				riskLevel += value + 1
				lowPoints = append(lowPoints, []int{i, j}) //This is needed for Puzzle 2
			}
		}
	}
	fmt.Println("Puzzle 1 =", riskLevel)

	//Puzzle 2
	var basinSizes []int
	for _, p := range lowPoints {
		var basin [][]int
		basin = append(basin, []int{p[0], p[1]})
		basin = FindBasinCoordinates(points, basin, p[0], p[1])
		basinSizes = append(basinSizes, len(basin))
	}
	sort.Ints(basinSizes)
	i := len(basinSizes) - 1
	result := basinSizes[i] * basinSizes[i-1] * basinSizes[i-2]
	fmt.Println("Puzzle 2 =", result)
}

func FindBasinCoordinates(points [][]int, basin [][]int, i int, j int) [][]int {
	up := []int{i - 1, j}
	down := []int{i + 1, j}
	left := []int{i, j - 1}
	right := []int{i, j + 1}
	directions := [][]int{up, down, left, right}

	for _, d := range directions {
		x, y := d[1], d[0]
		if x < 0 || x >= len(points[0]) || y < 0 || y >= len(points) {
			continue
		}
		if points[y][x] == 9 {
			continue
		}
		if ItemInSlice(d, basin) {
			continue
		}
		basin = append(basin, d)
		basin = FindBasinCoordinates(points, basin, y, x)
	}
	return basin
}

func ItemInSlice(i []int, basin [][]int) bool {
	for _, b := range basin {
		if i[0] == b[0] && i[1] == b[1] {
			return true
		}
	}
	return false
}

func IsLowPoint(points [][]int, i int, j int) bool {
	up := []int{i - 1, j}
	down := []int{i + 1, j}
	left := []int{i, j - 1}
	right := []int{i, j + 1}
	directions := [][]int{up, down, left, right}

	for _, d := range directions {
		x, y := d[1], d[0]
		if x < 0 || x >= len(points[0]) || y < 0 || y >= len(points) {
			continue
		}
		if points[i][j] >= points[y][x] {
			return false
		}
	}
	return true
}

func GetPuzzleInput() [][]int {
	pwd, _ := os.Getwd()
	f, _ := os.Open(pwd + "/2021/day09/input.txt")

	scanner := bufio.NewScanner(f)
	var puzzleInput [][]int
	for scanner.Scan() {
		var newRow []int
		for _, v := range scanner.Text() {
			num, _ := strconv.Atoi(string(v))
			newRow = append(newRow, num)
		}
		puzzleInput = append(puzzleInput, newRow)
	}
	return puzzleInput
}
