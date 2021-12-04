package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	numbers 	[][]string
	isBingo		 bool
}

func (b Board) MarkNumber(num string) {
	for r, row := range b.numbers {
		for v, value := range row {
			if value == "X" {
				continue
			}
			if value == num {
				b.numbers[r][v] = "X"
			}
		}
	}
}

func (b Board) CheckBingo() bool {
	for _, row := range b.numbers {
		if strings.Join(row, "") == "XXXXX" {
			return true
		}
	}
	transposeBoard := transpose(b.numbers)
	for _, row := range transposeBoard {
		if strings.Join(row, "") == "XXXXX" {
			return true
		}
	}
	return false
}

func main() {
	//Puzzle input
	numbers, boards := GetPuzzleInput()

	//Puzzle 1
	var foundBingo bool
	var lastNumber, winnerIndex int
	var winnersPuzzle1 []int
	for _, num := range numbers {
		Round(boards, num)
		CheckBingo(boards, winnersPuzzle1)
		for b, board := range boards {
			if board.isBingo {
				foundBingo = true
				lastNumber, _ = strconv.Atoi(num)
				winnerIndex = b
				break
			}
		}
		if foundBingo {
			break
		}
	}
	sumUnmarkedNumbers := GetUnmarkedNumbers(boards[winnerIndex])
	fmt.Println("Puzzle 1 =", sumUnmarkedNumbers * lastNumber)

	//Puzzle 2
	numbers, boards = GetPuzzleInput()
	lastNumber = 0
	var winnersPuzzle2 []int
	for _, num := range numbers {
		Round(boards, num)
		CheckBingo(boards, winnersPuzzle2)
		for b, board := range boards {
			if BoardInWinners(b, winnersPuzzle2) {
				continue
			}
			if board.isBingo {
				lastNumber, _ = strconv.Atoi(num)
				winnersPuzzle2 = append(winnersPuzzle2, b)
			}
		}
		if len(winnersPuzzle2) == len(boards) {
			break
		}
	}
	lastWinnerIndex := winnersPuzzle2[len(winnersPuzzle2) - 1]
	sumUnmarkedNumbers = GetUnmarkedNumbers(boards[lastWinnerIndex])
	fmt.Println("Puzzle 1 =", sumUnmarkedNumbers * lastNumber)
}

func BoardInWinners(b int, winners []int) bool {
	for _, v := range winners {
		if v == b {
			return true
		}
	}
	return false
}

func GetUnmarkedNumbers(board Board) int {
	var sum int
	for _, row := range board.numbers {
		for _, value := range row {
			num, _ := strconv.Atoi(value)
			sum += num
		}
	}
	return sum
}

func CheckBingo(boards []Board, winners []int) {
	for i, board := range boards {
		if board.CheckBingo() {
			boards[i].isBingo = true
			winners = append(winners, i)
		}
	}
}

func Round(boards []Board, num string) {
	for _, board := range boards {
		board.MarkNumber(num)
	}
}

func GetPuzzleInput() ([]string, []Board ) {
	pwd, _ := os.Getwd()
	dataBytes, _ := os.ReadFile(pwd + "/2021/day04/input.txt")
	raw := strings.Split(string(dataBytes), "\n\n")
	numbers := GetNumbers(raw[0])
	boards := GetBoards(raw[1:])

	return numbers, boards
}

func GetBoards(input []string) []Board {
	var boards []Board
	for _, b := range input {
		var board Board
		var numbers [][]string
		for _, r := range strings.Split(b, "\n") {
			var row []string
			r = strings.Replace(r, "  ", " ", -1)
			r = strings.Trim(r, " ")
			for _, num := range strings.Split(r, " ") {
				row = append(row, num)
			}
			numbers = append(numbers, row)
		}
		board.numbers = numbers
		board.isBingo = false
		boards = append(boards, board)
	}
	return boards
}

func GetNumbers(input string) []string {
	var numbers []string
	for _, num := range strings.Split(input, ",") {
		numbers = append(numbers, num)
	}
	return numbers
}

func PrintAllBoards(boards []Board) {
	for _, board := range boards {
		PrintBoard(board)
		fmt.Println()
	}
}

func PrintBoard(board Board) {
	for _, row := range board.numbers {
		for _, num := range row {
			fmt.Printf("%3s", num)
		}
		fmt.Println()
	}
}

//Source of transpose func = https://gist.github.com/tanaikech/5cb41424ff8be0fdf19e78d375b6adb8
func transpose(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}