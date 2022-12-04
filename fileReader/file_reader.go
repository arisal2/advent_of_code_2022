package fileReader

import (
	"bufio"
	"fmt"
	"os"
)

func Reader() (fileLines []string) {
	if len(os.Args) == 1 {
		fmt.Println("error: please enter the input file name")
		os.Exit(3)
	}

	filePath := os.Args[1]
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()
	return
}
