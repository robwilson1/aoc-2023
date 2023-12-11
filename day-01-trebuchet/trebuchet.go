package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// This function takes in a byte, converts it to a rune, and checks if it is '0' - '9'
func isDigit(char rune) bool {
    return unicode.IsDigit(char)
}

func getWordMap() map[string]string {
	return map[string]string{
		"one": "1",
		"two": "2",
		"three": "3",
		"four": "4",
		"five": "5",
		"six": "6",
		"seven": "7",
		"eight": "8",
		"nine": "9",
	}
}

func main() {
	wordMap := getWordMap()

	file, err := os.Open("./input.txt")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// gets called after main returns
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum int

	for scanner.Scan() {
		line := scanner.Text()
		var front, back string
		var frontWord, backWord string

		// Reverse the indicies to begin with to aid in loop logic
		frontIdx := len(line) - 1
		backIdx := 0

		// Set the front/back indicies based on digits
		for idx, char := range line {
			if isDigit(char) {
				if idx < frontIdx {
					frontIdx = idx
				}

				if idx > backIdx {
					backIdx = idx
				}
			}
		}

		// Override front/back indicies if a word substring appears before/after current indicies
		for key := range wordMap {
			wordIdx := strings.Index(line, key)
			if wordIdx != -1 {
				if wordIdx < frontIdx {
					frontIdx = wordIdx
					frontWord = key
				}

				lastWordIdx := strings.LastIndex(line, key)
				if lastWordIdx != -1 && lastWordIdx > backIdx {
					backIdx = lastWordIdx
					backWord = key
				}
			}
		}

		// TODO: handle here if neither front or frontIdx found digit/word
		// TODO: handle here if neither back or backIdx found digit/word

		// If frontWord/backWord is not empty, then use that
		if (frontWord != "") {
			front = wordMap[frontWord]
		} else {
			front = string(line[frontIdx])
		}

		if (backWord != "") {
			back = wordMap[backWord]
		} else {
			back = string(line[backIdx])
		}

		lineSum, lineSumErr := strconv.Atoi(front + back)

		if (lineSumErr != nil) {
			fmt.Println("Error summing the line:", lineSumErr)
			return
		}

		sum += lineSum
	}

	scanErr := scanner.Err()

	if scanErr != nil {
		fmt.Println("Error reading from file:", scanErr)
		return
	}

	fmt.Println(sum)
}
