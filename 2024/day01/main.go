package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const INPUTPATH string = "./2024/day01/input.txt"

func main() {
	f, err := os.Open(INPUTPATH)
	if err != nil {
		log.Fatalf("failed to open input file: ", err)
	}
	defer f.Close()

	var list1 []int
	var list2 []int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lineslice := strings.Fields(line)

		first, err := strconv.ParseInt(lineslice[0], 10, 32)
		if err != nil {
			log.Fatal("failed to parse first int:", err)
		}
		list1 = append(list1, int(first))

		second, err := strconv.ParseInt(lineslice[1], 10, 32)
		if err != nil {
			log.Fatal("failed to parse second int:", err)
		}
		list2 = append(list2, int(second))
	}

	log.Print("list1:", list1)
	log.Print("list2:", list2)

	sort.Sort(sort.Reverse(sort.IntSlice(list1)))
	sort.Sort(sort.Reverse(sort.IntSlice(list2)))

	log.Print("sort finished")
	log.Print("list1:", list1)
	log.Print("list2:", list2)

	sum := 0
	for i := range list1 {
		diff := list2[i] - list1[i]
		sum += abs(diff)
	}

	log.Print("The answer is part 1 is: ", sum)

	m := make(map[int]int)
	for _, key := range list1 {
		if _, ok := m[key]; !ok {
			m[key] = 0
		}
	}

	for _, key := range list2 {
		if _, ok := m[key]; ok {
			m[key] += 1
		}
	}

	var day2ans int
	for key, val := range m {
		day2ans += key * val
	}
	log.Print("The answer for part 2 is: ", day2ans)
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
