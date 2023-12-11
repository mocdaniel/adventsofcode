package day9

import "testing"

func TestGetNext(t *testing.T) {
	sequence := []int{0, 3, 6, 9, 12, 15}
	prev, next := getPrevAndNext(&sequence)
	if next != 18 || prev != -3 {
		t.Errorf("Expected -3, 18, got %v, %v", prev, next)
	}

	sequence = []int{1, 3, 6, 10, 15, 21}
	prev, next = getPrevAndNext(&sequence)
	if next != 28 || prev != 0 {
		t.Errorf("Expected 0, 28, got %v, %v", prev, next)
	}

	sequence = []int{10, 13, 16, 21, 30, 45}
	prev, next = getPrevAndNext(&sequence)
	if next != 68 || prev != 5 {
		t.Errorf("Expected 5, 68, got %v, %v", prev, next)
	}

	sequence = []int{1, 2}
	prev, next = getPrevAndNext(&sequence)
	if next != 3 || prev != 0 {
		t.Errorf("Expected 0, 3, got %v, %v", prev, next)
	}

	sequence = []int{1, 3, -5, 6, 9}

}
