package main

import (
	"fmt"
	"math"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"advent2025/pkg/file"
)

func cmpSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func getJoltages(joltages string) []int {
	strJoltages := strings.Split(joltages[1:len(joltages)-1], ",")
	finalJoltage := make([]int, len(strJoltages))
	for i := 0; i < len(strJoltages); i++ {
		finalJoltage[i], _ = strconv.Atoi(strJoltages[i])
	}
	return finalJoltage
}

func getButtons(buttons []string) [][]int {
	intButtons := make([][]int, len(buttons))
	for i, strButton := range buttons {
		values := strings.Split(strButton[1:len(strButton)-1], ",")
		intButtons[i] = make([]int, len(values))
		for j, value := range values {
			intButtons[i][j], _ = strconv.Atoi(value)
		}
	}
	return intButtons
}

func getFinalState(lights string) []int {
	finalState := make([]int, len(lights)-2)
	for i := 1; i < len(lights)-1; i++ {
		if lights[i] == '.' {
			finalState[i-1] = 0
		} else {
			finalState[i-1] = 1
		}
	}
	return finalState
}

func getMinCombinations(lights string, buttons []string, _ string) int {
	finalState := getFinalState(lights)
	initialState := make([]int, len(finalState)) // by default all zeroes
	intButtons := getButtons(buttons)
	type Tuple struct {
		state  []int
		clicks int
	}
	visited := make(map[string]int)
	queue := make([]Tuple, 0)
	queue = append(queue, Tuple{
		state:  initialState,
		clicks: 0,
	})
	for len(queue) > 0 {
		tile := queue[0]  // Get first
		queue = queue[1:] // Remove it
		if cmpSlice(tile.state, finalState) {
			return tile.clicks
		}
		if _, ok := visited[fmt.Sprint(tile.state)]; ok {
			continue
		}
		visited[fmt.Sprint(tile.state)] = tile.clicks
		for _, button := range intButtons {
			newState := make([]int, len(tile.state))
			copy(newState, tile.state)
			for i := 0; i < len(button); i++ {
				if newState[button[i]] == 0 {
					newState[button[i]] = 1
				} else {
					newState[button[i]] = 0
				}
			}
			if _, ok := visited[fmt.Sprint(newState)]; ok {
				continue
			}
			queue = append(queue, Tuple{
				state:  newState,
				clicks: tile.clicks + 1,
			})
		}
	}
	return -1
}

func getFewestButtonClicks(s []string) int {
	sum := 0
	for _, line := range s {
		parts := strings.Split(line, " ")
		lights := parts[0]
		joltage := parts[len(parts)-1]
		buttons := parts[1 : len(parts)-1]
		sum += getMinCombinations(lights, buttons, joltage)
	}
	return sum
}

type ButtonCombination struct {
	jolatges  []int
	numClicks int
}

func getCombinations(buttons [][]int, joltagesLenght int) []ButtonCombination {
	combinations := make([]ButtonCombination, 0)
	for i := 0; i < (1 << len(buttons)); i++ { // 2^len(buttons) because the each button has two states clicked or not
		clicks := 0
		jolatge := make([]int, joltagesLenght)
		for j := 0; j < len(buttons); j++ {
			if (i & (1 << j)) != 0 { // if the j-th button is clicked
				clicks++
				for _, k := range buttons[j] {
					jolatge[k]++
				}
			}
		}
		combinations = append(combinations, ButtonCombination{jolatge, clicks})
	}
	return combinations
}

func getMinCombinations2(joltages []int, combinations []ButtonCombination) (int, bool) {
	// check if joltages are zero
	isZero := true
	for i := 0; i < len(joltages); i++ {
		if joltages[i] != 0 {
			isZero = false
			break
		}
	}
	if isZero {
		return 0, true
	}
	result := math.MaxInt
	for _, combination := range combinations {
		isMinorOrEqual, isEqualModulo2 := true, true
		for i := 0; i < len(combination.jolatges); i++ {
			if combination.jolatges[i] > joltages[i] {
				isMinorOrEqual = false
			}
			if combination.jolatges[i]%2 != joltages[i]%2 {
				isEqualModulo2 = false
			}
		}
		if !isMinorOrEqual || !isEqualModulo2 {
			continue
		}
		nextJolatges := make([]int, len(joltages))
		for i := 0; i < len(joltages); i++ {
			nextJolatges[i] = (joltages[i] - combination.jolatges[i]) / 2
		}
		if val, ok := getMinCombinations2(nextJolatges, combinations); ok {
			if n := 2*val + combination.numClicks; n < result {
				result = n
			}
		}
	}
	if result < math.MaxInt {
		return result, true
	}
	return 0, false
}

func getFewestButtonClicks2(s []string) int {
	sum := 0
	for _, line := range s {
		parts := strings.Split(line, " ")
		joltage := parts[len(parts)-1]
		buttons := parts[1 : len(parts)-1]
		intButtons := getButtons(buttons)
		intJoltages := getJoltages(joltage)
		// Get unique button combinations
		combinations := getCombinations(intButtons, len(intJoltages))
		clicks, _ := getMinCombinations2(intJoltages, combinations)
		sum += clicks
	}
	return sum
}

