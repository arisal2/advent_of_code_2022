package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	filePath := os.Args[1]
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []int
	var totalCalories []int

	for fileScanner.Scan() {
		convertedValue, _ := strconv.Atoi(fileScanner.Text())
		fileLines = append(fileLines, convertedValue)
	}

	readFile.Close()

	total := 0
	for _, line := range fileLines {
		if line == 0 {
			totalCalories = append(totalCalories, total)
			total = 0
		} else {
			total += line
		}
	}

	sort.Ints(totalCalories)

	// Part 1
	fmt.Println("Elf with highest calories:", totalCalories[len(totalCalories)-1])

	// Part 2
	topThreeElves := 0
	for i := len(totalCalories) - 1; i >= len(totalCalories)-3; i-- {
		topThreeElves += totalCalories[i]
	}

	fmt.Println("Total of top three elves:", topThreeElves)
}
