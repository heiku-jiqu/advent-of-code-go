package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const INPUTFILE = "./2024/day10/input.txt"

func main() {
	f, err := os.Open(INPUTFILE)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var topoMap [][]int
	for scanner.Scan() {
		line := scanner.Text()
		parsedLine := toInt([]rune(line))
		topoMap = append(topoMap, parsedLine)
	}

	part1ans := 0
	for i := range topoMap {
		for j := range topoMap[0] {
			if topoMap[i][j] == 0 {
				// log.Print("found trailhead at ", i, ",", j)

				score := getTrailheadScore(topoMap, i, j)
				// log.Print("trailhead at ", i, ",", j, " score is ", score)

				part1ans += score
			}
		}
	}

	log.Print("Day 10 Part 1 Answer is: ", part1ans)
}

// Finds the score of the given trailhead position i,j in this topoMap.
// A trailhead gets a score for every reachable '9'.
func getTrailheadScore(topoMap [][]int, i, j int) int {
	visited := make(map[[2]int]struct{})
	return traverse(topoMap, i, j, visited)
}

func traverse(topoMap [][]int, i, j int, visited map[[2]int]struct{}) int {
	_, isVisited := visited[[2]int{i, j}]
	if topoMap[i][j] == 9 && !isVisited {
		visited[[2]int{i, j}] = struct{}{}
		return 1
	}
	left := 0
	if i > 0 && topoMap[i][j]-topoMap[i-1][j] == -1 {
		left = traverse(topoMap, i-1, j, visited)
	}
	right := 0
	if i < len(topoMap)-1 && topoMap[i][j]-topoMap[i+1][j] == -1 {
		right = traverse(topoMap, i+1, j, visited)
	}
	up := 0
	if j > 0 && topoMap[i][j]-topoMap[i][j-1] == -1 {
		up = traverse(topoMap, i, j-1, visited)
	}
	down := 0
	if j < len(topoMap[0])-1 && topoMap[i][j]-topoMap[i][j+1] == -1 {
		down = traverse(topoMap, i, j+1, visited)
	}
	return left + right + up + down
}

// Converts each rune into int
// Panics if invalid rune
func toInt(numbers []rune) []int {
	out := make([]int, len(numbers))
	for i, r := range numbers {
		num, err := strconv.Atoi(string(r))
		if err != nil {
			panic(err)
		}
		out[i] = num
	}
	return out
}