func main() {
	absPathName, _ := filepath.Abs("src/day10/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getFewestButtonClicks(output))
	fmt.Println(getFewestButtonClicks2(output))
}

// -- NOT WORKING solution

func getMinCombinationsXX(_ string, buttons []string, joltages string) int {
	finalJoltage := getJoltages(joltages)
	intButtons := getButtons(buttons)
	A := make([][]int64, len(finalJoltage))
	b := make([]int64, len(finalJoltage))
	// Generate Matrices Ax = b
	for i := 0; i < len(finalJoltage); i++ {
		b[i] = int64(finalJoltage[i])
		A[i] = make([]int64, len(intButtons))
		for j, button := range intButtons {
			if slices.Contains(button, i) {
				A[i][j] = 1.0
			}
		}
	}
	// Extended Matrix
	M := make([][]int64, len(A))
	for i := 0; i < len(A); i++ {
		M[i] = make([]int64, len(A[i])+1)
		for j := 0; j < len(A[i]); j++ {
			M[i][j] = A[i][j]
		}
		M[i][len(A[i])] = b[i]
	}
	// -- Reduce Matrix
	reducedRows := make(map[int]struct{})
	reducedOn := make([]int, len(intButtons))
	for x := 0; x < len(intButtons); x++ {
		y := 0
		for ; y < len(finalJoltage); y++ {
			if _, ok := reducedRows[y]; !ok && M[y][x] != 0 {
				break
			}
		}
		if y < len(finalJoltage) {
			reducedOn[x] = y
			reducedRows[y] = struct{}{}
			factorA := M[y][x]
			for i := 0; i < len(finalJoltage); i++ {
				if _, ok := reducedRows[i]; !ok && M[i][x] != 0 {
					factorB := M[i][x]
					M[i][x] = 0
					for j := x + 1; j < len(M[i]); j++ {
						M[i][j] = factorA*M[i][j] - factorB*M[y][j]
					}
				}
			}
		} else {
			reducedOn[x] = -1
		}
	}
	// -- Solve the reduced matrix
	buttonClicks := make([]int64, len(intButtons))
	minClicks := 1000
	solveEquations(M, reducedOn, &buttonClicks, 0, len(intButtons)-1, &minClicks)
	return minClicks
}

func solveEquations(M [][]int64, reducedOn []int, buttonClicks *[]int64, clicks, unknownIdx int, minClicks *int) {
	if unknownIdx < 0 {
		if clicks < *minClicks {
			*minClicks = clicks
		}
		return
	}
	if clicks >= *minClicks {
		return
	}
	y := reducedOn[unknownIdx]
	if y >= 0 {
		value := M[y][len(M[y])-1]
		for x := unknownIdx + 1; x < len(M[y])-1; x++ {
			value -= M[y][x] * (*buttonClicks)[x]
		}
		factor := M[y][unknownIdx]
		if factor < 0 {
			factor = -factor
			value = -value
		}
		if value < 0 {
			return
		}
		if factor > 1 {
			if value%factor != 0 {
				return
			}
			value /= factor
		}
		clicks += int(value)
		if clicks < *minClicks {
			(*buttonClicks)[unknownIdx] = value
			solveEquations(M, reducedOn, buttonClicks, clicks, unknownIdx-1, minClicks)
		}
	} else {
		(*buttonClicks)[unknownIdx] = 0
		for clicks < *minClicks {
			solveEquations(M, reducedOn, buttonClicks, clicks, unknownIdx-1, minClicks)
			(*buttonClicks)[unknownIdx]++
		}
	}
}

// Old lp Ax = b solution
/*
import "github.com/alexchao1/golp"

func getMinCombinations2(_ string, buttons []string, joltages string) int {
	finalJoltage := getJoltages(joltages)
	intButtons := getButtons(buttons)
	const maxClicks = 1000 // adjust as needed
	numJoltages := len(finalJoltage)
	lp := golp.NewLP(0, len(intButtons))
	lp.SetVerboseLevel(golp.NEUTRAL)
	objectiveCoeffs := make([]float64, len(intButtons))
	for i := 0; i < len(intButtons); i++ {
		objectiveCoeffs[i] = 1.0
		lp.SetInt(i, true)
		lp.SetBounds(i, 0.0, float64(maxClicks)) // Number of clicks per button >= 0 to macClicks restriction for equations
	}
	lp.SetObjFn(objectiveCoeffs)
	for i := 0; i < numJoltages; i++ {
		var entries []golp.Entry
		for j, button := range intButtons { // set A matrix
			if slices.Contains(button, i) {
				entries = append(entries, golp.Entry{Col: j, Val: 1.0})
			}
		}
		if err := lp.AddConstraintSparse(entries, golp.EQ, float64(finalJoltage[i])); err != nil { // set b Matrix
			panic(err)
		}
	}
	if lp.Solve() != golp.OPTIMAL {
		return 0
	}
	solution := lp.Variables()
	clicks := 0
	for _, val := range solution {
		clicks += int(val + 0.5) // round float64 up to int
	}
	return clicks
}
*/
