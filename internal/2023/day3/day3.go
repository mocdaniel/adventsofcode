package day3

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

func Solve(files ...string) {
	var filePath string
	if len(files) > 0 && len(files[0]) > 0 {
		filePath = files[0]
	} else {
		filePath = "prompts/2023/day3.txt"
	}
	// Read file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := [][]byte{}

	for scanner.Scan() {
		lines = append(lines, []byte(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error scanning file: %v\n", err)
		return
	}

	sum1 := 0
	gears := map[int][]int{}

	for idx, line := range lines {
		prevLine := []byte{}
		if idx > 0 {
			prevLine = lines[idx-1]
		}
		nextLine := []byte{}
		if idx < len(lines)-1 {
			nextLine = lines[idx+1]
		}
		foundNumbers := numberRegex.FindAllIndex(line, -1)
		for _, number := range foundNumbers {
			n, _ := strconv.Atoi(string(line[number[0]:number[1]]))
			if isSchematic(number, &prevLine, &line, &nextLine) {
				sum1 += n
			}
			findGears(number, &prevLine, &line, &nextLine, &gears, idx)
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
