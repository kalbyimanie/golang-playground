package main

import (
	"bufio"
	"fmt"
	"os"
)

func readContentFromLine(filename string, lineNumber int) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var content []string
	currentLine := 1

	// Read the file until the desired line number is reached
	for scanner.Scan() {
		if currentLine >= lineNumber {
			content = append(content, scanner.Text())
		}
		currentLine++

		if currentLine > lineNumber {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return content, nil
}

func main() {
	filename := "yourfile.txt"
	lineNumber := 5

	content, err := readContentFromLine(filename, lineNumber)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if len(content) == 0 {
		fmt.Println("Line number is out of range.")
	} else {
		fmt.Println("Content from line", lineNumber, "onwards:")
		for _, line := range content {
			fmt.Println(line)
		}
	}
}
