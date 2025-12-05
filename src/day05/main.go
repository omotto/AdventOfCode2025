package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"advent2025/pkg/file"
)

func getRangeIDs(s []string) ([][2]int, []int) {
	var (
		ranges [][2]int
		ids    []int
	)
	for _, line := range s {
		ss := strings.Split(line, "-")
		if len(line) == 0 {
			continue
		}
		if len(ss) == 2 {
			a, _ := strconv.Atoi(ss[0])
			b, _ := strconv.Atoi(ss[1])
			ranges = append(ranges, [2]int{a, b})
		} else {
			i, _ := strconv.Atoi(ss[0])
			ids = append(ids, i)
		}
	}
	return ranges, ids
}

func getNumFresh(s []string) int {
	num := 0
	ranges, ids := getRangeIDs(s)
	for _, id := range ids {
		for _, r := range ranges {
			if id >= r[0] && id <= r[1] {
				num++
				break
			}
		}
	}
	return num
}

func getNewRanges(ranges [][2]int) [][2]int {
	for {
		numRanges := len(ranges)
		for i := 0; i < len(ranges)-1; i++ {
			for j := i + 1; j < len(ranges); j++ {
				if ranges[j][0] >= ranges[i][0] && ranges[j][1] <= ranges[i][1] {
					ranges = append(ranges[:j], ranges[j+1:]...) // remove it from list
				} else if ranges[j][0] <= ranges[i][0] && ranges[j][1] >= ranges[i][1] {
					ranges[i][0], ranges[i][1] = ranges[j][0], ranges[j][1]
					ranges = append(ranges[:j], ranges[j+1:]...) // remove it from list
				} else if ranges[j][0] <= ranges[i][1] && ranges[j][0] >= ranges[i][0] {
					ranges[i][1] = ranges[j][1]
					ranges = append(ranges[:j], ranges[j+1:]...) // remove it from list
				} else if ranges[j][1] <= ranges[i][1] && ranges[j][1] >= ranges[i][0] {
					ranges[i][0] = ranges[j][0]
					ranges = append(ranges[:j], ranges[j+1:]...) // remove it from list
				}
			}
		}
		if numRanges == len(ranges) {
			break
		}
	}
	return ranges
}

func getNumUniqueFresh(s []string) int {
	num := 0
	ranges, _ := getRangeIDs(s)
	newRanges := getNewRanges(ranges)
	for _, r := range newRanges {
		num += r[1] - r[0] + 1
	}
	return num
}

func main() {
	absPathName, _ := filepath.Abs("src/day05/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getNumFresh(output))
	fmt.Println(getNumUniqueFresh(output))
}
