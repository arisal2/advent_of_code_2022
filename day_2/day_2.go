package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	total, point, firstTotal, strategyTotal = 0, 0, 0, 0
	delim                                   = " "
)

const (
	ROCK, idx1    = 1, 1
	PAPER         = 2
	SCISSOR, DRAW = 3, 3
	WIN           = 6
	LOSE, idx0    = 0, 0
	STRATEGY      = "strategy"
	NORMAL        = "normal"
)

func main() {
	filePath := os.Args[1]
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

	for _, op := range fileLines {
		firstTotal += rockPaperScissor(strings.Split(op, delim), NORMAL)
		strategyTotal += rockPaperScissor(strings.Split(op, delim), STRATEGY)
	}
	// Part 1
	fmt.Printf("The winning total is: %d", firstTotal)

	// Part 2
	fmt.Printf("The strategy guide total is: %d", strategyTotal)
}

func swapElements(op []string, idx0, idx1 int) []string {
	temp := op[idx0]
	op[idx0] = op[idx1]
	op[idx1] = temp
	return op
}

func rockPaperScissor(op []string, condition string) (result int) {
	var rival string
	swapCondition := op[1]
	if condition == STRATEGY {
		swapElements(op, idx0, idx1)
		swapCondition = op[0]
	}

	switch swapCondition {
	case "X":
		rival = "A"
		point = ROCK
	case "Y":
		rival = "B"
		point = PAPER
	case "Z":
		rival = "C"
		point = SCISSOR
	}

	if op[0] == rival {
		return DRAW + point
	} else if (op[0] == "A" && rival == "C") || (op[0] == "C" && rival == "B") || (op[0] == "B" && rival == "A") {
		return LOSE + point
	} else if (rival == "A" && op[0] == "C") || (rival == "C" && op[0] == "B") || (rival == "B" && op[0] == "A") {
		return WIN + point
	} else {
		return 0
	}
}
