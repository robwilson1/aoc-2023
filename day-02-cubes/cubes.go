package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}

	// Run this after `main` returns
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0 // The sum of game IDs that are 'possible'
	maxRed := 12 // The number of red cubes in the bag
	maxGreen := 13 // The number of green cubes in the bag
	maxBlue := 14 // The number of blue cubes in the bag

	gameNumber := 0
	for scanner.Scan() {
		// Get the next line
		line := scanner.Text()
		gameNumber++

		// Is the game possible? Assume true initially
		isPossible := true

		// Remove the prefix of `game n:`
		withoutPrefix := strings.Split(line, ":")[1]

		// Split into cube reveals
		reveals := strings.Split(withoutPrefix, ";")

		//look at the each reveal substring
		for _, reveal := range reveals {
			// Break out early if we have determined the game is not possible
			if !isPossible { break }
			
			//For each reveal does number of blue/red/green cubes exceed the number of possible cubes for that colour?
			for _, cube := range strings.Split(reveal, ",") {
				trimmedCube := strings.TrimSpace(cube)
				num := strings.Split(trimmedCube, " ")[0]
				colour := strings.Split(trimmedCube, " ")[1]

				numInt, parseErr := strconv.Atoi(num)

				if parseErr != nil {
					fmt.Println("Unable to parse number from string:", parseErr)
					return
				}

				if ((colour == "red" && numInt > maxRed) || (colour == "green" && numInt > maxGreen) || (colour == "blue" && numInt > maxBlue)) {
					isPossible = false
					break
				}
			}
		}

		// Add game number to sum if game was possible
		if isPossible {
			sum += gameNumber
		}	
	}

	fmt.Println(sum)
}