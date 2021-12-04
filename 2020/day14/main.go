package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Puzzle input
	pwd, _ := os.Getwd()
	f, _ := os.Open(pwd + "/2020/day14/input.txt")

	scanner := bufio.NewScanner(f)
	var puzzleInput []string
	for scanner.Scan() {
		puzzleInput = append(puzzleInput, scanner.Text())
	}

	//Puzzle 1
	memory := make(map[int]int)
	var mask string

	for _, record := range puzzleInput {
		if strings.Contains(record, "mask") {
			mask = record[7:]
			continue
		}
		address, value := ParseInput(record)
		valueAsBinary:= fmt.Sprintf("%036s", strconv.FormatInt(int64(value), 2))
		valueAsBinary = ApplyMaskPuzzle1(valueAsBinary, mask)
		result, _ := strconv.ParseInt(valueAsBinary, 2, 64)
		memory[address] = int(result)
	}
	var sum int
	for _, v := range memory {
		sum += v
	}
	fmt.Println("Puzzle 1 =", sum)

	//Puzzle 2
	memoryPuzzle2 := make(map[int]int)

	for _, record := range puzzleInput {
		if strings.Contains(record, "mask") {
			mask = record[7:]
			continue
		}
		address, originValue := ParseInput(record)
		addressAsBinary := fmt.Sprintf("%036s", strconv.FormatInt(int64(address), 2))
		result := ApplyMaskPuzzle2(addressAsBinary, mask)
		listOfResults := CreateCombinations(result)

		// Save Results into memory
		for _, valueAsBinary := range listOfResults {
			num, _ := strconv.ParseInt(valueAsBinary, 2, 64)
			memoryPuzzle2[int(num)] = originValue
		}
	}

	sum = 0
	for _, v := range memoryPuzzle2 {
		sum += v
	}
	fmt.Println("Puzzle 2 =", sum)

}

func CreateCombinations(input string) []string {
	k := strings.Count(input, "X")

	//Create combinations of 01
	inputNumbers := "01"
	n := len(inputNumbers)
	var listCombinations []string
	combinations := GetListOfCombinations(listCombinations, inputNumbers, "", n, k)

	//Create list of mask according combinations
	result := GetListOfResults(combinations, input)

	return result
}

func GetListOfResults(combinations []string, input string) []string {
	var result []string
	for _, combination := range combinations {
		r := input
		for _, c := range combination {
			r = strings.Replace(r, "X", string(c), 1)
		}
		result = append(result, r)
	}
	return result
}

func GetListOfCombinations(combinations []string, input string, prefix string, n int, k int) []string {

	if k == 0 {
		combinations = append(combinations, prefix)
		return combinations
	}
	for i := 0; i < n; i++ {
		newPrefix := prefix + string(input[i])
		combinations = GetListOfCombinations(combinations, input, newPrefix, n, k - 1)
	}
	return combinations
}

func ApplyMaskPuzzle2(value string, mask string) string {
	result := ""
	for i, bit := range mask {
		switch string(bit) {
		case "0":
			result += string(value[i])
		case "1", "X":
			result += string(bit)
		}
	}
	return result
}

func ApplyMaskPuzzle1(value string, mask string) string {
	result := ""
	for i, v := range value {
		if string(mask[i]) == "X" {
			result += string(v)
		} else {
			result += string(mask[i])
		}
	}
	return result
}

func ParseInput(line string) (int, int) {
	line = strings.Replace(line, "mem[", "", 1)
	line = strings.Replace(line, "] =", "", 1)
	result := strings.Split(line, " ")
	address, _ := strconv.Atoi(result[0])
	value, _ := strconv.Atoi(result[1])
	return address, value
}
