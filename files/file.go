package file

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFromFile(path string) []string {
	readFile, err := os.Open(path)
	var input []string

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		input = append(input, fileScanner.Text())
	}

	readFile.Close()
	return input
}
