package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"advent2025/pkg/file"
)

func replaceAt(s string, idx int, char rune) string {
	runes := []rune(s)
	if idx >= 0 && idx < len(runes) {
		runes[idx] = char
		return string(runes)
	}
	return s
}

func getNumSplits(s []string) int {
	total := 0
	for y := 1; y < len(s); y++ {
		for x := 0; x < len(s[y]); x++ {
			if s[y-1][x] == 'S' || s[y-1][x] == '|' {
				if s[y][x] == '^' {
					total++
					s[y] = replaceAt(s[y], x+1, '|')
					s[y] = replaceAt(s[y], x-1, '|')
				} else {
					s[y] = replaceAt(s[y], x, '|')
				}
			}
		}
	}
	return total
}

func recursive(s []string, x, y int, cache map[struct{ x, y int }]int) int {
	for ; y < len(s); y++ {
		if s[y][x] == '^' {
			if v, ok := cache[struct{ x, y int }{x, y}]; ok {
				return v
			} else {
				v = recursive(s, x-1, y, cache) + recursive(s, x+1, y, cache)
				cache[struct{ x, y int }{x, y}] = v
				return v
			}
		}
	}
	return 1
}

func getTimelines(s []string) int {
	y, x := 0, strings.Index(s[0], "S")
	cache := make(map[struct{ x, y int }]int)
	return recursive(s, x, y, cache)
}

func main() {
	absPathName, _ := filepath.Abs("src/day07/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getNumSplits(output))
	fmt.Println(getTimelines(output))
}
