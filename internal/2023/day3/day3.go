package day3

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/mocdaniel/adventsofcode/internal/lib/files"
)

var adjacentSymbolRegex = regexp.MustCompile(`[^\d\.]`)
var numberRegex = regexp.MustCompile(`\d+`)
var gearRegex = regexp.MustCompile(`\*`)

func findGears(number []int, prevLine *[]byte, line *[]byte, nextLine *[]byte, gears *map[int][]int, lineNumber int) {
	lines := [][]byte{*prevLine, *line, *nextLine}

	for idx, currentLine := range lines {
		if len(currentLine) > 0 {
			foundGears := gearRegex.FindAllIndex(currentLine, -1)
			for _, gear := range foundGears {
				if gear[0] < number[0]-1 || gear[0] > number[1] {
					continue
				}
				gearNumber := (lineNumber+idx-1)*len(currentLine) + gear[0]
				number, _ := strconv.Atoi(string((*line)[number[0]:number[1]]))
				_, exists := (*gears)[gearNumber]
				if exists {
					(*gears)[gearNumber] = append((*gears)[gearNumber], number)
				} else {
					(*gears)[gearNumber] = []int{number}
				}
			}
		}
	}
}

func isSchematic(number []int, prevLine *[]byte, line *[]byte, nextLine *[]byte) bool {
	lines := [][]byte{*prevLine, *line, *nextLine}

	for _, line := range lines {
		if len(line) > 0 {
			slice := line[max(0, number[0]-1):min(len(line), number[1]+1)]
			if adjacentSymbolRegex.Match(slice) {
				return true
			}
		}
	}
	return false
}

func Solve(f ...string) {
	var filePath string
	if len(f) > 0 && len(f[0]) > 0 {
		filePath = f[0]
	} else {
		filePath = "prompts/2023/day3.txt"
	}

	lines, err := files.GetLines(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	sum1 := 0
	gears := map[int][]int{}

	for idx, line := range lines {
		prevLine := []byte{}
		if idx > 0 {
			prevLine = []byte(lines[idx-1])
		}
		nextLine := []byte{}
		if idx < len(lines)-1 {
			nextLine = []byte(lines[idx+1])
		}
		foundNumbers := numberRegex.FindAllIndex([]byte(line), -1)
		for _, number := range foundNumbers {
			n, _ := strconv.Atoi(string(line[number[0]:number[1]]))
			l := []byte(line)
			if isSchematic(number, &prevLine, &l, &nextLine) {
				sum1 += n
			}
			findGears(number, &prevLine, &l, &nextLine, &gears, idx)
		}
	}

	sum2 := 0
	for _, gear := range gears {
		if len(gear) == 2 {
			sum2 += gear[0] * gear[1]
		}
	}

	fmt.Printf("PART 1: The sum of all part numbers in the schematic is %v\nPART 2: The sum of all gear rations in the schematic is %v\n", sum1, sum2)
}
