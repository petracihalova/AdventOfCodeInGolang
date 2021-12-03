package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Puzzle input
	pwd, _ := os.Getwd()
	dataBytes, _ := os.ReadFile(pwd + "/day15/input.txt")

	raw := strings.Split(string(dataBytes), ",")
	data := CreateStartingNumbers(raw)

	//Puzzle 1 - style = brute force
	turn := len(data)
	end := 2020
	for i := turn; i < end; i++ {
		newNum := Turn(data)
		data = append(data, newNum)
	}
	fmt.Println("Puzzle 1 =", data[len(data)-1])

	//Puzzle 2
	//Brute force not applicable = take too long
	//I use map to save last 2 indexes of all numbers
	results := make(map[int][]int) //key = number, value = last (max 2) indexes
	data = CreateStartingNumbers(raw)

	for i, v := range data {
		results[v] = []int {i}
	}

	turn = len(data)
	end = 30000000
	last := data[len(data) - 1]

	for i := turn; i < end; i++ {
		if len(results[last]) == 1 {
			results[0] = append(results[0], i)
			if len(results[0]) > 2 {
				results[0] = results[0][1:]
			}
			last = 0
			continue
		}
		newNumber := results[last][1] - results[last][0]
		if _, ok := results[newNumber]; ok {
			if len(results[newNumber]) > 1 {
				results[newNumber] = results[newNumber][1:]
			}
		} else {
			results[newNumber] = []int {}
		}
		results[newNumber] = append(results[newNumber], i)
		last = newNumber
	}

	fmt.Println("Puzzle 2 =", last)
}

func Turn(data []int) int {
	last := data[len(data) - 1]
	previousIndex := GetPreviousIndex(last, data[:len(data)-1])
	if previousIndex == -1 {
		return 0
	}
	return len(data) - 1 - previousIndex
}

func GetPreviousIndex(last int, data []int) int {
	for i := len(data) - 1; i >= 0; i-- {
		if data[i] == last {
			return i
		}
	}
	return -1
}

func CreateStartingNumbers(raw []string) []int {
	var data []int
	for _, v := range raw {
		num, _ := strconv.Atoi(v)
		data = append(data, num)
	}
	return data
}
