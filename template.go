package main

import (
	"bufio"
	"fmt"
	"os"
)

// Template to read and process input text files
func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}

	// Run this after `main` returns
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}