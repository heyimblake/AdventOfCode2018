package day1

import "testing"

func TestGetResultingFrequency(t *testing.T) {
	expected := 553
	actual := GetResultingFrequency(0, "input.txt")

	if actual != expected {
		t.Fatalf("Expected %d but got %d\n", expected, actual)
	}
}

func TestGetFirstRepeatingFrequency(t *testing.T) {
	expected := 78724
	actual := GetFirstRepeatingFrequency(0, "input.txt", 2)

	if actual != expected {
		t.Fatalf("Expected %d but got %d\n", expected, actual)
	}
}