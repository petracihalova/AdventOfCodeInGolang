package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//Puzzle 1
	caves := GetPuzzleInput()
	var result int
	result = GetCountOfPathsPuzzle1(caves, "start", result, "")

	fmt.Println("Puzzle 1 =", result)

	//Puzzle 2
	caves = GetPuzzleInput()
	result = 0
	result = GetCountOfPathsPuzzle2(caves, "start", result, "")

	fmt.Println("Puzzle 2 =", result)

}

func GetCountOfPathsPuzzle1(caves map[string][]string, pointer string, result int, path string) int {
	path += pointer + " "
	for _, direction := range caves[pointer] {
		if direction == "end" {
			result += 1
			continue
		}
		if IsLower(direction) && strings.Contains(path, direction) {
			continue
		}
		result = GetCountOfPathsPuzzle1(caves, direction, result, path)
	}
	return result
}

func GetCountOfPathsPuzzle2(caves map[string][]string, pointer string, result int, path string) int {
	path += pointer + " "
	for _, direction := range caves[pointer] {
		if direction == "end" {
			result += 1
			continue
		}
		if IsLower(direction) {
			counts := make(map[string]int)
			for _, cave := range strings.Split(strings.Trim(path, " "), " ") {
				if !IsLower(cave) || cave == "start" {
					continue
				}
				if _, ok := counts[cave]; ok {
					counts[cave] += 1
				} else {
					counts[cave] = 1
				}
			}
			anyCaveTwice := false
			for _, v := range counts {
				if v == 2 {
					anyCaveTwice = true
				}
			}
			if _, ok := counts[direction]; ok && anyCaveTwice{
				continue
			}
		}
		result = GetCountOfPathsPuzzle2(caves, direction, result, path)
	}
	return result
}

func GetPuzzleInput() map[string][]string {
	pwd, _ := os.Getwd()
	f, _ := os.Open(pwd + "/2021/day12/input.txt")

	scanner := bufio.NewScanner(f)
	var puzzleInput [][]string
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "-")
		puzzleInput = append(puzzleInput, line)
		puzzleInput = append(puzzleInput, []string{line[1], line[0]})
	}
	caves := make(map[string][]string)
	for _, item := range puzzleInput {
		v1 := item[0]
		v2 := item[1]
		if v2 == "start" || v1 == "end" {
			continue
		}
		if _, ok := caves[v1]; ok {
			caves[v1] = append(caves[v1], v2)
		} else {
			caves[v1] = []string{v2}
		}
	}
	return caves
}

func IsLower(s string) bool {
	return s == strings.ToLower(s)
}