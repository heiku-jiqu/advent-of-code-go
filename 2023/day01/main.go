package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	f, err := os.Open("./2023/day01/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var sum int
	var sum2 int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		scanText := scanner.Text()
		curr, err := parseString(scanText)
		if err != nil {
			panic(err)
		}
		sum += curr

		curr2, err := parseStringPart2(scanText)
		if err != nil {
			panic(err)
		}
		sum2 += curr2
	}
	fmt.Printf("Part 1: %d", sum)
	fmt.Printf("Part 2: %d", sum2)
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

func parseStringPart2(s string) (int, error) {
	firstIdx := -1
	lastIdx := -1
	var first string
	var last string
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

	spellFirstIdx, spellLastIdx, spellFirstVal, spellLastVal := parseSpelledNumber(s)
	if firstIdx == -1 || spellFirstIdx != -1 && spellFirstIdx < firstIdx {
		firstIdx = spellFirstIdx
		first = fmt.Sprint(spellFirstVal)
	} else {
		first = string(runes[firstIdx])
	}
	if lastIdx == -1 || spellLastIdx != -1 && spellLastIdx > lastIdx {
		lastIdx = spellLastIdx
		last = fmt.Sprint(spellLastVal)
	} else {
		last = string(runes[lastIdx])
	}

	combined, err := strconv.Atoi(first + last)
	fmt.Printf("%s: %d\n", s, combined)
	if err != nil {
		return 0, err
	}
	return combined, nil
}

// returns index of substring
func parseSpelledNumber(s string) (firstIdx, lastIdx int, firstVal, lastVal int) {
	spelled := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	firstIdx = -1
	lastIdx = -1
	for i, spelledNum := range spelled {
		firstFind := strings.Index(s, spelledNum)
		if firstFind != -1 && (firstFind < firstIdx || firstIdx == -1) {
			firstIdx = firstFind
			firstVal = i + 1
		}
		lastFind := strings.LastIndex(s, spelledNum)
		if lastFind != -1 && (lastFind > lastIdx || lastIdx == -1) {
			lastIdx = lastFind
			lastVal = i + 1
		}
	}
	return
}

// kfdqhfml55
// 4dpdvdqdqq6mtnl
// ckpqmllrqrcbgdcmzpxcpqxtx66
