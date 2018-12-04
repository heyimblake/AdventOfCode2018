package util

import (
	"testing"
)

func TestGetFileContents(t *testing.T) {
	actual, _ := GetFileContents("../day1//Day1.txt")

	expected := "-10"
	checkStrings(&t, 5, &actual, &expected)

	expected = "+4"
	checkStrings(&t, 0, &actual, &expected)

	expected = "+6"
	checkStrings(&t, 20, &actual, &expected)

	expected = "-79113"
	checkStrings(&t, 1023, &actual, &expected)
}

func checkStrings(t **testing.T, i int, actual **[]string, expected *string) {
	if (**actual)[i] != *expected {
		(*t).Fatalf("Expected \"%s\" but got \"%s\"\n", *expected, (**actual)[i])
	}
}
