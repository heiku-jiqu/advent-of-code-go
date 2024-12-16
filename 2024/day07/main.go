package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const INPUTFILE = "./2024/day07/input.txt"

type Equation struct {
	TestVal int
	Nums    []int
}

func main() {
	f, err := os.Open(INPUTFILE)
	if err != nil {
		log.Print(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var eqns []Equation
	for scanner.Scan() {
		line := scanner.Text()
		eqn := strings.Split(line, ":")
		testValue := eqn[0]
		testValueInt := strToInt(testValue)
		numbers := strings.Split(strings.TrimSpace(eqn[1]), " ")
		numbersInt := make([]int, len(numbers))
		for i, val := range numbers {
			numbersInt[i] = strToInt(val)
		}
		eqns = append(eqns, Equation{testValueInt, numbersInt})
	}

	part1ans := 0
	part2ans := 0
	for _, eqn := range eqns {
		// log.Print(eqn, eqn.isTally())
		if eqn.isTally(1) {
			part1ans += eqn.TestVal
		}
		if eqn.isTally(2) {
			part2ans += eqn.TestVal
		}
	}
	log.Print("Part 1 answer is: ", part1ans)
	log.Print("Part 2 answer is: ", part2ans)
}

// Check if equation tallies
func (e Equation) isTally(part int) bool {
	// recursively check by adding or multiplying
	addOrMultiply := e.recurse(0, 0)
	if part == 1 {
		return addOrMultiply
	}
	if part == 2 {
		return e.recurseConcat(0, 0)
	}
	panic("invalid part")
}

// recursively calculates left to right
// starting from 0,0
func (e Equation) recurse(cumulative int, idx int) bool {
	if idx == len(e.Nums) {
		return cumulative == e.TestVal
	}
	add := e.recurse(cumulative+e.Nums[idx], idx+1)
	multiply := e.recurse(max(cumulative, 1)*e.Nums[idx], idx+1)

	return add || multiply
}

func (e Equation) recurseConcat(cumulative int, idx int) bool {
	if idx == len(e.Nums) {
		return cumulative == e.TestVal
	}
	add := e.recurseConcat(cumulative+e.Nums[idx], idx+1)
	multiply := e.recurseConcat(max(cumulative, 1)*e.Nums[idx], idx+1)
	concatCumulative := strToInt(intToStr(cumulative) + intToStr(e.Nums[idx]))
	concat := e.recurseConcat(concatCumulative, idx+1)

	return add || multiply || concat
}

// Convert string to int. Panics if fails
func strToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

// Convert int to string. Panics if fails
func intToStr(i int) string {
	s := strconv.Itoa(i)
	return s
}
