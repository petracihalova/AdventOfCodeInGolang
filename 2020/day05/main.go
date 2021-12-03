package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//Puzzle input
	pwd, _ := os.Getwd()
	f, _ := os.Open(pwd + "/day05/input.txt")

	scanner := bufio.NewScanner(f)
	var puzzleInput []string
	for scanner.Scan() {
		puzzleInput = append(puzzleInput, scanner.Text())
	}

	//Puzzle 1
	var highestSeatId int64
	for _, seat := range puzzleInput {
		result := GetSeatId(seat)
		if result > highestSeatId {
			highestSeatId = result
		}
	}
	fmt.Println("Puzzle 1 =", highestSeatId)

	//Puzzle 2
	var seatIds []int64
	for _, seat := range puzzleInput {
		seatIds = append(seatIds, GetSeatId(seat))
	}
	sort.Slice(seatIds, func(i, j int) bool { return seatIds[i] < seatIds[j] })

	var mySeatId int64
	for i := 1; i < len(seatIds) - 1; i++ {
		if seatIds[i] - seatIds[i-1] != 1 {
			mySeatId = seatIds[i] - 1
			break
		}
	}
	fmt.Println("Puzzle 2 =", mySeatId)
}

func GetSeatId(seat string) int64 {
	row, column := seat[:7], seat[7:]
	row = strings.Replace(row, "F", "0", -1)
	row = strings.Replace(row, "B", "1", -1)

	rowNumber, _ := strconv.ParseInt(row, 2, 64)

	column = strings.Replace(column, "L", "0", -1)
	column = strings.Replace(column, "R", "1", -1)

	colNumber, _ := strconv.ParseInt(column, 2, 64)

	return rowNumber * 8 + colNumber
}