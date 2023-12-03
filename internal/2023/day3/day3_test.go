package day3

import "testing"

func TestIsSchematic(t *testing.T) {
	// no gear
	prevLine := []byte(".....")
	line := []byte("..3..")
	nextLine := []byte(".....")
	got := isSchematic([]int{2, 3}, &prevLine, &line, &nextLine)
	want := false
	if got != want {
		t.Errorf("isSchematic() = %v, want %v", got, want)
	}

	// beginning of line
	prevLine = []byte(".....")
	line = []byte("3+...")
	nextLine = []byte(".....")
	got = isSchematic([]int{0, 1}, &prevLine, &line, &nextLine)
	want = true
	if got != want {
		t.Errorf("isSchematic() = %v, want %v", got, want)
	}

	// end of line
	prevLine = []byte(".....")
	line = []byte("...+3")
	nextLine = []byte(".....")
	got = isSchematic([]int{4, 5}, &prevLine, &line, &nextLine)
	want = true
	if got != want {
		t.Errorf("isSchematic() = %v, want %v", got, want)
	}
}
