package day5

import (
	"testing"
)

func TestGetRanges(t *testing.T) {
	source, dest, r := getRange("50 98 2")
	if source != 98 || dest != 50 || r != 2 {
		t.Errorf("getRange() = %v, %v, %v; want 98, 50, 2", dest, source, r)
	}
	if dest != 50 {
		t.Errorf("getRange() = %v, %v, %v; want 98, 50, 2", dest, source, r)
	}
	if r != 2 {
		t.Errorf("getRange() = %v, %v, %v; want 98, 50, 2", dest, source, r)
	}
}
