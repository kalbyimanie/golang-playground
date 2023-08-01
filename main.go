package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func readNewLinesFromLastCommit(filename string) ([]string, error) {
	cmd := exec.Command("git", "diff", "HEAD^", "--", filename)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var newLines []string
	lines := strings.Split(string(output), "\n")

	for _, line := range lines {
		if strings.HasPrefix(line, "+") && !strings.HasPrefix(line, "+++") {
			newLines = append(newLines, strings.TrimPrefix(line, "+"))
		}
	}

	return newLines, nil
}

func main() {
	filename := "yourfile.txt"

	lines, err := readNewLinesFromLastCommit(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if len(lines) == 0 {
		fmt.Println("No new lines were added to the file in the last commit.")
	} else {
		fmt.Println("New lines added to the file in the last commit:")
		for _, line := range lines {
			fmt.Println(line)
		}
	}
}
