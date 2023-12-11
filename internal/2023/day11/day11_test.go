package day11

import (
	"strings"
	"testing"
)

var text = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

func TestAdjustVerticalExpansion(t *testing.T) {
	starMap := [][]string{
		{".", ".", "#", ".", ".", "."},
		{".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", "#", "."}}

	adjustVerticalExpansion(&starMap)

	if len(starMap) != 4 {
		t.Errorf("Expected star map to have 4 rows, got %d", len(starMap))
	}
}

func TestAdjustHorizontalExpansion(t *testing.T) {
	starMap := [][]string{
		{".", ".", "#", ".", ".", "."},
		{".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", "#", "."}}

	adjustHorizontalExpansion(&starMap)

	if len(starMap[0]) != 10 {
		t.Errorf("Expected star map to have 10 columns, got %d", len(starMap[0]))
	}
}

func TestShortestPath(t *testing.T) {
	galaxy1 := galaxy{x: 0, y: 1}
	galaxy2 := galaxy{x: 4, y: 6}
	aH := []int{2, 5, 8}
	aV := []int{3, 7}

	path := shortestPath(galaxy1, galaxy2, nil, nil, 1)

	if path != 9 {
		t.Errorf("Expected path to be 9, got %d", path)
	}

	path = shortestPath(galaxy1, galaxy2, &aH, &aV, 10)

	if path != 27 {
		t.Errorf("Expected path to be 27, got %d", path)
	}
}

func TestExampleInput(t *testing.T) {

	starMap := make([][]string, 0)
	for _, line := range strings.Split(text, "\n") {
		starMap = append(starMap, strings.Split(strings.TrimSpace(line), ""))
	}
	adjustHorizontalExpansion(&starMap)
	adjustVerticalExpansion(&starMap)

	galaxies := make([]galaxy, 0)
	for y := range starMap {
		for x := range starMap[y] {
			if starMap[y][x] == "#" {
				galaxies = append(galaxies, galaxy{x, y})
			}
		}
	}

	sum1 := 0
	for idx := 0; idx < len(galaxies)-1; idx++ {
		for jdx := idx + 1; jdx < len(galaxies); jdx++ {
			sum1 += shortestPath(galaxies[idx], galaxies[jdx], nil, nil, 1)
		}
	}

	if sum1 != 374 {
		t.Errorf("Expected sum to be 374, got %d", sum1)
	}
}
