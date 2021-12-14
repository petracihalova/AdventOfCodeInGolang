package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	//Puzzle 1
	formula, rules := GetPuzzleInput()
	for formula.Step < 10 {
		formula.NextStep(rules)
	}
	result := CountResult(formula.CountsElements)
	fmt.Println("Puzzle 1 =", result)

	//Puzzle 2
	formula, rules = GetPuzzleInput()
	for formula.Step < 40 {
		formula.NextStep(rules)
	}
	result = CountResult(formula.CountsElements)
	fmt.Println("Puzzle 1 =", result)

}

func CountResult(elements map[string]int) int {
	var numbers []int
	for _, v := range elements{
		numbers = append(numbers, v)
	}
	sort.Ints(numbers)
	return numbers[len(numbers)-1] - numbers[0]
}

type Formula struct {
	Template 		 string
	Step	 		 int
	CountsElements 	 map[string]int
	PairsElements 	 map[string]int
}

func (f *Formula) NextStep(rules map[string]string) {
	f.Step += 1
	temp := make(map[string]int)
	for key, value := range f.PairsElements {
		newPair1 := string(key[0]) + rules[key]
		newPair2 := rules[key] + string(key[1])
		if _, ok := temp[newPair1]; ok {
			temp[newPair1] += value
		} else {
			temp[newPair1] = value
		}
		if _, ok := temp[newPair2]; ok {
			temp[newPair2] += value
		} else {
			temp[newPair2] = value
		}
		if _, ok := f.CountsElements[rules[key]]; ok {
			f.CountsElements[rules[key]] += value
		} else {
			f.CountsElements[rules[key]] = value
		}

	}
	f.PairsElements = temp
}

func GetPuzzleInput() (Formula, map[string]string) {
	pwd, _ := os.Getwd()
	dataBytes, _ := os.ReadFile(pwd + "/2021/day14/input.txt")
	raw := strings.Split(string(dataBytes), "\n\n")

	rules := make(map[string]string)
	for _, item := range strings.Split(raw[1], "\n") {
		line := strings.Split(item, " -> ")
		key := line[0]
		value := line[1]
		rules[key] = value
	}

	template := raw[0]
	counts := make(map[string]int)
	for _, v := range template {
		if _, ok := counts[string(v)]; ok {
			counts[string(v)] += 1
		} else {
			counts[string(v)] = 1
		}
	}
	pairs := make(map[string]int)
	for i := 1; i < len(template); i++ {
		pair := template[i-1: i+1]
		if _, ok := pairs[pair]; ok {
			pairs[pair] += 1
		} else {
			pairs[pair] = 1
		}
	}
	f := Formula{
		Template:       template,
		Step:           0,
		CountsElements: counts,
		PairsElements:  pairs,
	}

	return f, rules
}