package day13

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func findSmudgedReflection(m []string) int {
	ohM := findHorizontalReflection(m)
	ovM := findVerticalReflection(m)
	for jdx, row := range m {
		for idx := range row {
			smudge := row[idx]
			rowSplit := []byte(row)
			if smudge == '#' {
				rowSplit[idx] = '.'
			} else {
				rowSplit[idx] = '#'
			}
			m[jdx] = string(rowSplit)
			hM := findHorizontalReflection(m)
			for _, mirror := range hM {
				if !slices.Contains(ohM, mirror) {
					return mirror * 100
				}
			}
			vM := findVerticalReflection(m)
			for _, mirror := range vM {
				if !slices.Contains(ovM, mirror) {
					return mirror + 1
				}
			}
			rowSplit[idx] = smudge
			m[jdx] = string(rowSplit)
		}
	}
	return 0
}

func findHorizontalReflection(m []string) []int {

	// swap dimensions of the map
	swappedMap := make([]string, len(m[0]))
	for row := range swappedMap {
		for idx := len(m) - 1; idx >= 0; idx-- {
			swappedMap[row] += string(m[idx][row])
		}
	}

	// since we transform the map to search vertically, we need to translate the found index back to the original map
	transformedMirror := findVerticalReflection(swappedMap)
	if len(transformedMirror) == 0 {
		return make([]int, 0)
	} else {
		for idx, mirror := range transformedMirror {
			transformedMirror[idx] = len(m) - 1 - mirror
		}
	}
	return transformedMirror
}

func findVerticalReflection(m []string) []int {
	leftSides := make([]int, 0)

	// collect all possible vertical reflections by iterating over line 0
	for idx := range m[0] {
		if idx == len(m[0])-1 {
			break
		}
		// get the max length of the reflection to compare
		scope := int(math.Min(float64(idx)+1, float64(len(m[0])-1-idx)))
		leftSide := []byte(m[0][idx-scope+1 : idx+1])
		rightSide := []byte(m[0][idx+1 : idx+scope+1])
		slices.Reverse(rightSide)

		if slices.Equal(leftSide, rightSide) {
			leftSides = append(leftSides, idx)
		}

	}

	// check if the vertical reflection is valid for all rows
	results := make([]int, 0)
	for _, mirrorLine := range leftSides {

		for idx, row := range m {
			if idx == 0 {
				continue
			}
			// get the max length of the reflection to compare
			scope := int(math.Min(float64(mirrorLine)+1, float64(len(row)-1-mirrorLine)))
			leftSide := []byte(row[mirrorLine-scope+1 : mirrorLine+1])
			rightSide := []byte(row[mirrorLine+1 : mirrorLine+scope+1])
			slices.Reverse(rightSide)

			if !slices.Equal(leftSide, rightSide) {
				break
			}
			if idx == len(m)-1 {
				results = append(results, mirrorLine)
			}
		}
	}

	return results
}

func Solve() {

	formations := strings.Split(input, "\n\n")

	maps := make([][]string, 0, len(formations))

	for _, formation := range formations {
		rows := strings.Split(formation, "\n")
		maps = append(maps, rows)
	}

	sum1 := 0
	for _, m := range maps {
		hM := findHorizontalReflection(m)
		if len(hM) != 0 {
			sum1 += hM[0] * 100
		}
		vM := findVerticalReflection(m)
		if len(vM) != 0 {
			sum1 += vM[0] + 1
		}
	}

	sum2 := 0
	for _, m := range maps {
		sum2 += findSmudgedReflection(m)
	}

	fmt.Printf("PART 1: The sum of all notes is %d\nPART 2: The sum of all notes for smudged mirrors is %d\n", sum1, sum2)
}
