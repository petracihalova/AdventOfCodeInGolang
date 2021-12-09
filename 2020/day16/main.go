package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Puzzle input
	fields, _, nearbyTickets := GetPuzzleInput()

	//Puzzle 1
	ticketScanningErrorRate := 0
	for _, ticket := range nearbyTickets {
		for _, value := range ticket {
			validValue := ValueIsValidForAnyField(value, fields)
			if !validValue {
				ticketScanningErrorRate += value
			}
		}
	}
	fmt.Println("Puzzle 1 =", ticketScanningErrorRate)

	//Puzzle 2
	var nearbyTicketsValid [][]int
	fieldsPuzzle2, myTicket, nearbyTicketsPuzzle2 := GetPuzzleInput()

	//Get only valid nearby tickets
	for _, ticket := range nearbyTicketsPuzzle2 {
		validTicket := true
		for _, value := range ticket {
			validValue := ValueIsValidForAnyField(value, fieldsPuzzle2)
			if !validValue {
				validTicket = false
				break
			}
		}
		if validTicket {
			nearbyTicketsValid = append(nearbyTicketsValid, ticket)
		}
	}

	//Create map key=field and value=columns which match field conditions (from to)
	key := make(map[string][]int)
	for _, v := range fieldsPuzzle2 {
		key[v.name] = []int{}
	}

	for _, v := range fieldsPuzzle2 {
		for i := 0; i < len(nearbyTicketsValid[0]); i++ {
			ok := true
			for _, row := range nearbyTicketsValid {
				if row[i] < v.from1 || (row[i] > v.to1 && row[i] < v.from2) || row[i]> v.to2 {
					ok = false
				}
			}
			if ok {
				key[v.name] = append(key[v.name], i)
			}
		}
	}

	//Create final map where key=field name and value=column in nearby tickets
	finalKey := make(map[string]int)
	for len(finalKey) < len(fieldsPuzzle2) {
		for k, v := range key {
			if len(v) == 1 {
				finalKey[k] = v[0]
				key = DeleteIndexFromAllFields(key, v[0])
				break
			}
			if len(v) == 0 {
				delete(key, k)
				break
			}
		}
	}

	//Multiply fields in my ticket that start with the word "departure"
	result := 1
	for k, v := range finalKey {
		if strings.Contains(k, "departure") {
			result *= myTicket[v]
		}
	}
	fmt.Println("Puzzle 2 =", result)
}

func DeleteIndexFromAllFields(key map[string][]int, index int) map[string][]int {
	for k, v := range key {
		for i, num := range v {
			if num == index {
				v[i] = v[len(v)-1]
				key[k] = v[:len(v)-1]
			}
		}
	}
	return key
}

func ValueIsValidForAnyField(value int, fields []Field) bool {
	isValid := false
	for _, field := range fields {
		if field.from1 <= value && value <= field.to1 {
			isValid = true
			break
		}
		if field.from2 <= value && value <= field.to2 {
			isValid = true
			break
		}
	}
	return isValid
}

func GetPuzzleInput() ([]Field, []int, [][]int) {
	pwd, _ := os.Getwd()
	dataBytes, _ := os.ReadFile(pwd + "/2020/day16/input.txt")
	raw := strings.Split(string(dataBytes), "\n\n")

	fields := GetListFields(raw[0])
	myTicket := GetMyTicket(raw[1])
	nearbyTickets := GetNearbyTickets(raw[2])

	return fields, myTicket, nearbyTickets
}

func GetNearbyTickets(input string) [][]int {
	var tickets [][]int
	for _, row := range strings.Split(input, "\n")[1:] {
		var ticket []int
		for _, value := range strings.Split(row, ",") {
			num, _ := strconv.Atoi(value)
			ticket = append(ticket, num)
		}
		tickets = append(tickets, ticket)
	}
	return tickets
}

func GetMyTicket(input string) []int {
	var values []int
	input = strings.Split(input, "\n")[1]
	for _, v := range strings.Split(input, ",") {
		num, _ := strconv.Atoi(v)
		values = append(values, num)
	}
	return values
}

func GetListFields(input string) []Field {
	rawFields := strings.Split(input, "\n")
	var fields []Field
	for _, line := range rawFields {
		nameAndValues := strings.Split(line, ": ")
		fieldName := nameAndValues[0]
		values := strings.Split(nameAndValues[1], " or ")
		fromTo1 := strings.Split(values[0], "-")
		from1, _ := strconv.Atoi(fromTo1[0])
		to1, _ := strconv.Atoi(fromTo1[1])
		fromTo2 := strings.Split(values[1], "-")
		from2, _ := strconv.Atoi(fromTo2[0])
		to2, _ := strconv.Atoi(fromTo2[1])

		field := Field{name: fieldName, from1: from1, to1: to1, from2: from2, to2: to2}
		fields = append(fields, field)
	}
	return fields
}

type Field struct {
	name  string
	from1 int
	to1   int
	from2 int
	to2   int
}
