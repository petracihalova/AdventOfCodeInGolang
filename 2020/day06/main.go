package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//Puzzle input
	pwd, _ := os.Getwd()
	dataBytes, _ := os.ReadFile(pwd + "/day06/input.txt")

	data := string(dataBytes)
	groups := CreateGroupList(data)

	//Puzzle 1
	var counterPuzzle1 int
	for _, group := range groups {
		counterPuzzle1 += YesAnyoneFromGroup(group)
	}
	fmt.Println("Puzzle 1 = ", counterPuzzle1)

	//Puzzle 2
	var counterPuzzle2 int
	for _, group := range groups {
		counterPuzzle2 += YesEveryoneFromGroup(group)
	}
	fmt.Println("Puzzle 2 = ", counterPuzzle2)
}

func YesEveryoneFromGroup(group string) int {
	uniqYesQuestions := GetUniqQuestionsFromGroup(group)
	groupList := strings.Split(group, "\n")

	var counter int
	for _, item := range uniqYesQuestions {
		if len(groupList) == strings.Count(group, string(item)) {
			counter++
		}
	}
	return counter
}

func YesAnyoneFromGroup(group string) int {
	uniqYesQuestions := GetUniqQuestionsFromGroup(group)
	return len(uniqYesQuestions)
}

func GetUniqQuestionsFromGroup(group string) string {
	var uniqYesQuestions string
	group = strings.Replace(group, " ", "", -1)
	group = strings.Replace(group, "\n", "", -1)
	for _, char := range group {
		if !strings.Contains(uniqYesQuestions, string(char)) {
			uniqYesQuestions += string(char)
		}
	}
	return uniqYesQuestions
}

func CreateGroupList(data string) []string {
	groups := strings.Split(data, "\n\n")
	return groups
}
