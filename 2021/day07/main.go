package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Puzzle input
	crabs := GetPuzzleInput()

	//Puzzle 1
	var bestPosition int
	var fuelBestPosition int
	for i := 0; i <= MaxInteger(crabs); i++ {
		var fuel int
		for _, crab := range crabs {
			fuel += int(math.Abs(float64(crab - i)))
		}
		if i == 0 || fuel < fuelBestPosition {
			bestPosition = i
			fuelBestPosition = fuel
		}
	}
	fmt.Println("Puzzle 1: Best position = ",bestPosition, "with fuel =", fuelBestPosition)

	//Puzzle 2
	biggestPosition := MaxInteger(crabs)
	listFuelConsumption := GetFuelConsumption(biggestPosition)

	for i := 0; i <= MaxInteger(crabs); i++ {
		var fuel int
		for _, crab := range crabs {
			add := listFuelConsumption[int(math.Abs(float64(crab - i)))]
			fuel += add
		}
		if i == 0 || fuel < fuelBestPosition {
			bestPosition = i
			fuelBestPosition = fuel
		}
	}

	fmt.Println("Puzzle 2: Best position = ",bestPosition, "with fuel =", fuelBestPosition)
}

func GetFuelConsumption(max int) []int {
	var consumptions []int
	for i := 0; i <= max; i++ {
		var fuel int
		for j := 0; j <= i; j++ {
			fuel += j
		}
		consumptions = append(consumptions, fuel)
	}
	return consumptions
}

func MaxInteger(numbers []int) int {
	var max int
	for _, v := range numbers {
		if v > max {
			max = v
		}
	}
	return max
}

func GetPuzzleInput() []int {
	pwd, _ := os.Getwd()
	dataBytes, _ := os.ReadFile(pwd + "/2021/day07/input.txt")
	raw := strings.Split(string(dataBytes), ",")

	var crabs []int
	for _, v := range raw {
		num, _ := strconv.Atoi(v)
		crabs = append(crabs, num)
	}
	return crabs
}