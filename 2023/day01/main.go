package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	f, err := os.Open("./2023/day01/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var sum int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		curr, err := parseString(scanner.Text())
		if err != nil {
			panic(err)
		}
		sum += curr
	}
	fmt.Printf("Part 1: %d", sum)
}

func parseString(s string) (int, error) {
	var firstIdx int
	var lastIdx int
	var err error
	runes := []rune(s)
	runesLastIdx := len(runes) - 1

	for i := range runes {
		if unicode.IsDigit(runes[i]) {
			firstIdx = i
			break
		}
	}
	for i := range s {
		if unicode.IsDigit(runes[runesLastIdx-i]) {
			lastIdx = runesLastIdx - i
			break
		}
	}
	combined, err := strconv.Atoi(string(runes[firstIdx]) + string(runes[lastIdx]))
	if err != nil {
		return 0, err
	}
	return combined, nil
}
