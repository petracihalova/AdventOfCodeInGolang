package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	command string
	value int32
}

type Submarine struct {
	horizontalPosition int32
	depth int32
	aim int32
}

func main() {
	//Puzzle input
	pwd, _ := os.Getwd()
	f, _ := os.Open(pwd + "/day02/input.txt")

	scanner := bufio.NewScanner(f)
	var puzzleInput []Command
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		direction := line[0]
		value, _ := strconv.Atoi(line[1])
		command := Command{command: direction, value: int32(value)}
		puzzleInput = append(puzzleInput, command)
	}

	//Puzzle 1
	submarine := Submarine{}
	for _, v := range puzzleInput {
		switch v.command {
		case "forward":
			submarine.horizontalPosition += v.value
		case "down":
			submarine.depth += v.value
		case "up":
			submarine.depth -= v.value
		default:
			panic("unexpected command")
		}
	}

	result := submarine.horizontalPosition * submarine.depth
	fmt.Println("Puzzle 1 = ", result)

	//Puzzle 2
	submarine2 := Submarine{}
	for _, v := range puzzleInput {
		switch v.command {
		case "forward":
			submarine2.horizontalPosition += v.value
			submarine2.depth += submarine2.aim * v.value
		case "down":
			submarine2.aim += v.value
		case "up":
			submarine2.aim -= v.value
		default:
			panic("unexpected command")
		}
	}

	result2 := submarine2.horizontalPosition * submarine2.depth
	fmt.Println("Puzzle 1 = ", result2)


}
