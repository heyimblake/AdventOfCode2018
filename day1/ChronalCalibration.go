package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
	Part 1
 */
func GetResultingFrequency(startingFrequency int, filePath string) int {
	fileScanner, file := getScanner(filePath)
	result := startingFrequency

	// Ensure file will be closed when the function is finished
	defer (*file).Close()

	// Scan each line
	for (*fileScanner).Scan() {
		line := (*fileScanner).Text()
		freq, _ := strconv.Atoi(line)
		result += freq
	}

	return result
}

/*
	Part 2
 */
func GetFirstRepeatingFrequency(startingFrequency int, filePath string, timesRepeated int) int {
	var dupes = make(map[int]int)// Key: Frequency, Value: # times appeared

	// Add starting frequency to map
	dupes[startingFrequency] = 1

	// Current Frequency
	result := startingFrequency

	for {
		fileScanner, file := getScanner(filePath)

		// Scan each line
		for (*fileScanner).Scan() {
			line := (*fileScanner).Text()
			freq, _ := strconv.Atoi(line)
			result += freq

			// Update Map
			dupes[result]++

			// Check for first duplicate
			if dupes[result] == timesRepeated {
				(*file).Close()
				return result
			}
		}

		// Close file
		(*file).Close()
	}
}

func check(e error) {
	if e != nil {
		fmt.Printf("ERROR: Cannot open file. Does it exist?\n")
		os.Exit(1)
	}
}

func getScanner(filePath string) (**bufio.Scanner, **os.File) {
	// Open Data File
	file, err := os.Open(filePath)

	// Check if we can continue
	check(err)

	// Create scanner
	scanner := bufio.NewScanner(file)

	// Return scanner and file double pointers
	return &scanner, &file
}