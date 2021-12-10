package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	//Puzzle input
	input := GetPuzzleInput()


	//Puzzle 1
	scorePoints := map[string]int {
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	score := 0
	for _, line := range input {
		index := IncompleteBracketIndex(line)
		if index == -1 {
			continue
		}
		score += scorePoints[string(line[index])]
	}
	fmt.Println("Puzzle 1 =", score)


	//Puzzle 2
	var scores []int
	for _, line := range input {
		brackets := GetMissingBrackets(line)
		if brackets == "" {
			continue
		}
		bracketsScore := CalculateScore(brackets)
		scores = append(scores, bracketsScore)
	}
	sort.Ints(scores)
	index := len(scores) / 2


	fmt.Println("Puzzle 2 =", scores[index])

}

func CalculateScore(brackets string) int {
	points := map[string]int {
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}
	var score int
	for _, b := range brackets {
		score = score * 5 + points[string(b)]
	}
	return score
}

func GetMissingBrackets(line string) string {
	openBrackets := "([{<"
	closeBrackets := ")]}>"
	temp := ""
	for _, b := range line {
		if strings.Index(openBrackets, string(b)) != -1 {
			temp += string(b)
		} else {
			if strings.Index(closeBrackets, string(b)) == strings.Index(openBrackets, string(temp[len(temp)-1])) {
				temp = temp[:len(temp) - 1]
			} else {
				return ""
			}
		}
	}
	var missingBrackets string
	for i:= len(temp) - 1; i >= 0; i-- {
		index := strings.Index(openBrackets, string(temp[i]))
		missingBrackets += string(closeBrackets[index])
	}
	return missingBrackets
}

func IncompleteBracketIndex(line string) int {
	openBrackets := "([{<"
	closeBrackets := ")]}>"
	temp := ""
	for i, b := range line {
		if strings.Index(openBrackets, string(b)) != -1 {
			temp += string(b)
		} else {
			if strings.Index(closeBrackets, string(b)) == strings.Index(openBrackets, string(temp[len(temp)-1])) {
				temp = temp[:len(temp) - 1]
			} else {
				return i
			}
		}
	}
	return -1
}

func GetPuzzleInput() []string {
	pwd, _ := os.Getwd()
	f, _ := os.Open(pwd + "/2021/day10/input.txt")

	scanner := bufio.NewScanner(f)
	var puzzleInput []string
	for scanner.Scan() {
		puzzleInput = append(puzzleInput, scanner.Text())
	}
	return puzzleInput
}
