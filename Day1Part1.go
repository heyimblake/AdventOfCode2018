package AdventOfCode_2018

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Printf("ERROR: Provide the input file name.\n")
		os.Exit(1)
	}

	fmt.Printf("The resulting frequency is %d\n", getResultingFrequency(0, args[1]))
}
*/

func getResultingFrequency(startingFrequency int, filePath string) int {
	fileScannerDoublePointer, fileDoublePointer := getScanner(filePath)
	scanner := *fileScannerDoublePointer
	file := *fileDoublePointer
	result := startingFrequency

	// Ensure file will be closed when the function is finished
	defer file.Close()

	// Scan each line
	for scanner.Scan() {
		line := scanner.Text()
		freq, _ := strconv.Atoi(line)
		//fmt.Printf("Adding %d (New Total %d)\n", freq, result + freq)
		result += freq
	}

	return result
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