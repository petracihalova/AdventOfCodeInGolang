package main

import (
	"fmt"
	"strings"
)

func Puzzle1(grid [][][]string) int {
	cyclesCount := 6
	for i:= 0; i < cyclesCount; i++ {
		grid = Cycle3D(grid)
	}
	result := CountActiveCubes3D(grid)
	return result
}

func CountActiveCubes3D(grid [][][]string) int{
	var activeCubes int
	for _, layer := range grid {
		for _, row := range layer {
			for _, cube := range row {
				if cube == "#" {
					activeCubes++
				}
			}
		}
	}
	return activeCubes
}

func Cycle3D(grid [][][]string) [][][]string {
	grid = CoverGrid3D(grid)
	grid = NewState3D(grid)
	return grid
}

func NewState3D(grid [][][]string) [][][]string {
	var newGrid [][][]string
	for z, layer := range grid {
		var newLayer [][]string
		for y, row := range layer {
			var newRow []string
			for x, cube := range row {
				activeNeighbors := CountActiveNeighbors3D(grid, x, y, z)
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
		newGrid = append(newGrid, newLayer)
	}
	return newGrid
}

func CountActiveNeighbors3D(grid [][][]string, x int, y int, z int) int {
	var activeNeighbors int
	for newZ := z - 1; newZ <= z + 1; newZ++ {
		for newY := y - 1; newY <= y + 1; newY++ {
			for newX := x - 1; newX <= x + 1; newX++ {
				if newX == x && newY == y && newZ == z {
					continue
				}
				if newX < 0 || newX >= len(grid[0][0]) {
					continue
				}
				if newY < 0 || newY >= len(grid[0]) {
					continue
				}
				if newZ < 0 || newZ >= len(grid) {
					continue
				}
				if grid[newZ][newY][newX] == "#" {
					activeNeighbors++
				}
			}
		}
	}
	return activeNeighbors
}

func CoverGrid3D(grid [][][]string) [][][]string {
	var newGrid [][][]string
	emptyRow := strings.Split(strings.Repeat(".", len(grid[0][0])+2), "")
	var emptyLayer [][]string
	for i := 0; i < len(grid[0]) + 2; i++ {
		emptyLayer = append(emptyLayer, emptyRow)
	}
	newGrid = append(newGrid, emptyLayer)
	for _, layer := range grid {
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
		newGrid = append(newGrid, newLayer)
	}
	newGrid = append(newGrid, emptyLayer)
	return newGrid
}

func Print3DGrid3D(grid [][][]string) {
	for z, layer := range grid {
		fmt.Println("z =", z)
		for _, row := range layer {
			fmt.Println(row)
		}
		fmt.Println()
	}
	fmt.Println("----------END---------")
	fmt.Println()
}