package day9

import "testing"

func TestGetNext(t *testing.T) {
	sequence := []int{0, 3, 6, 9, 12, 15}
	next := getNext(&sequence)
	if next != 18 {
		t.Errorf("Expected 18, got %v", next)
	}

	sequence = []int{1, 3, 6, 10, 15, 21}
	next = getNext(&sequence)
	if next != 28 {
		t.Errorf("Expected 28, got %v", next)
	}

	sequence = []int{10, 13, 16, 21, 30, 45}
	next = getNext(&sequence)
	if next != 68 {
		t.Errorf("Expected 68, got %v", next)
	}

	sequence = []int{1, 2}
	next = getNext(&sequence)
	if next != 3 {
		t.Errorf("Expected 3, got %v", next)
	}

	sequence = []int{1, 3, -5, 6, 9}

}
