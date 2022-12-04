package main

import (
	fr "advent-of-code/fileReader"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var total, convertedLine int
	var totalCalories []int
	
	fileLines := fr.Reader()

	for _, line := range fileLines {
		fmt.Println(line)
		convertedLine, _ = strconv.Atoi(line)
		if convertedLine != 0 {
			total += convertedLine
			continue
		}
		totalCalories = append(totalCalories, total)
		total = 0
	}

	sort.Ints(totalCalories) // Part 1
	fmt.Println("Elf with highest calories:", totalCalories[len(totalCalories)-1])

	// Part 2
	topThreeElves := 0
	for i := len(totalCalories) - 1; i >= len(totalCalories)-3; i-- {
		topThreeElves += totalCalories[i]
	}

	fmt.Println("Total of top three elves:", topThreeElves)
}
