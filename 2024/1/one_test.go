package one

import (
	"os"
	"testing"
)

func TestDay1Part1(t *testing.T) {
	input, err := os.ReadFile("full_input.txt")

	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}

	got := Day1_Part1(input)
	want := 1341714

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestDay2Part2(t *testing.T) {
	input, err := os.ReadFile("full_input.txt")

	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}

	got := Day1_Part2(input)
	want := 27384707

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
