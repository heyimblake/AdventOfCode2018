package day2

import (
	"github.com/heyimblake/AdventOfCode-2018/util"
	"strings"
)

/*
	Part 1
 */
func GenerateChecksum(fileName string) int {
	// Get Scanner and Files (returned as double pointers)
	scanner, file := util.GetScanner(fileName)

	// Ensure file will be closed
	defer (*file).Close()

	// Variables to calculate checksum
	doubles := 0
	triples := 0

	// Scan each line
	for (*scanner).Scan() {
		// Get line of text
		line := (*scanner).Text()

		// Create map of characters with the number of times it appears
		var chars = make(map[string]int)

		// Misc variables used in the below loop
		checkForDouble := true
		checkForTriple := true

		// Go through each character in the string and add it to the map
		for i := 0; i < len(line); i++ {
			currentChar := string(line[i])
			chars[currentChar]++
		}

		// Check for doubles and triples
		for _, v := range chars {
			if !checkForDouble && !checkForTriple {
				break
			}

			// Check for double characters
			if checkForDouble && v == 2 {
				checkForDouble = false
				doubles++
			}

			// Check for triple characters
			if checkForTriple && v == 3 {
				checkForTriple = false
				triples++
			}
		}
	}

	return doubles * triples
}

/*
	Part 2
 */
func GetCommonBoxIDLetters(fileName string) string {
	var common string

	// Get lines from file
	lines, numlines := util.GetFileContents(fileName)

	outer:
	for i := 0; i < numlines; i++ {
		for j := 0; j < numlines; j++ {
			// Don't compare the same line!
			if i == j {
				continue
			}

			// If they aren't the same length, then it doesn't qualify to check
			length := len((*lines)[i])

			if length != len((*lines)[j])  {
				continue
			}

			diff := getCommonLetters((*lines)[i], (*lines)[j])

			// We found a match!
			if len(diff) == length - 1 {
				common = diff
				break outer // Using a label to break out of both loops
			}

		}
	}
	
	return common
}

/*
	Returns a string of common letters
 */
func getCommonLetters(str1 string, str2 string) string {
	max1 := len(str1)
	max2 := len(str2)
	var max int
	var stringBuilder strings.Builder

	// Ensure that we don't go out of bounds of a string
	if max1 <= max2 {
		max = max1
	} else {
		max = max2
	}

	for i := 0; i < max; i++ {
		if str1[i] == str2[i] {
			stringBuilder.WriteString(string(str1[i]))
		}
	}

	return stringBuilder.String()
}