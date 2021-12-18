package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//Puzzle 1
	grid := GetPuzzleInput()
	nodes := GetListNodes(grid)
	start := "0 0"
	end := fmt.Sprintf("%d %d", len(grid) - 1, len(grid) - 1)
	result := DijkstraShortestPath(nodes, start)
	fmt.Println("Puzzle 1 =", result[end])

	//Puzzle 2
	grid = GetPuzzleInput()
	grid = ExtendGrid(grid)
	nodes = GetListNodes(grid)
	start = "0 0"
	end = fmt.Sprintf("%d %d", len(grid) - 1, len(grid) - 1)
	result = DijkstraShortestPath(nodes, start)
	fmt.Println("Puzzle 2 =", result[end])


}

func ExtendGrid(grid [][]int) [][]int {
	var newGrid [][]int
	origSize := len(grid)
	size := origSize * 5
	for i := 0; i < size; i++ {
		var newRow []int
		for j := 0; j < size; j++ {
			var newItem int
			if i < origSize && j < origSize {
				newItem = grid[i][j]
			} else {
				newItem = grid[i % origSize][j % origSize] + i / origSize + j / origSize
				for newItem > 9 {
					newItem -= 9
				}
			}
			newRow = append(newRow, newItem)
		}
		newGrid = append(newGrid, newRow)
	}
	return newGrid
}

func DijkstraShortestPath(nodes map[string]Node, pointer string) map[string]int {
	visited := make(map[string]bool)
	temp := make(map[string]bool)
	distance := make(map[string]int)
	distance[pointer] = 0
	for {
		visited[pointer] = true
		for _, coor := range nodes[pointer].Children {
			if _, ok := visited[coor]; ok {
				continue
			}
			temp[coor] = true
			newDistance := distance[pointer] + nodes[coor].Value
			if _, ok := distance[coor]; !ok || newDistance < distance[coor] {
				distance[coor] = newDistance
			}
		}
		if _, ok := temp[pointer]; ok {
			delete(temp, pointer)
		}
		if len(temp) == 0 {
			break
		}
		maximum := 100000
		minDistance := maximum
		var index string

		for node, _ := range temp {
			if minDistance == maximum || distance[node] < minDistance {
				minDistance = distance[node]
				index = node
			}
		}

		pointer = index
	}
	return distance
}

type Node struct {
	Coor 		string
	Value		int
	Children	[]string
}

func GetListNodes(grid [][]int) map[string]Node {
	nodes := make(map[string]Node)
	for i, row := range grid {
		for j, _ := range row {
			coor := fmt.Sprintf("%d %d", i, j)
			a := Node{coor, grid[i][j], []string{}}
			var directions = [][]int{
				{i + 1, j},
				{i, j + 1},
				{i - 1, j},
				{i, j - 1},
			}
			for _, item := range directions {
				newI, newJ := item[0], item[1]
				if newI >= 0 && newI < len(grid) && newJ >= 0 && newJ < len(grid) {
					childrenCoor := fmt.Sprintf("%d %d", newI, newJ)
					a.Children = append(a.Children, childrenCoor)
				}
			}
			nodes[coor] = a
		}
	}
	return nodes
}

func GetPuzzleInput() [][]int {
	pwd, _ := os.Getwd()
	f, _ := os.Open(pwd + "/2021/day15/input.txt")

	scanner := bufio.NewScanner(f)
	var grid [][]int

	for scanner.Scan() {
		line := scanner.Text()
		var newLine []int
		for _, item := range line {
			num, _ := strconv.Atoi(string(item))
			newLine = append(newLine, num)
		}
		grid = append(grid, newLine)
	}
	return grid
}
