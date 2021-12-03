package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func PasswordIsValidPuzzle1 (psw string) bool {
	re := regexp.MustCompile(`(\d+)-(\d+) ([a-z]): ([a-z]+)`)
	reResult := re.FindStringSubmatch(psw)

	start, _ := strconv.Atoi(reResult[1])
	end, _ := strconv.Atoi(reResult[2])
	char, psw := reResult[3], reResult[4]

	count := strings.Count(psw, char)
	if count >= start && count <= end {
		return true
	}
	return false
}

func PasswordIsValidPuzzle2 (psw string) bool {
	re := regexp.MustCompile(`(\d+)-(\d+) ([a-z]): ([a-z]+)`)
	reResult := re.FindStringSubmatch(psw)

	first, _ := strconv.Atoi(reResult[1])
	second, _ := strconv.Atoi(reResult[2])
	char, psw := reResult[3], reResult[4]

	var containLetterCount int
	if string(psw[first-1]) == char{
		containLetterCount++
	}
	if string(psw[second-1]) == char{
		containLetterCount++
	}

	return containLetterCount == 1
}

func main() {
	//Puzzle input
	pwd, _ := os.Getwd()
	f, _ := os.Open(pwd + "/day02/input.txt")

	scanner := bufio.NewScanner(f)
	var puzzleInput []string
	for scanner.Scan() {
		puzzleInput = append(puzzleInput, scanner.Text())
	}

	//Puzzle 1
	var validCount1 int
	for i := range puzzleInput {
		if PasswordIsValidPuzzle1(puzzleInput[i]) {
			validCount1++
		}
	}
	fmt.Println("Puzzle 1 =", validCount1)

	//Puzzle 2
	var validCount2 int
	for i := range puzzleInput {
		if PasswordIsValidPuzzle2(puzzleInput[i]) {
			validCount2++
		}
	}
	fmt.Println("Puzzle 2 =", validCount2)

}
