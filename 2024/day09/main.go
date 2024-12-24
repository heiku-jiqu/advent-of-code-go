package main

import (
	"log"
	"os"
	"strconv"
)

var INPUTFILE = "./2024/day09/input.txt"

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
}

// expands dense file format into actual blocks
// with ID number and . for empty
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

func compact(expandedFS []int) []int {
	expanded := expandedFS
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

func toInt(s string) []int {
	out := make([]int, len(s))
	for i, c := range s {
		if c == '.' {
			out[i] = -1
			continue
		}
		num, err := strconv.Atoi(string(c))
		if err != nil {
			log.Panic(err)
		}
		out[i] = num
	}
	return out
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
