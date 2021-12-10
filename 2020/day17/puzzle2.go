package main

import (
	"fmt"
	"strings"
)

func Puzzle2(grid [][][][]string) int {
	cyclesCount := 6
	for i:= 0; i < cyclesCount; i++ {
		grid = Cycle4D(grid)
	}
	result := CountActiveCubes4D(grid)
	return result
}

func CountActiveCubes4D(grid [][][][]string) int{
	var activeCubes int
	for _, dimension := range grid {
		for _, layer := range dimension {
			for _, row := range layer {
				for _, cube := range row {
					if cube == "#" {
						activeCubes++
					}
				}
			}
		}
	}
	return activeCubes
}

func Cycle4D(grid4D [][][][]string) [][][][]string {
	grid4D = CoverGrid4D(grid4D)
	grid4D = NewState4D(grid4D)
	return grid4D
}

func NewState4D(grid4D [][][][]string) [][][][]string {
	var newGrid4D [][][][]string
	for w, dimension := range grid4D {
		var newDimension [][][]string
		for z, layer := range dimension {
			var newLayer [][]string
			for y, row := range layer {
				var newRow []string
				for x, cube := range row {
					activeNeighbors := CountActiveNeighbors4D(grid4D, x, y, z, w)
					if cube == "#" {
						//Cube is active
						if activeNeighbors == 2 || activeNeighbors == 3 {
							//cube remains active
							newRow = append(newRow, "#")
						} else {
							//cube becomes inactive
							newRow = append(newRow, ".")
						}
					} else if cube == "." {
						//Cube is inactive
						if activeNeighbors == 3 {
							//cube becomes active
							newRow = append(newRow, "#")
						} else {
							//cube remains inactive
							newRow = append(newRow, ".")
						}
					}
				}
				newLayer = append(newLayer, newRow)
			}
			newDimension = append(newDimension, newLayer)
		}
		newGrid4D = append(newGrid4D, newDimension)
	}
	return newGrid4D
}

func CountActiveNeighbors4D(grid4D [][][][]string, x int, y int, z int, w int) int {
	var activeNeighbors int
	for newW := w - 1; newW <= w + 1; newW++ {
		for newZ := z - 1; newZ <= z+1; newZ++ {
			for newY := y - 1; newY <= y+1; newY++ {
				for newX := x - 1; newX <= x+1; newX++ {
					if newX == x && newY == y && newZ == z && newW == w {
						continue
					}
					if newX < 0 || newX >= len(grid4D[0][0][0]) {
						continue
					}
					if newY < 0 || newY >= len(grid4D[0][0]) {
						continue
					}
					if newZ < 0 || newZ >= len(grid4D[0]) {
						continue
					}
					if newW < 0 || newW >= len(grid4D) {
						continue
					}
					if grid4D[newW][newZ][newY][newX] == "#" {
						activeNeighbors++
					}
				}
			}
		}
	}
	return activeNeighbors
}

func CoverGrid4D(grid4D [][][][]string) [][][][]string {
	var newGrid4D [][][][]string

	emptyRow := strings.Split(strings.Repeat(".", len(grid4D[0][0][0])+2), "")

	var emptyLayer [][]string
	for i := 0; i < len(grid4D[0][0]) + 2; i++ {
		emptyLayer = append(emptyLayer, emptyRow)
	}

	var emptyDimension [][][]string
	for i := 0; i < len(grid4D[0]) + 4; i++ {
		emptyDimension = append(emptyDimension, emptyLayer)
	}

	newGrid4D = append(newGrid4D, emptyDimension)

	for _, dimension := range grid4D {
		var newDimension [][][]string
		newDimension = append(newDimension, emptyLayer)
		newDimension = append(newDimension, emptyLayer)
		for _, layer := range dimension {
			var newLayer [][]string
			newLayer = append(newLayer, emptyRow)
			for _, row := range layer {
				newRow := []string{"."}
				for _, item := range row {
					newRow = append(newRow, item)
				}
				newRow = append(newRow, ".")
				newLayer = append(newLayer, newRow)
			}
			newLayer = append(newLayer, emptyRow)
			newDimension = append(newDimension, newLayer)
		}
		newDimension = append(newDimension, emptyLayer)
		newDimension = append(newDimension, emptyLayer)
		newGrid4D = append(newGrid4D, newDimension)
	}
	newGrid4D = append(newGrid4D, emptyDimension)

	return newGrid4D
}

func Print3DGrid4D(grid [][][][]string) {
	for w, dimension := range grid {
		fmt.Print("w =", w, " ")
		for z, layer := range dimension {
			fmt.Println("z =", z)
			for _, row := range layer {
				fmt.Println(row)
			}
			fmt.Println()
		}
		fmt.Println("----------dimension---------")
	}
	fmt.Println("----------END---------")
	fmt.Println()
}