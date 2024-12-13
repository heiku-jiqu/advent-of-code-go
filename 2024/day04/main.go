package main

import (
	"bufio"
	"errors"
	"log"
	"os"
)

const INPUTFILE = "./2024/day04/input.txt"

func main() {
	f, err := os.Open(INPUTFILE)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var puzzle [][]rune

	for scanner.Scan() {
		line := scanner.Text()
		puzzle = append(puzzle, []rune(line))
	}

	part1ans := 0
	part2ans := 0
	for rowNum := range puzzle {
		for colNum := range puzzle[rowNum] {
			if puzzle[rowNum][colNum] == 'X' {
				// log.Print("Found 'X' at ", rowNum, colNum)
				allIndices := generateIndices(rowNum, colNum, len(puzzle[0])-1, len(puzzle)-1)
				for _, indices := range allIndices {
					if matchXMAS(puzzle, indices) {
						part1ans++
					}
				}
			}

			if puzzle[rowNum][colNum] == 'A' {
				cornerIndices, err := generateIndicesCorners(rowNum, colNum, len(puzzle[0])-1, len(puzzle)-1)
				if err == nil && matchMS(puzzle, cornerIndices) {
					part2ans++
				}
			}
		}
	}

	log.Print("Part 1 answer is ", part1ans)
	log.Print("Part 2 answer is ", part2ans)

}

// matchXMAS returns whether the corresponding text indexed by indices in puzzle is
// "XMAS"
func matchXMAS(puzzle [][]rune, indices [4][2]int) bool {
	correct := []rune("XMAS")
	for i, index := range indices {
		if correct[i] != puzzle[index[0]][index[1]] {
			return false
		}
	}
	return true
}

// generateIndices returns slice of array(4) of indexes that goes in all 8 directions
// does not return incomplete or invalid indices (i.e <0 or > max)
func generateIndices(rowNum, colNum, maxRowNum, maxColNum int) [][4][2]int {
	var output [][4][2]int

	var leftIndices [4][2]int
	var rightIndices [4][2]int
	var upIndices [4][2]int
	var downIndices [4][2]int

	var upLeftIndices [4][2]int
	var upRightIndices [4][2]int
	var downLeftIndices [4][2]int
	var downRightIndices [4][2]int

	for i := range 4 {
		leftIndices[i] = [2]int{rowNum - i, colNum}
		rightIndices[i] = [2]int{rowNum + i, colNum}
		upIndices[i] = [2]int{rowNum, colNum + i}
		downIndices[i] = [2]int{rowNum, colNum - i}

		upLeftIndices[i] = [2]int{rowNum - i, colNum + i}
		upRightIndices[i] = [2]int{rowNum + i, colNum + i}
		downLeftIndices[i] = [2]int{rowNum - i, colNum - i}
		downRightIndices[i] = [2]int{rowNum + i, colNum - i}
	}

	all := [][4][2]int{
		leftIndices,
		rightIndices,
		upIndices,
		downIndices,
		upLeftIndices,
		upRightIndices,
		downLeftIndices,
		downRightIndices,
	}

	for _, indices := range all {
		if outOfBound(indices, maxRowNum, maxColNum) {
			continue
		}
		output = append(output, indices)
	}

	return output
}

// returns whether any of the element in indices is out of bounds.
func outOfBound(indices [4][2]int, maxRowNum, maxColNum int) bool {
	for _, index := range indices {
		if index[0] < 0 || index[0] > maxRowNum || index[1] < 0 || index[1] > maxColNum {
			return true
		}
	}
	return false
}

// generate indices for all 4 corners,
// indices ordered in {topleft, topright, bottomleft, bottomright}
// returns error if it is the set of indices contains outofbounds
func generateIndicesCorners(rowNum, colNum, maxRowNum, maxColNum int) ([4][2]int, error) {
	output := [4][2]int{
		[2]int{rowNum - 1, colNum - 1},
		[2]int{rowNum - 1, colNum + 1},
		[2]int{rowNum + 1, colNum - 1},
		[2]int{rowNum + 1, colNum + 1},
	}
	if outOfBound(output, maxRowNum, maxColNum) {
		return output, errors.New("Generated out of bound indices")
	}
	return output, nil
}

// check whether diagonal corners forms two "MAS" for part 2
// indices should be ordered in {topleft, topright, bottomleft, bottomright}
func matchMS(puzzle [][]rune, indices [4][2]int) bool {
	toCheck := make([]rune, 4)
	for i, index := range indices {
		toCheck[i] = puzzle[index[0]][index[1]]
	}
	valids := [][]rune{
		[]rune("MSMS"), []rune("MMSS"), []rune("SSMM"), []rune("SMSM"),
	}
	for _, valid := range valids {
		if string(valid) == string(toCheck) {
			return true
		}
	}
	return false
}
