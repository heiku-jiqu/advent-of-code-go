package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const INPUTFILEPATH string = "./2024/day03/input.txt"

func main() {
	f, err := os.Open(INPUTFILEPATH)
	if err != nil {
		log.Panicln("Failed to open file: ", err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanRunes)

	// isParsingToken := false

	for scanner.Scan() {
		if scanner.Text() == "m" {

		}
	}
}

func ScanMul(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF {
		return 0, nil, bufio.ErrFinalToken
	}
	str := string(data)
	toAdvance := 0
	for i := range str {
		if str[i] == 'm' {
			toAdvance = i
			str = str[i:]

			if len(str) < 4 {
				return i, nil, nil // read more into data and advancing to 'm'
			}

			found := strings.Index("mul(")

		}
	}
}
