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
	f, _ := os.Open(pwd + "/day08/input.txt")

	scanner := bufio.NewScanner(f)
	var puzzleInput []string
	for scanner.Scan() {
		puzzleInput = append(puzzleInput, scanner.Text())
	}
	instructions := CreateListInstructions(puzzleInput)

	//Puzzle 1
	accumulator, status := RunProgram(instructions)
	switch status {
	case "loop":
		fmt.Println("Puzzle 1 =", accumulator, status)
	case "end":
		panic("status \"loop\" expected, got \"end\"")
	default:
		panic("something wrong")
	}

	//Puzzle 2
	for i := 0; i < len(instructions); i++ {
		accumulator, status = RunProgram(instructions)
		switch status {
		case "loop":
			if instructions[i].operation == "acc" {
				continue
			} else {
				instructions = CreateListInstructions(puzzleInput)
				ChangeInstructions(instructions, i)
			}
		case "end":
			fmt.Println("Puzzle 2 =", accumulator, status)
		default:
			panic("something wrong")
		}

		if status == "end" {
			break
		}
	}

}

func ChangeInstructions(instructions []Instruction, i int) {
	toChange := instructions[i]
	switch toChange.operation {
	case "jmp":
		instructions[i].operation = "nop"
	case "nop":
		instructions[i].operation = "jmp"
	}
}

func RunProgram(instructions []Instruction) (int, string) {
	var pointer int
	var status string
	var accumulator int
	var listIndexes []int
	for {
		if pointer == len(instructions) {
			status = "end"
			break
		}
		i := instructions[pointer]
		if IndexInList(listIndexes, pointer) {
			status = "loop"
			break
		} else {
			listIndexes = append(listIndexes, pointer)
		}

		switch i.operation {
		case "nop":
			pointer++
		case "acc":
			accumulator += i.argument
			pointer++
		case "jmp":
			pointer += i.argument
		}
	}
	return accumulator, status
}

func IndexInList(listIndexes []int, index int) bool {
	for _, i := range listIndexes {
		if i == index {
			return true
		}
	}
	return false
}

type Instruction struct {
	operation string
	argument int
}

func CreateListInstructions(input []string) []Instruction {
	var result []Instruction
	for _, value := range input {
		splitValue := strings.Split(value, " ")
		o := splitValue[0]
		a, _ := strconv.Atoi(splitValue[1])
		instruction := Instruction{
			operation: o,
			argument: a,
		}
		result = append(result, instruction)
	}
	return result
}