package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const INPUTFILE = "./2024/day06/input.txt"

func main() {
	f, err := os.Open(INPUTFILE)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// parse the map
	posnMap := [][]rune{}
	startingLoc := [2]int{0, 0}
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if j := strings.Index(line, "^"); j != -1 {
			startingLoc = [2]int{i, j}
		}
		posnMap = append(posnMap, []rune(line))

		i++
	}
	// log.Print("starting location is: ", startingLoc)

	guard := NewGuard(posnMap, startingLoc)
	for guard.isInsideMap {
		guard.traverse()
	}

	log.Print("Part 1 answer: ", len(guard.distinctPos))
}

type Guard struct {
	posnMap     [][]rune
	isInsideMap bool
	route       [][2]int
	distinctPos map[[2]int]struct{}
	currentLoc  [2]int
	direction   [2]int
}

func NewGuard(posnMap [][]rune, startingLoc [2]int) *Guard {
	distinctPos := make(map[[2]int]struct{})
	distinctPos[startingLoc] = struct{}{}
	return &Guard{
		posnMap:     posnMap,
		isInsideMap: true,
		route:       [][2]int{startingLoc},
		distinctPos: distinctPos,
		currentLoc:  startingLoc,
		direction:   [2]int{-1, 0},
	}
}

func (g *Guard) traverse() {
	// log.Print("before move: isInsideMap ", g.isInsideMap, ", currentLoc ", g.currentLoc)
	if !g.isInsideMap {
		// log.Print("Warning: ignoring traversal because guard is already out of map: ", g.currentLoc)
		return
	}
	newLoc := [2]int{g.currentLoc[0] + g.direction[0], g.currentLoc[1] + g.direction[1]}

	// check if guard is out of map in newLoc
	if newLoc[0] < 0 || newLoc[0] >= len(g.posnMap) || newLoc[1] < 0 || newLoc[1] >= len(g.posnMap[0]) {
		// move and dont need to do anything else
		g.isInsideMap = false
		g.currentLoc = newLoc
		return
	}

	// check if newLoc is blocked,
	if g.posnMap[newLoc[0]][newLoc[1]] == '#' {
		// turn direction before we move
		// math coordinate transformation
		newDirection := [2]int{
			g.direction[1],
			g.direction[0] * -1,
		}
		g.direction = newDirection
		newLoc = [2]int{g.currentLoc[0] + g.direction[0], g.currentLoc[1] + g.direction[1]}
		// log.Print("turning! new direction is: ", g.direction)
	}
	// add new location into traveled since we are still inside map, then move
	if _, contain := g.distinctPos[newLoc]; !contain {
		g.distinctPos[newLoc] = struct{}{}
		g.route = append(g.route, newLoc)
	}
	g.currentLoc = newLoc

	// log.Print("after move: isInsideMap ", g.isInsideMap, ", currentLoc ", g.currentLoc)

}

func findStartingLoc(posnMap [][]rune) [2]int {
	// TODO: find starting loc
	return [2]int{0, 0}
}
