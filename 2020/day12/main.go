package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	//Puzzle input
	pwd, _ := os.Getwd()
	f, _ := os.Open(pwd + "/2020/day12/input.txt")

	scanner := bufio.NewScanner(f)
	var instructions []string
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	//Puzzle 1
	ship := Ship{x: 0, y: 0, direction: 90}
	for _, instruction := range instructions {
		ship.MoveShipPuzzle1(instruction)
	}

	manhattanDistance := math.Abs(float64(ship.x)) + math.Abs(float64(ship.y))
	fmt.Println("Puzzle 1 = ", manhattanDistance)

	//Puzzle 2
	w := Waypoint{x: 10, y: 1}
	ship2 := Ship{x: 0, y: 0, direction: 90, waypoint: w}
	for _, instruction := range instructions {
		ship2.MoveShipPuzzle2(instruction)
	}
	manhattanDistance = math.Abs(float64(ship2.x)) + math.Abs(float64(ship2.y))
	fmt.Println("Puzzle 2 = ", manhattanDistance)
}

func (ship *Ship) MoveShipPuzzle1 (instruction string) {
	action := string(instruction[0])
	if action == "F" {
		action = ReturnStringDirection(ship.direction)
	}
	value, _ := strconv.Atoi(instruction[1:])
	switch action {
	case "N":
		ship.y += value
	case "S":
		ship.y -= value
	case "E":
		ship.x += value
	case "W":
		ship.x -= value
	case "L":
		ship.direction = (ship.direction - value) % 360
		for ship.direction < 0 {
			ship.direction += 360
		}
	case "R":
		ship.direction = (ship.direction + value) % 360
	}
}

func (ship *Ship) MoveShipPuzzle2(instruction string) {
	action := string(instruction[0])
	value, _ := strconv.Atoi(instruction[1:])
	switch action {
	case "F":
		ship.x += value * ship.waypoint.x
		ship.y += value * ship.waypoint.y
	case "N":
		ship.waypoint.y += value
	case "S":
		ship.waypoint.y -= value
	case "E":
		ship.waypoint.x += value
	case "W":
		ship.waypoint.x -= value
	case "L", "R":
		for i := 0; i < (value / 90); i++ {
			xNew := ship.waypoint.y
			yNew := ship.waypoint.x
			if action == "L" {
				xNew *= -1
			} else if action == "R"{
				yNew *= -1
			}
			ship.waypoint.x = xNew
			ship.waypoint.y = yNew
		}
	default:
		panic("unexpected input")
	}
}

func ReturnStringDirection(direction int) string {
	switch direction {
	case 0:
		return "N"
	case 90:
		return "E"
	case 180:
		return "S"
	case 270:
		return "W"
	default:
		panic("wrong direction")
	}
}

type Ship struct {
	x int
	y int
	direction int // 0 = north, 90 = east, 180 = south, 270 = west
	waypoint Waypoint
}

type Waypoint struct {
	x int
	y int
}
