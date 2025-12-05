package main

import (
	"fmt"
	"path/filepath"

	"advent2025/pkg/file"
)

func getSumRolls(s []string) int {
	sum := 0
	for y := 0; y < len(s); y++ {
		for x := 0; x < len(s[y]); x++ {
			if s[y][x] != '@' {
				continue
			}
			numAdjRolls := 0
			if y > 0 && s[y-1][x] == '@' {
				numAdjRolls++
			}
			if y < len(s)-1 && s[y+1][x] == '@' {
				numAdjRolls++
			}
			if x > 0 && s[y][x-1] == '@' {
				numAdjRolls++
			}
			if x < len(s[y])-1 && s[y][x+1] == '@' {
				numAdjRolls++
			}
			if y > 0 && x > 0 && s[y-1][x-1] == '@' {
				numAdjRolls++
			}
			if y < len(s)-1 && x < len(s[y])-1 && s[y+1][x+1] == '@' {
				numAdjRolls++
			}
			if y > 0 && x < len(s[y])-1 && s[y-1][x+1] == '@' {
				numAdjRolls++
			}
			if y < len(s)-1 && x > 0 && s[y+1][x-1] == '@' {
				numAdjRolls++
			}
			if numAdjRolls < 4 {
				sum++
			}
		}
	}
	return sum
}

func getSumRolls2(s []string) int {
	totalSum := 0
	m := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		m[i] = make([]int, len(s[i]))
		for j := 0; j < len(s[i]); j++ {
			if s[i][j] == '@' {
				m[i][j] = 1
			} else {
				m[i][j] = 0
			}
		}
	}
	newM := make([][]int, len(m))
	for i := 0; i < len(m); i++ {
		newM[i] = make([]int, len(m[i]))
		copy(newM[i], m[i])
	}
	for {
		sum := 0
		for y := 0; y < len(m); y++ {
			for x := 0; x < len(m[y]); x++ {
				if m[y][x] != 1 {
					continue
				}
				numAdjRolls := 0
				if y > 0 && m[y-1][x] == 1 {
					numAdjRolls++
				}
				if y < len(m)-1 && m[y+1][x] == 1 {
					numAdjRolls++
				}
				if x > 0 && m[y][x-1] == 1 {
					numAdjRolls++
				}
				if x < len(m[y])-1 && m[y][x+1] == 1 {
					numAdjRolls++
				}
				if y > 0 && x > 0 && m[y-1][x-1] == 1 {
					numAdjRolls++
				}
				if y < len(m)-1 && x < len(m[y])-1 && m[y+1][x+1] == 1 {
					numAdjRolls++
				}
				if y > 0 && x < len(m[y])-1 && m[y-1][x+1] == 1 {
					numAdjRolls++
				}
				if y < len(m)-1 && x > 0 && m[y+1][x-1] == 1 {
					numAdjRolls++
				}
				if numAdjRolls < 4 {
					sum++
					newM[y][x] = 0
				}
			}
		}
		if sum == 0 {
			break
		}
		totalSum += sum
		for i := 0; i < len(m); i++ {
			m[i] = make([]int, len(m[i]))
			copy(m[i], newM[i])
		}
	}
	return totalSum
}

func main() {
	absPathName, _ := filepath.Abs("src/day04/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getSumRolls(output))
	fmt.Println(getSumRolls2(output))
}
