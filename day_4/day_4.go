package main

import (
	fr "advent-of-code/fileReader"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fileLines := fr.Reader()
	var pairs []string
	var total int

	for _, lines := range fileLines {
		pairs = strings.Split(lines, ",")
		firstPairRange := strings.Split(pairs[0], "-")
		secondPairRange := strings.Split(pairs[1], "-")
		startLimitOne, _ := strconv.Atoi(firstPairRange[0])
		endLimitOne, _ := strconv.Atoi(firstPairRange[1])
		startLimitTwo, _ := strconv.Atoi(secondPairRange[0])
		endLimitTwo, _ := strconv.Atoi(secondPairRange[1])
		firstPair := generator(startLimitOne, endLimitOne)
		secondPair := generator(startLimitTwo, endLimitTwo)

		fmt.Println()
		if len(firstPair) > len(secondPair) {
			total += subset(secondPair, firstPair)
		} else {
			total += subset(firstPair, secondPair)
		}
	}
	fmt.Printf("Total number of pairs are: %d", total)
}

func subset(first, second []int) (result int) {
	set := make(map[int]int)
	for _, value := range second {
		set[value] += 1
	}

	for _, value := range first {
		if count, found := set[value]; !found {
			return 0
		} else if count < 1 {
			return 0
		} else {
			set[value] = count - 1
		}
	}

	return 1
}

func generator(start, limit int) (result []int) {
	for i := start; i <= limit; i++ {
		result = append(result, i)
	}

	return
}
