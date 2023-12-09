package lists

import "testing"

func TestCount(t *testing.T) {
	l := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := 1
	count := Count(n, &l)
	if count != 1 {
		t.Errorf("Count(%v, %v) = %v, want %v", n, l, count, 1)
	}

	l = []int{1, 2, 3, 4, 5, 6, 7, 8, 1}
	count = Count(n, &l)
	if count != 2 {
		t.Errorf("Count(%v, %v) = %v, want %v", n, l, count, 2)
	}

	l = []int{2, 3, 4, 5, 6, 7, 8, 9}
	n = 1
	count = Count(n, &l)
	if count != 0 {
		t.Errorf("Count(%v, %v) = %v, want %v", n, l, count, 0)
	}
}
