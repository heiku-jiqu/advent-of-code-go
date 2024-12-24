package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

var INPUTFILE = "./2024/day09/input_test.txt"

func main() {
	fileBytes, err := os.ReadFile(INPUTFILE)
	if err != nil {
		log.Panic(err)
	}

	inputString := string(fileBytes)
	log.Print(inputString)
	expandedFileFormat := expandFileFormat(inputString)
	log.Print(expandedFileFormat)
}

// expands dense file format into actual blocks
// with ID number and . for empty
func expandFileFormat(denseString string) string {
	out := ""
	for i, char := range denseString {
		if i%2 == 0 {
			// char represents size of file
			fileId := i / 2
			size, err := strconv.Atoi(string(char))
			if err != nil {
				log.Panic(err)
			}
			fileIdStr := strconv.Itoa(fileId)
			fullFileBlock := strings.Repeat(fileIdStr, size)
			out = out + fullFileBlock
		} else {
			// char represents size of empty space
			size, err := strconv.Atoi(string(char))
			if err != nil {
				log.Panic(err)
			}
			fullEmptyBlock := strings.Repeat(".", size)
			out = out + fullEmptyBlock
		}
	}
	return out
}
