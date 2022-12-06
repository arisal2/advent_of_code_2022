package main

import (
	fr "advent-of-code/fileReader"
	"fmt"
	"strings"
)

const ()

func main() {
	readLines := fr.Reader()
	// PART 1
	fmt.Printf("First Solution: %v", checker(readLines, 0))
	fmt.Println()

	// PART 2
	fmt.Printf("Second Solution: %v", checker(readLines, 1))
}

func checker(readLines []string, part int) (result int) {
	var i, condition int

	condition = 4
	if part == 1 {
		condition = 14
	}

	for i = 0; i < len(readLines[0]); i++ {
		compareArray := strings.Split(readLines[0][i:i+condition], "")
		cleanedArray := removeDuplicates(compareArray)
		if len(compareArray) == len(cleanedArray) {
			result = i + condition
			break
		}
	}

	return
}

func removeDuplicates(arr []string) (result []string) {
	occurred := map[string]bool{}
	for e := range arr {
		if !occurred[arr[e]] {
			occurred[arr[e]] = true
			result = append(result, arr[e])
		}
	}
	return
}
