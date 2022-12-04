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
	var total, ot int

	for _, lines := range fileLines {
		pairs = strings.Split(lines, ",")
		fpr := strings.Split(pairs[0], "-")
		spr := strings.Split(pairs[1], "-")
		sOne, _ := strconv.Atoi(fpr[0])
		eOne, _ := strconv.Atoi(fpr[1])
		sTwo, _ := strconv.Atoi(spr[0])
		eTwo, _ := strconv.Atoi(spr[1])
		fp := generator(sOne, eOne)
		sp := generator(sTwo, eTwo)

		if len(fp) > len(sp) {
			total += subset(sp, fp)
		} else {
			total += subset(fp, sp)
		}

		ot += checkOverlap(sOne, eOne, sTwo, eTwo)
	}

	// PART 1
	fmt.Printf("Total number of pairs are: %d", total)
	fmt.Println()

	// PART 2
	fmt.Printf("Total number of overlapping pairs are: %d", ot)
}

// we can also do a <= c <= d <= b or c <= a <= b <= d
// but golang returns boolean for each compare
// need to implement above solution
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

// [2 8]
// [3 7]
// ! 7 < 2 && ! 8 < 3
// [6 6]
// [4 6]
// ! 6 < 6 && ! 6 < 3
func checkOverlap(ot ...int) (result int) {
	result = 0
	if !(ot[3] < ot[0]) && !(ot[1] < ot[2]) {
		return 1
	}

	return
}

func generator(start, limit int) (result []int) {
	for i := start; i <= limit; i++ {
		result = append(result, i)
	}

	return
}
