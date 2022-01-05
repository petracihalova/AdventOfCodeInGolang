package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Puzzle 1 + Puzzle 2
	imageEnhAlgorithm, image := GetInputData()
	rounds := 50

	// for test data is needed to use: backgrounds := []string{".", "."}
	backgrounds := []string{".", "#"}
	for i := 0; i < rounds; i++ {
		image = ConvertImage(image, imageEnhAlgorithm, backgrounds[i % 2])
		if i == 1 {
			fmt.Println("Puzzle 1 =", len(image))
		}
	}
	fmt.Println("Puzzle 2 =", len(image))

}
//PrintGrid only helper for debugging. Not needed for final solution.
func PrintGrid(image map[string]bool) {
	size := GetSizeImage(image)
	for row := size["min i"] - 10; row < size["max i"] + 11; row++ {
		for col := size["min j"] - 10; col < size["max j"] + 11; col++ {
			coor := fmt.Sprintf("%d,%d", row, col)
			if _, ok := image[coor]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func GetSizeImage(image map[string]bool) map[string]int {
	size := map[string]int{
		"min i": 100000,
		"max i": 0,
		"min j": 100000,
		"max j": 0,
	}
	for k, _ := range image {
		coor := strings.Split(k, ",")
		i, _ := strconv.Atoi(coor[0])
		j, _ := strconv.Atoi(coor[1])
		if i < size["min i"] {
			size["min i"] = i
		}
		if i > size["max i"] {
			size["max i"] = i
		}
		if j < size["min j"] {
			size["min j"] = j
		}
		if j > size["max j"] {
			size["max j"] = j
		}
	}
	return size
}

func ConvertImage(image map[string]bool, algorithm map[string]bool, background string) map[string]bool {
	size := GetSizeImage(image)
	newImage := make(map[string]bool)
	for row := size["min i"] - 2; row < size["max i"] + 3; row++ {
		for col := size["min j"] - 2; col < size["max j"] + 3; col++ {
			binaryNumber := GetBinaryNumberForPixel(row, col, image, size, background)
			if _, ok := algorithm[binaryNumber]; ok {
				coor := fmt.Sprintf("%d,%d", row, col)
				newImage[coor] = true
			}
		}
	}
	return newImage
}

func GetBinaryNumberForPixel(row int, col int, image map[string]bool, size map[string]int, background string) string {
	var binary string
	for i := row - 1; i < row + 2; i++ {
		for j := col - 1; j < col + 2; j++ {
			coor := fmt.Sprintf("%d,%d", i, j)
			if _, ok := image[coor]; ok {
				binary += "1"
			} else if i < size["min i"] || i > size["max i"] || j < size["min j"] || j > size["max j"] {
				if background == "." {
					binary += "0"
				} else if background == "#" {
					binary += "1"
				}
			} else {
				binary += "0"
			}
		}
	}
	return binary
}

func GetInputData() (map[string]bool, map[string]bool) {
	pwd, _ := os.Getwd()
	dataBytes, _ := os.ReadFile(pwd + "/2021/day20/input.txt")

	imageEnhAlgorithm := make(map[string]bool)
	image := make(map[string]bool)

	raw := strings.Split(string(dataBytes), "\n\n")
	for index, value := range raw[0] {
		if string(value) == "#" {
			binary := fmt.Sprintf("%09b", index)
			imageEnhAlgorithm[binary] = true
		}
	}

	for i, row := range strings.Split(raw[1], "\n") {
		for j, v := range row {
			if string(v) == "#" {
				coor := fmt.Sprintf("%d,%d", i, j)
				image[coor] = true
			}
		}
	}
	return imageEnhAlgorithm, image
}