package main

import (
	"fmt"
	"os"
	"os/exec"
)

func readContentFromLastCommit(filename string) ([]byte, error) {
	cmd := exec.Command("git", "show", "HEAD:"+filename)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	return output, nil
}

func main() {
	filename := "yourfile.txt"

	content, err := readContentFromLastCommit(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Content of the file from the last commit:")
	fmt.Println(string(content))
}
