package day13

import "testing"

func TestFindHorizontalReflections(t *testing.T) {
	i := []string{"#...##..#",
		"#....#..#",
		"..##..###",
		"#####.##.",
		"#####.##.",
		"..##..###",
		"#....#..#"}
	got := findHorizontalReflection(i)
	if got[0] != 4 {
		t.Errorf("findHorizontalReflections(i) = %d; want 4", got)
	}

}

func TestFindVerticalReflections(t *testing.T) {
	i := []string{"#.##..##.",
		"..#.##.#.",
		"##......#",
		"##......#",
		"..#.##.#.",
		"..##..##.",
		"#.#.##.#."}
	got := findVerticalReflection(i)
	if got[0] != 4 {
		t.Errorf("findVerticalReflections(i) = %d; want 4", got)
	}

	i = []string{"#...##",
		".#...."}
	got = findVerticalReflection(i)
	if got[0] != 4 {
		t.Errorf("findVerticalReflections(i) = %d; want 4", got)
	}

	i = []string{"##...##",
		"......."}
	got = findVerticalReflection(i)
	if got[0] != 0 {
		t.Errorf("findVerticalReflections(i) = %d; want 0", got)
	}
}

func TestExample1(t *testing.T) {
	i := [][]string{
		{"#.##..##.",
			"..#.##.#.",
			"##......#",
			"##......#",
			"..#.##.#.",
			"..##..##.",
			"#.#.##.#."},
		{"#...##..#",
			"#....#..#",
			"..##..###",
			"#####.##.",
			"#####.##.",
			"..##..###",
			"#....#..#"},
	}

	sum1 := 0
	for _, m := range i {
		hM := findHorizontalReflection(m)
		if len(hM) != 0 {
			sum1 += hM[0] * 100
		}
		vM := findVerticalReflection(m)
		if len(vM) != 0 {
			sum1 += vM[0] + 1
		}
	}
	if sum1 != 405 {
		t.Errorf("sum1 = %d; want 405", sum1)
	}
}

func TestExample2(t *testing.T) {
	i := [][]string{
		{"#.##..##.",
			"..#.##.#.",
			"##......#",
			"##......#",
			"..#.##.#.",
			"..##..##.",
			"#.#.##.#."},
		{"#...##..#",
			"#....#..#",
			"..##..###",
			"#####.##.",
			"#####.##.",
			"..##..###",
			"#....#..#"},
	}

	sum2 := 0
	for _, m := range i {
		sum2 += findSmudgedReflection(m)
	}
	if sum2 != 400 {
		t.Errorf("sum2 = %d; want 400", sum2)
	}
}
