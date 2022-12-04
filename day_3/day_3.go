package main

import (
	fr "advent-of-code/fileReader"
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func main() {
	var (
		mid, end, total, groupTotal = 0, 0, 0, 0
	)

	fileLines := fr.Reader()

	for _, lines := range fileLines {
		mid = (len(lines) / 2)
		end = len(lines)
		total += compareRucksacks(lines[0:mid], lines[mid:end])
	}

	for i := 0; i < len(fileLines)-1; i = i + 3 {
		groupTotal += findCommon(sortString(fileLines[i]), sortString(fileLines[i+1]), sortString(fileLines[i+2]))
	}

	// PART 1
	fmt.Printf("The priority total is: %d", total)
	fmt.Println()

	// PART 2
	fmt.Printf("The group total is: %d", groupTotal)
}

// Need to update this method with above
func compareRucksacks(rsOne, rsTwo string) (result int) {
	for _, first := range rsOne {
		for _, second := range rsTwo {
			if first == second {
				return checkUnicode(first)
			}
		}
	}

	return
}

func checkUnicode(first rune) (result int) {
	if unicode.IsUpper(first) {
		return (int(first) + 27) - 65
	} else {
		return int(first) - 96
	}
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func findCommon(rsOne, rsTwo, rsThree string) (result int) {
	var i, j, k = 0, 0, 0
	for i < len(rsOne) && j < len(rsTwo) && k < len(rsThree) {
		if rsOne[i] == rsTwo[j] && rsTwo[j] == rsThree[k] {
			result = checkUnicode(rune(rsOne[i]))
			i++
			j++
			k++
		} else if rsOne[i] < rsTwo[j] {
			i++
		} else if rsTwo[j] < rsThree[k] {
			j++
		} else {
			k++
		}
	}

	return
}
