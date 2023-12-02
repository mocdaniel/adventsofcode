package day2

import "testing"

func TestParseProbe(t *testing.T) {
	got := parseProbe("3 blue, 4 red, 13 green")
	want := probe{3, 4, 13}
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = parseProbe("4 red, 13 green, 3 blue")
	want = probe{3, 4, 13}
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = parseProbe("4 red, 13 green")
	want = probe{0, 4, 13}
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = parseProbe("4 red")
	want = probe{0, 4, 0}
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestGetMaxProbe(t *testing.T) {
	got := getMaxProbe([]probe{
		{3, 4, 13},
		{4, 13, 3},
		{0, 4, 13},
		{0, 4, 0},
	})
	want := probe{4, 13, 13}
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
