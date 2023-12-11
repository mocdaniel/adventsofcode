package day11

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

type galaxy struct {
	x int
	y int
}

func adjustHorizontalExpansion(starMap *[][]string) {
	for idx := 0; idx < len((*starMap)[0]); idx++ {
		isEmpty := true
		for _, row := range *starMap {
			if row[idx] == "#" {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			for jdx, row := range *starMap {
				newRow := slices.Insert(row, idx, ".")
				(*starMap)[jdx] = newRow
			}
			idx++
		}
	}
}

func adjustVerticalExpansion(starMap *[][]string) {
	for idx := 0; idx < len(*starMap); idx++ {
		if !slices.Contains((*starMap)[idx], "#") {
			newRow := make([]string, len((*starMap)[idx]))
			for i := range newRow {
				newRow[i] = "."
			}
			*starMap = slices.Insert(*starMap, idx, newRow)
			idx++
		}
	}
}

func writeHorizontalAdjustement(starMap *[][]string, hA *[]int) {
	for idx := 0; idx < len((*starMap)[0]); idx++ {
		isEmpty := true
		for _, row := range *starMap {
			if row[idx] == "#" {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			*hA = append(*hA, idx)
			idx++
		}
	}
}

func writeVerticalAdjustement(starMap *[][]string, vA *[]int) {
	for idx := 0; idx < len(*starMap); idx++ {
		if !slices.Contains((*starMap)[idx], "#") {
			*vA = append(*vA, idx)
			idx++
		}
	}
}

func shortestPath(galaxy1 galaxy, galaxy2 galaxy, hA *[]int, vA *[]int, a int) int {
	if hA == nil || vA == nil {
		return int(math.Abs(float64(galaxy1.x-galaxy2.x)) + math.Abs(float64(galaxy1.y-galaxy2.y)))
	}

	count := 0
	for x := math.Min(float64(galaxy1.x), float64(galaxy2.x)) + 1; x <= math.Max(float64(galaxy1.x), float64(galaxy2.x)); x++ {
		if slices.Contains(*hA, int(x)) {
			count += a
		} else {
			count++
		}
	}
	for y := math.Min(float64(galaxy1.y), float64(galaxy2.y)) + 1; y <= math.Max(float64(galaxy1.y), float64(galaxy2.y)); y++ {
		if slices.Contains(*vA, int(y)) {
			count += a
		} else {
			count++
		}
	}
	return count
}

func Solve() {
	starMap := make([][]string, 0)

	// create star map
	for _, line := range strings.Split(input, "\n") {
		starMap = append(starMap, strings.Split(strings.TrimSpace(line), ""))
	}

	// note horizontal and vertical expansion for part 2 before expanding for part 1
	horizontalAdjustment := make([]int, 0)
	writeHorizontalAdjustement(&starMap, &horizontalAdjustment)
	verticalAdjustment := make([]int, 0)
	writeVerticalAdjustement(&starMap, &verticalAdjustment)

	// build galaxies without simple expansion
	galaxies := make([]galaxy, 0)
	for y := range starMap {
		for x := range starMap[y] {
			if starMap[y][x] == "#" {
				galaxies = append(galaxies, galaxy{x, y})
			}
		}
	}

	// sum up shortest path between heavily expanded galaxies before expanding starMap for part 1
	sum2 := 0
	for idx := 0; idx < len(galaxies)-1; idx++ {
		for jdx := idx + 1; jdx < len(galaxies); jdx++ {
			sum2 += shortestPath(galaxies[idx], galaxies[jdx], &horizontalAdjustment, &verticalAdjustment, 1000000)
		}
	}

	// adjust for vertical expansion
	adjustVerticalExpansion(&starMap)
	// adjust for horizontal expansion
	adjustHorizontalExpansion(&starMap)

	// create and populate list of galaxies
	galaxies = make([]galaxy, 0)
	for y := range starMap {
		for x := range starMap[y] {
			if starMap[y][x] == "#" {
				galaxies = append(galaxies, galaxy{x, y})
			}
		}
	}

	// sum up shortest path between galaxies
	sum1 := 0
	for idx := 0; idx < len(galaxies)-1; idx++ {
		for jdx := idx + 1; jdx < len(galaxies); jdx++ {
			sum1 += shortestPath(galaxies[idx], galaxies[jdx], nil, nil, 1) // last arg doesn't matter
		}
	}

	fmt.Printf("PART 1: The sum of each galaxy pair's shortest path is %d\nPART 2: The sum of each galax pair's shortest path is %d\n", sum1, sum2)
}
