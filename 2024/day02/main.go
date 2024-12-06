package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const INPUTFILE string = "./2024/day02/input.txt"

func main() {
	f, err := os.Open(INPUTFILE)
	if err != nil {
		log.Panic("Failed to open ", INPUTFILE, "; ", err)
	}

	part1ans := 0
	part2ans := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		splitline := strings.Fields(line)
		parsedline, err := toInt(splitline)
		if err != nil {
			log.Panic("Failed to convert string to int: ", err)
		}
		if isSafe(parsedline) {
			part1ans += 1
		}

		if isSafeWithDampener(parsedline) {
			part2ans += 1
		}

	}

	log.Print("Day 2 Part 1 answer is: ", part1ans)
	log.Print("Day 2 Part 2 answer is: ", part2ans)
}

func toInt(x []string) ([]int, error) {
	out := make([]int, len(x))
	for i, val := range x {
		converted, err := strconv.Atoi(val)
		if err != nil {
			return out, err
		}
		out[i] = converted
	}
	return out, nil
}

func isSafe(x []int) bool {
	increasing := x[1]-x[0] > 0
	for i := 1; i < len(x); i++ {
		diff := x[i] - x[i-1]
		if diff == 0 || (diff > 0) != increasing {
			return false
		}
		if diff > 3 || diff < -3 {
			return false
		}
	}
	return true
}

func isSafeWithDampener(x []int) bool {
	if isSafe(x) {
		return true
	}

	for i := range x {
		removed := make([]int, 0, len(x)-1)
		for j, val := range x {
			if i != j {
				removed = append(removed, val)
			}
		}
		if isSafe(removed) {
			return true
		}
	}
	return false
}
