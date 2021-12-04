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
	f, _ := os.Open(pwd + "/2020/day07/input.txt")

	scanner := bufio.NewScanner(f)
	var puzzleInput []string
	for scanner.Scan() {
		puzzleInput = append(puzzleInput, scanner.Text())
	}
	rules := CreateListOfRules(puzzleInput)

	//Puzzle 1
	var bags []string
	bags = HowManyBagColorsContain(rules, "shiny gold", bags)
	fmt.Println("Puzzle 1 =", len(bags))

	//Puzzle 2
	var bagsCount int
	bagsCount = HowManyBagsContain(rules, "shiny gold", bagsCount)
	fmt.Println("Puzzle 2 =", bagsCount - 1) // bagsCont without 1x shiny gold bag

}

func HowManyBagsContain(rules [][]string, bag string, bagsCount int) int {
	for _, rule := range rules {
		if rule[0] == bag && rule[2] == "no other" {
			return 1
		}
		if rule[0] == bag {
			num, _ := strconv.Atoi(rule[1])
			bagsCount += num * HowManyBagsContain(rules, rule[2], 0)
		}
	}
	return bagsCount + 1 // +1 bag itself
}

func HowManyBagColorsContain(rules [][]string, bag string, bags []string) []string {
	for _, rule := range rules {
		if rule[2] == bag {
			bags = AddUniqBag(bags, rule[0])
			bags = HowManyBagColorsContain(rules, rule[0], bags)
		}
	}
	return bags
}

func AddUniqBag(bags []string, bag string) []string {
	addBag := true
	for _, item := range bags {
		if item == bag {
			addBag = false
		}
	}
	if addBag {
		bags = append(bags, bag)
	}
	return bags
}

func CreateListOfRules(input []string) [][]string {
	var rules [][]string
	for _, line := range input {
		for _, rule := range CreateRule(line) {
			rules = append(rules, rule)
		}
	}
	return rules
}

func CreateRule(line string) [][]string {
	var rules [][]string

	result := strings.Split(line, " bags contain ")
	from := result[0]
	to := strings.Split(result[1], ", ")
	for _, bagTo := range to {
		bagTo = strings.Split(bagTo, " bag")[0]
		var count string
		if bagTo == "no other" {
			count = "0"
		} else {
			count = strings.Split(bagTo, " ")[0]
			bagTo = fmt.Sprintf("%s %s", strings.Split(bagTo, " ")[1], strings.Split(bagTo, " ")[2])
		}

		rule := []string {
			from,
			count,
			bagTo,
		}
		rules = append(rules, rule)
	}
	return rules

}