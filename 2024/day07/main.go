package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const INPUTFILE = "./2024/day07/input_test.txt"

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
	for _, eqn := range eqns {
		// log.Print(eqn, eqn.isTally())
		if eqn.isTally() {
			part1ans += eqn.TestVal
		}
	}
	log.Print("Part 1 answer is: ", part1ans)
}

// Check if equation tallies
func (e Equation) isTally() bool {
	// recursively check by adding or multiplying
	return e.recurse(0, 0)
}

// recursively calculates left to right
// starting from 0,0
func (e Equation) recurse(cumulative int, idx int) bool {
	if idx == len(e.Nums) {
		return cumulative == e.TestVal
	}
	add := e.recurse(cumulative+e.Nums[idx], idx+1)
	multiply := e.recurse(cumulative*e.Nums[idx], idx+1)

	return add || multiply
}

// Convert string to int. Panics if fails
func strToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
