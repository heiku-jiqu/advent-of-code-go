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
		// log.Print("main guard is at: ", guard.currentLoc)
		guard.checkIfObstructWillLoop()
	}
	log.Print("Part 1 answer: ", len(guard.distinctPos))
	log.Print("Part 2 answer: ", len(guard.loopStoneLoc))

	log.Print("BRUTE FORCING NOW")
	day2ans := 0
	for pos, _ := range guard.distinctPos {
		if pos == startingLoc {
			continue
		}
		if tile := posnMap[pos[0]][pos[1]]; tile != '.' {
			continue
		}
		newPosnMap := make([][]rune, len(posnMap))
		for i, row := range posnMap {
			newRow := make([]rune, len(row))
			for j, val := range row {
				newRow[j] = val
			}
			newPosnMap[i] = newRow
		}
		newPosnMap[pos[0]][pos[1]] = '#'
		guardCopy := NewGuard(newPosnMap, startingLoc)
		// log.Print("placing stone at ", pos)
		// log.Print(guardCopy.posnMap)
		for guardCopy.isInsideMap {
			if guardCopy.checkIfVisitedAndSameDirection(guardCopy.peekNextLoc(), guardCopy.peekDirection()) {
				day2ans++
				break
			}
			guardCopy.traverse()
		}

	}

	log.Print("Part 2 answer: ", day2ans)
}

type Guard struct {
	posnMap      [][]rune
	isInsideMap  bool
	route        [][2]int
	distinctPos  map[[2]int][][2]int // key: location, val: []directions
	currentLoc   [2]int
	direction    [2]int
	loopStoneLoc map[[2]int]struct{}
}

func NewGuard(posnMap [][]rune, startingLoc [2]int) *Guard {
	distinctPos := make(map[[2]int][][2]int)
	startDirection := [2]int{-1, 0}
	distinctPos[startingLoc] = [][2]int{startDirection}
	loopStoneLoc := make(map[[2]int]struct{})
	return &Guard{
		posnMap:      posnMap,
		isInsideMap:  true,
		route:        [][2]int{startingLoc},
		distinctPos:  distinctPos,
		currentLoc:   startingLoc,
		direction:    startDirection,
		loopStoneLoc: loopStoneLoc,
	}
}

func (g *Guard) peekNextLoc() [2]int {
	newDirection := g.peekDirection()
	newLoc := [2]int{g.currentLoc[0] + newDirection[0], g.currentLoc[1] + newDirection[1]}
	return newLoc
}

func (g *Guard) peekDirection() [2]int {
	newDirection := g.direction
	for range 4 {
		newLoc := [2]int{g.currentLoc[0] + newDirection[0], g.currentLoc[1] + newDirection[1]}
		if newLoc[0] < 0 || newLoc[0] >= len(g.posnMap) || newLoc[1] < 0 || newLoc[1] >= len(g.posnMap[0]) {
			return g.direction
		}
		if g.posnMap[newLoc[0]][newLoc[1]] == '#' {
			newDirection = [2]int{
				newDirection[1],
				newDirection[0] * -1,
			}
		}
	}
	return newDirection
}

func (g *Guard) traverse() {
	// log.Print("before move: isInsideMap ", g.isInsideMap, ", currentLoc ", g.currentLoc)
	if !g.isInsideMap {
		// log.Print("Warning: ignoring traversal because guard is already out of map: ", g.currentLoc)
		return
	}
	newLoc := g.peekNextLoc()
	newDirection := g.peekDirection()
	g.direction = newDirection

	// check if guard is out of map in newLoc
	if newLoc[0] < 0 || newLoc[0] >= len(g.posnMap) || newLoc[1] < 0 || newLoc[1] >= len(g.posnMap[0]) {
		// move and dont need to do anything else
		g.isInsideMap = false
		g.currentLoc = newLoc
		return
	}

	// add new location into traveled since we are still inside map, then move
	if directions, contain := g.distinctPos[newLoc]; !contain {
		// havent been here before
		g.distinctPos[newLoc] = [][2]int{g.direction}
	} else {
		// been here before, add where we are facing
		g.distinctPos[newLoc] = append(directions, g.direction)
	}
	g.route = append(g.route, newLoc)
	g.currentLoc = newLoc

	// log.Print("after move: isInsideMap ", g.isInsideMap, ", currentLoc ", g.currentLoc)
}

func (g *Guard) Copy() *Guard {
	// copy to new guard
	newPosnMap := make([][]rune, len(g.posnMap))
	for i, row := range g.posnMap {
		newRow := make([]rune, len(row))
		for j, val := range row {
			newRow[j] = val
		}
		newPosnMap[i] = newRow
	}
	newDistinctPos := make(map[[2]int][][2]int)
	for key, val := range g.distinctPos {
		newVal := make([][2]int, len(val))
		for i, v := range val {
			newVal[i] = v
		}
		newDistinctPos[key] = newVal
	}
	newRoute := make([][2]int, len(g.route))
	for i, v := range g.route {
		newRoute[i] = v
	}

	return &Guard{
		posnMap:     newPosnMap,
		isInsideMap: g.isInsideMap,
		route:       newRoute,
		distinctPos: newDistinctPos, // key: location, val: []directions
		currentLoc:  g.currentLoc,
		direction:   g.direction,
	}
}

// Returns whether placing an obstruction on next position will cause a loop in Guard's path
//
// Done by checking whether guard has been in the same tile+direction after
// obstructing, turning right and keep moving straight (don't have to be just move 1 tile)
func (g *Guard) checkIfObstructWillLoop() {
	stoneLoc := [2]int{
		g.currentLoc[0] + g.direction[0],
		g.currentLoc[1] + g.direction[1],
	}
	if stoneLoc[0] < 0 || stoneLoc[0] >= len(g.posnMap) || stoneLoc[1] < 0 || stoneLoc[1] >= len(g.posnMap[0]) {
		return
	}
	if _, ok := g.distinctPos[stoneLoc]; ok {
		return
	}
	if _, ok := g.loopStoneLoc[stoneLoc]; ok {
		return
	}
	if g.posnMap[stoneLoc[0]][stoneLoc[1]] == '#' {
		return
	}

	simulatedGuard := g.Copy()
	simulatedGuard.posnMap[stoneLoc[0]][stoneLoc[1]] = '#'

	// keep walking in new direction and checking whether guard has been in the same tile+direction
	for simulatedGuard.isInsideMap {
		// log.Print(simulatedGuard.currentLoc, simulatedGuard.direction, stoneLoc, simulatedGuard.peekNextLoc(), simulatedGuard.peekDirection())
		if simulatedGuard.checkIfVisitedAndSameDirection(simulatedGuard.peekNextLoc(), simulatedGuard.peekDirection()) {
			g.loopStoneLoc[stoneLoc] = struct{}{}
			return
		}
		simulatedGuard.traverse()
	}
}

func (g *Guard) checkIfVisitedAndSameDirection(loc [2]int, direction [2]int) bool {
	// check if visited here before
	if visitedDirections, visited := g.distinctPos[loc]; visited {
		// check if direction same as last time
		for _, visitedDirection := range visitedDirections {
			if direction == visitedDirection {
				// log.Print("found loop at current pos: ", g.currentLoc)
				return true
			}
		}
	}
	return false
}
