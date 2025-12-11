package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"advent2025/pkg/file"
)

func getMap(s []string) map[string][]string {
	values := make(map[string][]string)
	for _, line := range s {
		parts := strings.Split(line, ": ")
		outs := strings.Split(parts[1], " ")
		values[parts[0]] = outs
	}
	return values
}

func recursive(values map[string][]string, input, end string) int {
	outputs := values[input]
	paths := 0
	for _, output := range outputs {
		if output == end {
			return 1
		}
		paths += recursive(values, output, end)
	}
	return paths
}

func getNumPaths(s []string) int {
	values := getMap(s)
	return recursive(values, "you", "out")
}

// -- part 2

type State struct {
	id           string
	isfft, isdac bool
}

func recursive2(values map[string][]string, input, end string, fft, dac bool, visited map[State]int) int {
	outputs := values[input]
	if value, ok := visited[State{fmt.Sprintf("%s:%s", input, outputs), fft, dac}]; ok {
		return value
	}
	if input == end {
		if fft && dac {
			return 1
		}
		return 0
	}
	paths := 0
	for _, output := range outputs {
		isfft, isdac := fft, dac
		if input == "fft" {
			isfft = true
		}
		if input == "dac" {
			isdac = true
		}
		paths += recursive2(values, output, end, isfft, isdac, visited)
	}
	visited[State{fmt.Sprintf("%s:%s", input, outputs), fft, dac}] = paths
	return paths
}

func getNumPaths2(s []string) int {
	values := getMap(s)
	return recursive2(values, "svr", "out", false, false, make(map[State]int))
}

func main() {
	absPathName, _ := filepath.Abs("src/day11/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getNumPaths(output))
	fmt.Println(getNumPaths2(output))
}
