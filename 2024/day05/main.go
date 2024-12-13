package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const INPUTFILE = "./2024/day05/input.txt"

func main() {
	f, err := os.Open(INPUTFILE)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	part1ans := 0
	rules := make(map[int][]int)
	// First section. Parse order rules
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		orderRule := strings.Split(line, "|")
		if len(orderRule) > 2 {
			log.Panic("Expected order rule to only have 2 numbers, but got ", len(orderRule))
		}
		before, err := strconv.Atoi(orderRule[0])
		if err != nil {
			log.Panic("Failed to parse order rule LHS. ", err)
		}
		after, err := strconv.Atoi(orderRule[1])
		if err != nil {
			log.Panic("Failed to parse order rule RHS. ", err)
		}

		if priors, ok := rules[before]; ok {
			rules[before] = append(priors, after)
		} else {
			rules[before] = []int{after}
		}

		// log.Printf("Rule: %d before %d", before, after)
	}

	// Second section. Parse updates
	for scanner.Scan() {
		line := scanner.Text()
		pageUpdate := strings.Split(line, ",")
		pageUpdateNum := make([]int, len(pageUpdate))
		validUpdate := true
		for i, page := range pageUpdate {
			pageNum, err := strconv.Atoi(page)
			if err != nil {
				log.Panic("Failed to convert %s to int in line %s", page, line)
			}
			pageUpdateNum[i] = pageNum

			isValidPage := checkValidPageUpdateNum(pageUpdateNum, rules)
			if !isValidPage {
				validUpdate = false
				break
			}
		}
		if validUpdate {
			part1ans += pageUpdateNum[len(pageUpdateNum)/2]
		}
	}

	log.Print("Part 1 answer is: ", part1ans)
}

func checkValidPageUpdateNum(pageUpdateNum []int, rules map[int][]int) bool {
	for i, pageNum := range pageUpdateNum {
		for j := i + 1; j < len(pageUpdateNum); j++ {
			if priors, ok := rules[pageUpdateNum[j]]; ok {
				for _, prior := range priors {
					if prior == pageNum {
						// log.Print("this line is invalid")
						return false
					}
				}
			}
		}
	}
	return true
}
