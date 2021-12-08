package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Signal struct {
	pattern []string
	digits  []string
}

func main() {
	//Puzzle input
	input := GetPuzzleOneInput()

	//Puzzle 1
	var result int
	for _, v := range input {
		if len(v) == 2 || len(v) == 3 || len(v) == 4 || len(v) == 7 {
			result++
		}
	}
	fmt.Println("Puzzle 1 =", result)

	//Puzzle 2
	signals := GetPuzzleTwoInput()
	result = 0
	for _, s := range signals {
		result += GenerateOutput(s)
	}
	fmt.Println("Puzzle 2 =", result)
}

func GenerateOutput(s Signal) int {
	var digits []string
	for i := 0; i < 10; i++ {
		digits = append(digits, "")
	}
	// Numbers 1, 4, 7, 8
	for _, v := range s.pattern {
		if len(v) == 2 {
			digits[1] = v
		}
		if len(v) == 3 {
			digits[7] = v
		}
		if len(v) == 4 {
			digits[4] = v
		}
		if len(v) == 7 {
			digits[8] = v
		}
	}
	// Numbers 0, 6, 9
	for _, v := range s.pattern {
		if len(v) != 6 {
			continue
		}
		contains7, _ := FirstInSecond(digits[7], v)
		contains4, _ := FirstInSecond(digits[4], v)
		switch {
		case contains7 && contains4:
			digits[9] = v
		case contains7:
			digits[0] = v
		default:
			digits[6] = v
		}
	}

	// Numbers 2, 3, 5
	for _, v := range s.pattern {
		if len(v) != 5 {
			continue
		}
		contains7, _ := FirstInSecond(digits[7], v)
		_, diff := FirstInSecond(digits[6], v)

		switch {
		case contains7:
			digits[3] = v
		case diff == 1:
			digits[5] = v
		default:
			digits[2] = v
		}
	}

	var result string
	for _, digit := range s.digits {
		result += strconv.Itoa(indexOf(digit, digits))

	}
	num, _ := strconv.Atoi(result)

	return num
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found
}

func FirstInSecond(first string, second string) (bool, int) {
	var diff int
	for _, v := range first {
		if !strings.Contains(second, string(v)) {
			diff++
		}
	}
	if diff == 0 {
		return true, 0
	}
	return false, diff
}

func GetPuzzleOneInput() []string {
	pwd, _ := os.Getwd()
	f, _ := os.Open(pwd + "/2021/day08/input.txt")

	scanner := bufio.NewScanner(f)
	var puzzleInput []string
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " | ")
		digits := strings.Split(line[1], " ")
		for _, d := range digits {
			puzzleInput = append(puzzleInput, d)
		}

	}
	return puzzleInput
}

func GetPuzzleTwoInput() []Signal {
	pwd, _ := os.Getwd()
	f, _ := os.Open(pwd + "/2021/day08/input.txt")

	scanner := bufio.NewScanner(f)
	var puzzleInput []Signal
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " | ")
		pattern := strings.Split(line[0], " ")
		SortAllItemsAlphabetically(pattern)
		digits := strings.Split(line[1], " ")
		SortAllItemsAlphabetically(digits)
		puzzleInput = append(puzzleInput, Signal{pattern, digits})
	}
	return puzzleInput
}

func SortAllItemsAlphabetically(pattern []string) {
	for i, v := range pattern {
		s := []rune(v)
		sort.Sort(sortRunes(s))
		pattern[i] = string(s)
	}
}

//source: https://mariadesouza.com/2018/01/01/string-manipulation-in-go/
type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Len() int {
	return len(s)
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
