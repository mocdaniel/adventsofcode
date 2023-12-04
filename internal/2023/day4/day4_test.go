package day4

import (
	"testing"
)

func TestPrepareNumbers(t *testing.T) {
	input := "Card 1: 1 2 3 4 5 | 6 7 8 9 10"
	idx, winners, numbers := prepareNumbers(input)
	if !("1 2 3 4 5" == winners && "6 7 8 9 10" == numbers && 1 == idx) {
		t.Errorf("Expected winners to be %v, got %v", []byte("1 2 3 4 5"), winners)
		t.Errorf("Expected numbers to be %v, got %v", []byte("6 7 8 9 10"), numbers)
	}
}

func TestCountMatches(t *testing.T) {
	winners := "1 2 3 4 5"
	numbers := "1 2 3 4 5"
	matches := countMatches(winners, numbers)
	if matches != 5 {
		t.Errorf("Expected 5 matches, got %v", matches)
	}

	winners = "1 2 3 4 5"
	numbers = "1 2 3 4 6 7 9 10"
	matches = countMatches(winners, numbers)
	if matches != 4 {
		t.Errorf("Expected 4 matches, got %v", matches)
	}

	winners = "1 2 3 4 5"
	numbers = "1 2 3 4 6 7 9 10 11 12 13 14 1"
	matches = countMatches(winners, numbers)
	if matches != 5 {
		t.Errorf("Expected 5 matches, got %v", matches)
	}

	winners = "1 2 3 4 10"
	numbers = "6 7 8 9 10"
	matches = countMatches(winners, numbers)
	if matches != 1 {
		t.Errorf("Expected 1 matches, got %v", matches)
	}
}
