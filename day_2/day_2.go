package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	total, strategyTotal, point = 0, 0, 0
	delim                       = " "
	rival                       = ""
)

const (
	ROCK, idx1    = 1, 1
	PAPER         = 2
	SCISSOR, DRAW = 3, 3
	WIN           = 6
	LOSE, idx0    = 0, 0
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
		newOp := strings.Split(op, delim)
		total += rockPaperScissor(newOp)
		strategyTotal += strategyGuider(newOp)
	}
	// Part 1
	fmt.Printf("The winning total is: %d", total)
	fmt.Println()
	// Part 2
	fmt.Printf("The strategy guide total is: %d", strategyTotal)
}

func rockPaperScissor(op []string) (result int) {
	switch op[1] {
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

func strategyGuider(op []string) int {
	shapeValue := 0
	switch op[1] {
	case "X":
		point = LOSE
	case "Y":
		point = DRAW
	case "Z":
		point = WIN
	}

	// ROCK
	if op[0] == "A" {
		if point == LOSE {
			shapeValue = SCISSOR
		} else if point == DRAW {
			shapeValue = ROCK
		} else if point == WIN {
			shapeValue = PAPER
		}
	}

	// PAPER
	if op[0] == "B" {
		if point == LOSE {
			shapeValue = ROCK
		} else if point == DRAW {
			shapeValue = PAPER
		} else if point == WIN {
			shapeValue = SCISSOR
		}
	}

	// SCISSOR
	if op[0] == "C" {
		if point == LOSE {
			shapeValue = PAPER
		} else if point == DRAW {
			shapeValue = SCISSOR
		} else if point == WIN {
			shapeValue = ROCK
		}
	}

	return shapeValue + point
}
