package day1

import "testing"

func TestBuildTwoDigits(t *testing.T) {
	chars := []byte("o1e8t123def")
	n := buildTwoDigits(&chars)
	if n != 13 {
		t.Errorf("Expected 13, got %v", n)
	}

	chars = []byte("123def")
	n = buildTwoDigits(&chars)
	if n != 13 {
		t.Errorf("Expected 13, got %v", n)
	}

	chars = []byte("abc1t2o23s6x")
	n = buildTwoDigits(&chars)
	if n != 16 {
		t.Errorf("Expected 16, got %v", n)
	}

	chars = []byte("4")
	n = buildTwoDigits(&chars)
	if n != 44 {
		t.Errorf("Expected 44, got %v", n)
	}

	chars = []byte("f4r")
	n = buildTwoDigits(&chars)
	if n != 44 {
		t.Errorf("Expected 44, got %v", n)
	}

	chars = []byte("e8t3e")
	n = buildTwoDigits(&chars)
	if n != 83 {
		t.Errorf("Expected 83, got %v", n)
	}

	chars = []byte("42")
	n = buildTwoDigits(&chars)
	if n != 42 {
		t.Errorf("Expected 42, got %v", n)
	}

	chars = []byte("4e8t2")
	n = buildTwoDigits(&chars)
	if n != 42 {
		t.Errorf("Expected 42, got %v", n)
	}

	chars = []byte("48t")
	n = buildTwoDigits(&chars)
	if n != 48 {
		t.Errorf("Expected 48, got %v", n)
	}

	chars = []byte("84")
	n = buildTwoDigits(&chars)
	if n != 84 {
		t.Errorf("Expected 84, got %v", n)
	}
}
