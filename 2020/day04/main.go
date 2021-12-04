package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	//Puzzle input
	pwd, _ := os.Getwd()
	dataBytes, _ := os.ReadFile(pwd + "/2020/day04/input.txt")

	data := string(dataBytes)
	passports := CreatePassportList(data)

	//Puzzle 1
	var passportValidCounter1 int
	for _, passport := range passports {
		if PassportIsValidPuzzle1(passport) {
			passportValidCounter1++
		}
	}

	fmt.Println("Puzzle 1 = ", passportValidCounter1)

	//Puzzle 2
	var passportValidCounter2 int
	for _, passport := range passports {
		if PassportIsValidPuzzle2(passport) {
			passportValidCounter2++
		}
	}

	fmt.Println("Puzzle 2 = ", passportValidCounter2)
}

func CreatePassportList(data string) []map[string]string {
	var passports []map[string]string
	passportsListRaw := strings.Split(data, "\n\n")

	for _, passportRaw := range passportsListRaw {
		passportRaw = strings.Replace(passportRaw, "\n", " ", -1)
		passportItem := strings.Split(passportRaw, " ")
		passport := map[string]string{}

		for _, i := range passportItem {
			keyValuePair := strings.Split(i, ":")
			key, value := keyValuePair[0], keyValuePair[1]
			passport[key] = value
		}
		passports = append(passports, passport)
	}
	return passports
}

func PassportIsValidPuzzle1(passport map[string]string) bool {
	if len(passport) == 8 {
		return true
	}
	if len(passport) == 7 {
		if _, ok := passport["cid"]; !ok {
			return true
		}
	}
	return false
}

func PassportIsValidPuzzle2(passport map[string]string) bool {
	if !PassportIsValidPuzzle1(passport) {
		return false
	}

	birthYear, _ := strconv.Atoi(passport["byr"])
	if birthYear < 1920 || birthYear > 2002 {
		return false
	}

	issueYear, _ := strconv.Atoi(passport["iyr"])
	if issueYear < 2010 || issueYear > 2020 {
		return false
	}

	expirationYear, _ := strconv.Atoi(passport["eyr"])
	if expirationYear < 2020 || expirationYear > 2030 {
		return false
	}

	h, unit := passport["hgt"][:len(passport["hgt"])-2], passport["hgt"][len(passport["hgt"])-2:]
	height, _ := strconv.Atoi(h)
	switch unit {
	case "cm":
		if height < 150 || height > 193 {
			return false
		}
	case "in":
		if height < 59 || height > 76 {
			return false
		}
	default:
		return false
	}

	re, _ := regexp.MatchString("#[0-9a-f]{6}", passport["hcl"])
	if re == false {
		return false
	}

	eyeColors := []string { "amb", "blu", "brn", "gry", "grn", "hzl", "oth" }
	var result bool
	for _, value := range eyeColors {
		if value == passport["ecl"] {
			result = true
			break
		}
	}
	if !result {
		return false
	}

	if len(passport["pid"]) != 9 {
		return false
	}

	_, err := strconv.Atoi(passport["pid"])
	if err != nil {
		return false
	}

	return true
}