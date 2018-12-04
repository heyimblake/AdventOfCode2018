package day2

import "testing"

func TestGenerateChecksum(t *testing.T) {
	expected := 8892
	actual := GenerateChecksum("Day2.txt")

	if actual != expected {
		t.Fatalf("Expected %d but got %d\n", expected, actual)
	}
}

func TestGetCommonBoxIDLetters(t *testing.T) {
	expected := "zihwtxagifpbsnwleydukjmqv"
	actual := GetCommonBoxIDLetters("Day2.txt")

	if actual != expected {
		t.Fatalf("Expected \"%s\" but got \"%s\"\n", expected, actual)
	}
}