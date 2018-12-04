package day1

import (
	"github.com/heyimblake/AdventOfCode-2018/util"
	"strconv"
)

/*
	Part 1
 */
func GetResultingFrequency(startingFrequency int, filePath string) int {
	fileScanner, file := util.GetScanner(filePath)
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

	// Get contents of file
	lines, size := util.GetFileContents(filePath)

	// Keep on looping until we find a result and exit
	for {
		for i := 0; i < size; i++ {
			// Convert to int and add to frequency count
			freq, _ := strconv.Atoi((*lines)[i])
			result += freq

			// Update Map
			dupes[result]++

			// Check for duplicate
			if dupes[result] == timesRepeated {
				return result
			}
		}
	}
}

