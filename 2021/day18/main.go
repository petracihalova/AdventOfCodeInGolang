package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	//Puzzle 1
	input := GetPuzzleInput()
	number := CalculateResult(input)
	number = CalculateMagnitude(number)

	fmt.Println("Puzzle 1 =", number)

	//Puzzle 2
	var maxMagnitude int
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			if i == j {
				continue
			}
			twoNumbers := []string{input[i], input[j]}
			number = CalculateResult(twoNumbers)
			result, _ := strconv.Atoi(CalculateMagnitude(number))
			if result > maxMagnitude {
				maxMagnitude = result
			}

			twoNumbers = []string{input[j], input[i]}
			number = CalculateResult(twoNumbers)
			result, _ = strconv.Atoi(CalculateMagnitude(number))
			if result > maxMagnitude {
				maxMagnitude = result
			}
		}
	}
	fmt.Println("Puzzle 2 =", maxMagnitude)
}

func GetPuzzleInput() []string {
	pwd, _ := os.Getwd()
	f, _ := os.Open(pwd + "/2021/day18/input.txt")

	scanner := bufio.NewScanner(f)
	var snailfishNumbers []string

	for scanner.Scan() {
		snailfishNumbers = append(snailfishNumbers, scanner.Text())
	}
	return snailfishNumbers
}

func AdditionOfSnaifishNumbers(first string, second string) string {
	return "[" + first + "," + second + "]"
}

func Explode(number string) string {
	var level int
	for pointer, char := range number {
		switch string(char) {
		case "[":
			level += 1
			continue
		case "]":
			level -= 1
			continue
		case ",":
			continue
		}
		if level != 5 {
			continue
		}
		var end int
		for i := pointer; i < len(number); i++ {
			if string(number[i]) == "]" {
				end = i - 1
				break
			}
		}
		pair := strings.Split(number[pointer:end+1], ",")
		num1, _ := strconv.Atoi(pair[0])
		num2, _ := strconv.Atoi(pair[1])
		//originPair := fmt.Sprintf("[%d,%d]", num1, num2)

		var diff int
		for i := pointer - 1; i >= 0; i-- {
			if strings.Contains("[],", string(number[i])) {
				continue
			}
			originNumber := string(number[i])
			if !strings.Contains("[],", string(number[i-1])) {
				originNumber = number[i-1:i+1]
			}
			num, _ := strconv.Atoi(originNumber)
			newNumber := strconv.Itoa(num + num1)
			diff = len(newNumber) - len(originNumber)
			pointer += diff
			number = number[:i+1-len(originNumber)] + newNumber + number[i+1:]
			break
		}

		for j := end + 1 + diff; j < len(number); j++ {
			if strings.Contains("[],", string(number[j])) {
				continue
			}
			originNumber := string(number[j])
			if !strings.Contains("[],", string(number[j+1])) {
				originNumber = number[j:j+2]
			}
			num, _ := strconv.Atoi(originNumber)
			newNumber := strconv.Itoa(num + num2)
			number = number[:j] + newNumber + number[j+len(originNumber):]
			break
		}
		number = number[:pointer - 1] + "0" + number[end+2+diff:]
		break
	}
	return number
}

func Split(number string) string {
	for pointer, char := range number {
		switch string(char) {
		case "[", "]", ",":
			continue
		}
		num, err := strconv.Atoi(number[pointer: pointer+2])
		if err != nil {
			continue
		}
		first := num / 2
		second := num - first
		newPair := fmt.Sprintf("[%d,%d]", first, second)
		number = number[:pointer] + newPair + number[pointer+2:]
		break
	}
	return number
}

func CalculateResult(inputNumbers []string) string {
	number := inputNumbers[0]
	for i := 1; i < len(inputNumbers); i++ {
		number = AdditionOfSnaifishNumbers(number, inputNumbers[i])
		for {
			newNumber := Explode(number)
			if newNumber != number {
				number = newNumber
				continue
			}
			newNumber = Split(number)
			if newNumber != number {
				number = newNumber
				continue
			}
			break
		}
	}
	return number
}

func CalculateMagnitude(number string) string {
	r, _ := regexp.Compile("[0-9]+,[0-9]+")
	for r.FindString(number) != "" {
		pair := strings.Split(r.FindString(number), ",")
		first, _ := strconv.Atoi(pair[0])
		second, _ := strconv.Atoi(pair[1])
		sum := strconv.Itoa(first * 3 + second * 2)
		originPair := "[" + pair[0] + "," + pair[1] + "]"
		number = strings.Replace(number, originPair, sum, -1)
	}
	return number
}