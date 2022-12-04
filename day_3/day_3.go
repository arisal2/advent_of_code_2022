package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

var (
	mid, end, total = 0, 0, 0
)

func main() {
	//filePath := os.Args[1]
	filePath := "day_3.txt"
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	for _, lines := range fileLines {
		mid = (len(lines) / 2) - 1
		end = len(lines) - 1
		total += compareRucksacks(lines[0:mid+1], lines[mid+1:end+1])
	}

	// PART 1
	fmt.Printf("The priority total is: %d", total)

	// PART 2
	//fmt.Printf("The group total is: %d", groupTotal)
}

func compareRucksacks(ruckSackFirst, ruckSackSecond string) (result int) {
	for _, first := range ruckSackFirst {
		for _, second := range ruckSackSecond {
			if first == second {
				if unicode.IsUpper(first) {
					result = (int(first) + 27) - 65
				} else {
					result = int(first) - 96
				}
				break
			}
		}
		
	}
	return
}
