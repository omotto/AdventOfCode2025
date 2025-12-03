package main

import (
	"fmt"
	"math"
	"path/filepath"
	"strconv"

	"advent2025/pkg/file"
)

func getMaxJolage(bank []int, digits int) int {
	result, newIdx := 0, 0
	for j := 0; j < digits; j++ {
		maxValue, idx := 0, newIdx
		for i := idx; i < len(bank)-digits+j+1; i++ {
			if bank[i] > maxValue {
				maxValue = bank[i]
				newIdx = i + 1
			}
		}
		result += maxValue * int(math.Pow(10, float64(digits-j-1)))
	}
	return result
}

func getIntSlice(s string) []int {
	result := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		result[i], _ = strconv.Atoi(string(s[i]))
	}
	return result
}

func getSumJoltage(s []string, digits int) int {
	sum := 0
	for _, bank := range s {
		intBank := getIntSlice(bank)
		sum += getMaxJolage(intBank, digits)
	}
	return sum
}

func main() {
	absPathName, _ := filepath.Abs("src/day03/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getSumJoltage(output, 2))
	fmt.Println(getSumJoltage(output, 12))
}
