package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//Puzzle input
	pwd, _ := os.Getwd()
	f, _ := os.Open(pwd + "/2020/day13/input.txt")

	scanner := bufio.NewScanner(f)
	var puzzleInput []string
	for scanner.Scan() {
		puzzleInput = append(puzzleInput, scanner.Text())
	}

	//Puzzle 1 input data preparation
	timestamp, _ := strconv.Atoi(puzzleInput[0])
	var buses []int
	for _, item := range strings.Split(puzzleInput[1], ",") {
		if item == "x" {
			continue
		}
		bus, _ := strconv.Atoi(item)
		buses = append(buses, bus)
	}
	sort.Ints(buses)


	//Puzzle 1
	bus, departTimestamp := FindBus(timestamp, buses)
	//fmt.Println("Bus:", bus, "\nDepart time:", departTimestamp,"\nOrigin timestamp", timestamp)
	fmt.Println("Puzzle 1 =", bus * (departTimestamp - timestamp))


	//Puzzle 2 input data preparation
	var buses2 []int
	for _, item := range strings.Split(puzzleInput[1], ",") {
		bus2, _ := strconv.Atoi(item)
		buses2 = append(buses2, bus2)
	}

	//Puzzle 2
	var n []*big.Int
	var a []*big.Int

	for index, number := range buses2 {
		if number == 0 {
			continue
		}
		n = append(n, big.NewInt(int64(number)))
		a = append(a, big.NewInt(int64(len(buses2) - 1 - index)))
	}

	result, _ := crt(a, n)
	num, _ := strconv.Atoi(fmt.Sprintf("%v", result))

	fmt.Println("Puzzle 2 =", num - len(buses2) + 1)
}

func FindBus(timestamp int, buses []int) (int, int){
	for i := timestamp; i >= timestamp; i++ {
		for _, bus := range buses {
			if i % bus == 0 {
				return bus, i
			}
		}
	}
	return 0, 0
}

var one = big.NewInt(1)

//crt computes Chinese remainder theorem
//Source of function here: https://rosettacode.org/wiki/Chinese_remainder_theorem
func crt(a, n []*big.Int) (*big.Int, error) {
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		if z.Cmp(one) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p), nil
}