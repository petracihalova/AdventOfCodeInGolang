package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	//Puzzle input
	pwd, _ := os.Getwd()
	f, _ := os.Open(pwd + "/2020/day10/input.txt")

	scanner := bufio.NewScanner(f)
	var adapters []int
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		adapters = append(adapters, num)
	}

	//Puzzle 1
	adapters = append(adapters, 0)
	sort.Ints(adapters)
	myDevice := adapters[len(adapters) - 1] + 3
	adapters = append(adapters, myDevice)

	var counter1 int
	var counter3 int

	for i := 1; i < len(adapters); i++ {
		switch adapters[i] - adapters[i-1] {
		case 1:
			counter1++
		case 3:
			counter3++
		}
	}
	fmt.Println("Puzzle 1 =", counter1 * counter3)

	//Puzzle 2
	//Sepparated groups of numbers, eg 4:2 means that group of 4 consecutive numbers is 2x in the data
	groups := make(map[int]int)
	var groupSize int
	var start int
	for i := 0; i < len(adapters) - 1; i++ {
		if adapters[i + 1] - adapters[i] == 3 {
			groupSize = len(adapters[start:i+1])
			start = i + 1
			if groupSize <= 2 {
				continue
			}
			groups[groupSize] += 1
		}

	}
	var result int64
	var multiplicator int64
	result = 1
	for key, value := range groups {
		switch key {
		case 3:
			multiplicator = int64(math.Pow(2, float64(value)))
		case 4:
			multiplicator = int64(math.Pow(4, float64(value)))
		case 5:
			multiplicator = int64(math.Pow(7, float64(value)))
		}
		result = result * multiplicator
	}
	fmt.Println("Puzzle 2 =", result)
}
