package main

import (
	"log"
	"os"
	"strconv"
)

var INPUTFILE = "./2024/day09/input_test.txt"

func main() {
	fileBytes, err := os.ReadFile(INPUTFILE)
	if err != nil {
		log.Panic(err)
	}

	inputString := string(fileBytes)
	// log.Print(inputString)
	expandedFileFormat := expandFileFormatInt(inputString)
	// log.Print(expandedFileFormat)
	compacted := compact(expandedFileFormat)
	// log.Print(compacted)
	part1ans := checksum(compacted)
	log.Print("Day 9 Part 1 answer is:", part1ans)

	compactedP2 := compactWholeFiles(expandedFileFormat)
	part2ans := checksum(compactedP2)
	log.Print("Day 9 Part 2 answer is:", part2ans)
}

// expands dense file format into actual blocks
// with ID number and -1 for empty
func expandFileFormatInt(denseString string) []int {
	out := []int{}
	for i, char := range denseString {
		if i%2 == 0 {
			// char represents size of file
			fileId := i / 2
			size, err := strconv.Atoi(string(char))
			if err != nil {
				log.Panic(err)
			}
			fullFileBlock := make([]int, size)
			for i := range size {
				fullFileBlock[i] = fileId
			}
			out = append(out, fullFileBlock...)
		} else {
			// char represents size of empty space
			size, err := strconv.Atoi(string(char))
			if err != nil {
				log.Panic(err)
			}
			fullEmptyBlock := make([]int, size)
			for i := range size {
				fullEmptyBlock[i] = -1
			}
			out = append(out, fullEmptyBlock...)
		}
	}
	return out
}

// compact returns a new copy that is compacted per block
func compact(expandedFS []int) []int {
	expanded := make([]int, len(expandedFS))
	copy(expanded, expandedFS)

	left := 0
	right := len(expanded) - 1

	for left < right {
		if expanded[left] == -1 && expanded[right] != -1 {
			expanded[left], expanded[right] = expanded[right], expanded[left]
			// log.Print("after swap: ", string(expanded))
			left++
			right--
		}
		if expanded[left] != -1 {
			left += 1
		}
		if expanded[right] == -1 {
			right -= 1
		}
	}
	return expanded
}

func checksum(compacted []int) int {
	out := 0
	for i, val := range compacted {
		if val != -1 {
			out += i * val
		}
	}
	return out
}

// compactWholeFiles returns a copy that is compacted by whole files
func compactWholeFiles(expandedFS []int) []int {
	compacted := make([]int, len(expandedFS))
	copy(compacted, expandedFS)

	left := 0
	right := len(compacted) - 1
	for left < right {
		log.Print("compacting: ", compacted, " left: ", left, " right: ", right)
		if compacted[left] == -1 && compacted[right] != -1 {
			//TODO: check all possible spaces!!!

			// check if left space can fit right file
			spaceSize := findSizeStartingLeft(compacted, left)
			fileSize := findSizeStartingRight(compacted, right)
			log.Print("space size: ", spaceSize, " file size: ", fileSize)
			if spaceSize >= fileSize {
				log.Print("can fit, moving file: ", compacted[right])
				// if can fit, move file
				for i := range fileSize {
					compacted[left+i], compacted[right-i] = compacted[right-i], compacted[left+i]
				}
			} else {
				// if cannot fit, move right pointer to next fileID that is not -1
				curr := compacted[right]
				for curr == compacted[right] {
					right--
				}
			}
		}

		// move left ptr until hit it is -1
		if compacted[left] != -1 {
			left++
		}

		// move right ptr until it is on file
		if compacted[right] == -1 {
			right--
		}
	}

	return compacted
}

func findSizeStartingLeft(expandedFS []int, leftIdx int) int {
	curr := expandedFS[leftIdx]
	length := 0
	for curr == expandedFS[leftIdx+length] {
		length++
	}
	return length
}

func findSizeStartingRight(expandedFS []int, rightIdx int) int {
	curr := expandedFS[rightIdx]
	length := 0
	for curr == expandedFS[rightIdx-length] {
		length++
	}
	return length
}
