package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Puzzle input
	mainState := GetPuzzleInput()

	//Puzzle 1
	rounds := 80
	result := LanternfishCycle(mainState, rounds)
	fmt.Println("Puzzle 1 =", result)

	//Puzzle 2
	rounds = 256
	result = LanternfishCycle(mainState, rounds)
	fmt.Println("Puzzle 2 =", result)
}

func LanternfishCycle(mainState map[int]int, rounds int) int {
	for i := 0; i < rounds; i++ {
		temp := make(map[int]int)
		for k, v := range mainState {
			newState := 0
			if k == 0 {
				newState = 6
				AddNumIntoMap(temp, newState, v)
				AddNumIntoMap(temp, 8, v)
			} else {
				newState = k - 1
				AddNumIntoMap(temp, newState, v)
			}
		}
		mainState = make(map[int]int)
		for k, v := range temp {
			mainState[k] = v
		}
	}
	var result int
	for _, v := range mainState {
		result += v
	}
	return result
}

func GetPuzzleInput() map[int]int {
	pwd, _ := os.Getwd()
	dataBytes, _ := os.ReadFile(pwd + "/2021/day06/input.txt")
	raw := strings.Split(string(dataBytes), ",")

	mainState := make(map[int]int)
	for _, v := range raw {
		num, _ := strconv.Atoi(v)
		AddNumIntoMap(mainState, num, 1)
	}
	return mainState
}

func AddNumIntoMap(state map[int]int, num int, count int) {
	var numInMap bool
	for key := range state {
		if key == num {
			state[key] += count
			numInMap = true
		}
	}
	if !numInMap {
		state[num] = count
	}
}
