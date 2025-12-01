package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"advent2025/pkg/file"
)

func getFinalNumber(s []string) int {
	zeros := 0
	dialValue := 50 // initial dial value
	for _, line := range s {
		v, _ := strconv.Atoi(line[1:])
		v = v % 100 // for higher numbers than 100
		if strings.ToLower(string(line[0])) == "l" {
			dialValue = dialValue - v
			if dialValue < 0 {
				dialValue = 100 + dialValue
			}
		} else {
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
	zeros := 0
	dialValue := 50 // initial dial value
	for _, line := range s {
		v, _ := strconv.Atoi(line[1:])
		zeros = zeros + v/100
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
