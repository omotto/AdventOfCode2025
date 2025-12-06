package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"advent2025/pkg/file"
)

func getSumOperations(s []string) int {
	solutions := make([]int, len(s[len(s)-1]))
	operations := strings.Fields(s[len(s)-1])
	for i, line := range s[:len(s)-1] {
		numbers := strings.Fields(line)
		for j, number := range numbers {
			n, _ := strconv.Atoi(number)
			if i == 0 {
				solutions[j] = n
			} else if operations[j] == "+" {
				solutions[j] = solutions[j] + n
			} else {
				solutions[j] = solutions[j] * n
			}
		}
	}
	sum := 0
	for _, solution := range solutions {
		sum += solution
	}
	return sum
}

func getSumOperations2(s []string) int {
	var (
		newColumn []int
		totalSum  int = 0
	)
	for x := len(s[0]) - 1; x >= 0; x-- {
		number := 0
		isOperated := false
		for y := 0; y < len(s); y++ {
			if s[y][x] == ' ' {
				continue
			} else if s[y][x] == '+' {
				result := number
				for i := 0; i < len(newColumn); i++ {
					result += newColumn[i]
				}
				totalSum += result
				isOperated = true
				newColumn = []int{}
				x--
				break
			} else if s[y][x] == '*' {
				result := number
				for i := 0; i < len(newColumn); i++ {
					result *= newColumn[i]
				}
				totalSum += result
				isOperated = true
				newColumn = []int{}
				x--
				break
			} else {
				n, _ := strconv.Atoi(string(s[y][x]))
				number = number*10 + n
			}
		}
		if !isOperated {
			newColumn = append(newColumn, number)
		}
	}
	return totalSum
}

func main() {
	absPathName, _ := filepath.Abs("src/day06/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getSumOperations(output))
	fmt.Println(getSumOperations2(output))
}
