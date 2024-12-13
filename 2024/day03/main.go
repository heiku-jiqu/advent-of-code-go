package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const INPUTFILEPATH string = "./2024/day03/input.txt"

func main() {
	f, err := os.Open(INPUTFILEPATH)
	if err != nil {
		log.Panicln("Failed to open file: ", err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	part1ans := 0

	stop := false
	for {
		if stop {
			break
		}
		_, rerr := reader.ReadString('m')
		if rerr == io.EOF {
			stop = true
		}
		if rerr != nil && rerr != io.EOF {
			log.Print("no delim found in current string")
			continue
		}
		// need to peek 11 forwards after 'm': ul(123,456)
		peekbytes, perr := reader.Peek(11)
		if rerr == io.EOF {
			stop = true
		}
		if perr != nil && perr != io.EOF {
			log.Panic("unexpected error peeking: ", perr)
		}

		log.Print("peekbytes is: ", peekbytes, len(peekbytes))
		peekString := string(peekbytes)
		log.Print("peek string is: ", peekString)
		ubLoc := strings.Index(peekString, "ul")
		if ubLoc == -1 || ubLoc != 0 {
			continue
		}
		openParenLoc := strings.Index(peekString, "(")
		if openParenLoc == -1 || openParenLoc != 2 {
			continue
		}
		commaLoc := strings.Index(peekString, ",")
		if commaLoc == -1 || commaLoc < openParenLoc {
			continue
		}
		closeParenLoc := strings.Index(peekString, ")")
		if closeParenLoc == -1 || closeParenLoc < commaLoc {
			continue
		}
		first, err := strconv.Atoi(peekString[openParenLoc+1 : commaLoc])
		if err != nil || first > 999 {
			continue
		}

		second, err := strconv.Atoi(peekString[commaLoc+1 : closeParenLoc])
		if err != nil || second > 999 {
			continue
		}

		part1ans += first * second

		if rerr == io.EOF || perr == io.EOF {
			log.Print("finish reading")
			break
		}
	}

	log.Print("Day 3 Part 1 answer is: ", part1ans)
}
