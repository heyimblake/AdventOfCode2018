package util

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Printf("ERROR: Cannot open file. Does it exist?\n")
		os.Exit(1)
	}
}

func GetScanner(filePath string) (**bufio.Scanner, **os.File) {
	// Open Data File
	file, err := os.Open(filePath)

	// Check if we can continue
	check(err)

	// Create scanner
	scanner := bufio.NewScanner(file)

	// Return scanner and file double pointers
	return &scanner, &file
}

func GetFileContents(filePath string) (*[]string, int) {
	result := make([]string, 0)
	total := 0

	// Get scanner and file
	scanner, file := GetScanner(filePath)

	// Close when this ends
	defer (*file).Close()

	for (*scanner).Scan()  {
		result = append(result, (*scanner).Text())
		total++
	}

	return &result, total
}