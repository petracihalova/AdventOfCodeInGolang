package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//Puzzle input
	puzzleInput := GetPuzzleInput()

	//Puzzle 1
	var counts1 []int
	for i := 0; i < len(puzzleInput[0]); i++ {
		c0, _ := CountsOfZeroAndOne(i, puzzleInput)
		counts1 = append(counts1, c0)
	}

	gammaRate := ""
	epsilonRate := ""
	for _, num := range counts1 {
		if num > len(puzzleInput) / 2 {
			gammaRate += "1"
			epsilonRate += "0"
		} else {
			gammaRate += "0"
			epsilonRate += "1"
		}
	}
	gammaRateInt, _ := strconv.ParseInt(gammaRate, 2, 64)
	epsilonRateInt, _ := strconv.ParseInt(epsilonRate, 2, 64)
	fmt.Println("Puzzle 1 =", gammaRateInt * epsilonRateInt)

	//Puzzle 2
	puzzleInput = GetPuzzleInput()
	result := GetRating(MostCommonValueOxygenGenRating, puzzleInput)
	oxygenGeneratorRatingInt, _ := strconv.ParseInt(result, 2, 64)

	puzzleInput = GetPuzzleInput()
	result = GetRating(MostCommonValueCo2ScrubberRating, puzzleInput)
	co2ScrubberRatingInt, _ := strconv.ParseInt(result, 2, 64)

	fmt.Println("Puzzle 2 =", oxygenGeneratorRatingInt * co2ScrubberRatingInt)
}

func GetPuzzleInput() []string {
	pwd, _ := os.Getwd()
	f, _ := os.Open(pwd + "/2021/day03/input.txt")

	scanner := bufio.NewScanner(f)
	var puzzleInput []string
	for scanner.Scan() {
		puzzleInput = append(puzzleInput, scanner.Text())
	}
	return puzzleInput
}

func GetRating(MostCommonValueFunc func(pointer int, data []string) string, data []string) string {
	var pointer int
	for {
		var temp []string
		m := MostCommonValueFunc(pointer, data)
		for _, num := range data {
			if string(num[pointer]) == m {
				temp = append(temp, num)
			}
		}

		data = temp
		if len(data) == 1 {
			break
		}
		pointer++
	}
	return data[0]
}

func MostCommonValueOxygenGenRating(pointer int, data []string) string {
	counts0, counts1 := CountsOfZeroAndOne(pointer, data)
	if counts1 >= counts0 {
		return "1"
	}
	return "0"
}

func MostCommonValueCo2ScrubberRating(pointer int, data []string) string {
	counts0, counts1 := CountsOfZeroAndOne(pointer, data)
	if counts1 >= counts0 {
		return "0"
	}
	return "1"
}

func CountsOfZeroAndOne(pointer int, data []string) (int, int) {
	var counts1 int
	for _, num := range data {
		if string(num[pointer]) == "1" {
			counts1++
		}
	}
	return len(data) - counts1, counts1
}
