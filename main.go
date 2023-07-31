package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func getChangedLines(filename, prevContent string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(prevContent, string(content), false)

	var changedLines []string
	lineCount := 0
	for _, diff := range diffs {
		lines := strings.Split(diff.Text, "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "- ") || strings.HasPrefix(line, "+ ") {
				changedLines = append(changedLines, fmt.Sprintf("%d: %s", lineCount+1, line))
			}
			if !strings.HasPrefix(line, "+ ") {
				lineCount++
			}
		}
	}

	return changedLines, nil
}

func main() {
	filename := "yourfile.txt"
	prevContent := `
		Line 1
		Line 2
		Line 3
		Line 4
	`

	changedLines, err := getChangedLines(filename, prevContent)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if len(changedLines) == 0 {
		fmt.Println("No lines were changed.")
	} else {
		fmt.Println("Changed lines:")
		for _, line := range changedLines {
			fmt.Println(line)
		}
	}
}
