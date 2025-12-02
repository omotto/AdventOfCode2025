package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"advent2025/pkg/file"
)

func getFinalNumber(s []string) int {
	var (
		zeros     = 0  // times to get zero value
		dialValue = 50 // initial dial value
		dir       byte // Left or Right
		v         int
	)
	for _, line := range s {
		_, _ = fmt.Sscanf(line, "%c%d", &dir, &v)
		v = v % 100    // for higher numbers than 100
		if dir == 76 { // ASCII value of L
			dialValue = dialValue - v
			if dialValue < 0 {
				dialValue = 100 + dialValue
			}
		} else { // ASCII value of R is 82
			dialValue = dialValue + v
			if dialValue > 99 {
				dialValue = dialValue - 100
			}
		}
		if dialValue == 0 {
			zeros += 1
		}
	}
	return zeros
}

func getZeros(s []string) int {
	zeros, dialValue := 0, 50 // initial dial value
	for _, line := range s {
		v, _ := strconv.Atoi(line[1:])
		zeros += v / 100
		v = v % 100
		if strings.ToLower(string(line[0])) == "l" {
			dialValue = dialValue - v
			if dialValue < 0 {
				dialValue = 100 + dialValue
				if dialValue+v != 100 {
					zeros += 1
				}
			} else if dialValue == 0 {
				zeros += 1
			}
		} else {
			dialValue = dialValue + v
			if dialValue > 99 {
				dialValue = dialValue - 100
				if dialValue-v != 0 {
					zeros += 1
				}
			} else if dialValue == 0 {
				zeros += 1
			}
		}
	}
	return zeros
}

func main() {
	absPathName, _ := filepath.Abs("src/day01/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getFinalNumber(output))
	fmt.Println(getZeros(output))
}
