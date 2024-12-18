package main

import (
	"bufio"
	"log"
	"os"
)

const INPUTFILE = "./2024/day08/input.txt"

var Nodes map[rune][][2]int

func main() {
	f, err := os.Open(INPUTFILE)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	Nodes = make(map[rune][][2]int)
	nodeMap := [][]rune{}

	i := 0
	for scanner.Scan() {
		row := scanner.Text()
		// log.Print(row)
		nodeMap = append(nodeMap, []rune(row))
		for j, char := range row {
			if char != '.' {
				if _, exist := Nodes[char]; exist {
					Nodes[char] = append(Nodes[char], [2]int{i, j})
				} else {
					Nodes[char] = [][2]int{[2]int{i, j}}
				}
			}
		}
		i++
	}
	uniqueNodes := make(map[[2]int]struct{})
	for _, antennaPosn := range Nodes {
		// log.Print("antenna: ", string(antenna))
		// log.Print("antenna posn: ", (antennaPosn))
		antinodePosn := calcPossibleAntinodesPosn(antennaPosn)
		for _, loc := range antinodePosn {
			// log.Print("antinode position:", loc, "off map: ", isOffTheMap(nodeMap, loc))
			if isOffTheMap(nodeMap, loc) {
				continue
			}
			if _, exist := uniqueNodes[loc]; !exist {
				uniqueNodes[loc] = struct{}{}
			}
		}
	}
	log.Print("Part 1 answer is: ", len(uniqueNodes))
}

// Given antennas positions, returns all possible antinode positions
// Does not guarantee it is inside map
func calcPossibleAntinodesPosn(antennas [][2]int) [][2]int {
	var out [][2]int
	if len(antennas) < 2 {
		return out
	}
	for i := 0; i < len(antennas); i++ {
		for j := i + 1; j < len(antennas); j++ {
			antennaOne := antennas[i]
			antennaTwo := antennas[j]
			distance := [2]int{
				antennaOne[0] - antennaTwo[0],
				antennaOne[1] - antennaTwo[1],
			}
			out = append(out, [2]int{antennaOne[0] + distance[0], antennaOne[1] + distance[1]})
			out = append(out, [2]int{antennaTwo[0] - distance[0], antennaTwo[1] - distance[1]})
		}
	}
	return out
}

func isOffTheMap(nodeMap [][]rune, posn [2]int) bool {
	if posn[0] < 0 || posn[1] < 0 || posn[0] >= len(nodeMap) || posn[1] >= len(nodeMap[0]) {
		return true
	}
	return false
}
