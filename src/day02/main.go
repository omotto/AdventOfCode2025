package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"advent2025/pkg/file"
)

func checkInvalid(str string) bool {
	a := str[:len(str)/2]
	b := str[len(str)/2:]
	return a == b
}

func checkInvalid2(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		invalid := true
		for j := i + 1; j < len(str); j = j + i + 1 {
			a := str[j-i-1 : j]
			if j+i+1 > len(str) {
				invalid = false
				break
			}
			b := str[j : j+i+1]
			if a != b {
				invalid = false
				break
			}
		}
		if invalid {
			return true
		}
	}
	return false
}

func getSumInvalidIDs(s []string, checkInvalid func(string) bool) int {
	sum := 0
	ranges := strings.Split(s[0], ",")
	for _, r := range ranges {
		v := strings.Split(r, "-")
		init, _ := strconv.Atoi(v[0])
		end, _ := strconv.Atoi(v[1])
		for i := init; i <= end; i++ {
			if checkInvalid(strconv.Itoa(i)) {
				sum += i
			}
		}
	}
	return sum
}

func main() {
	absPathName, _ := filepath.Abs("src/day02/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getSumInvalidIDs(output, checkInvalid))
	fmt.Println(getSumInvalidIDs(output, checkInvalid2))
}
