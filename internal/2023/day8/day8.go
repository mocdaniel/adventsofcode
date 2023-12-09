package day8

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/mocdaniel/adventsofcode/internal/lib/files"
	"github.com/mocdaniel/adventsofcode/internal/lib/primes"
)

var regex = regexp.MustCompile(`^([\dA-Z]{3}) = \(([\dA-Z]{3}), ([\dA-Z]{3})\)$`)

type ghostStep struct {
	start    string
	position string
}

func Solve(f ...string) {
	var filePath string
	if len(f) > 0 && len(f[0]) > 0 {
		filePath = f[0]
	} else {
		filePath = "prompts/2023/day8.txt"
	}

	lines, err := files.GetLines(filePath)
	if err != nil {
		fmt.Printf("Error reading prompt: %v\n", err)
		return
	}

	instructions := ""
	mapMap := make(map[string][]string)

	for _, line := range lines {
		matches := regex.FindStringSubmatch(line)
		if len(line) == 0 {
			continue
		}
		if len(matches) != 4 && len(line) != 0 {
			instructions = line
			continue
		}
		if _, ok := mapMap[matches[1]]; !ok {
			mapMap[matches[1]] = []string{matches[2], matches[3]}
		}
	}

	steps := strings.Split(instructions, "")

	stepCount1 := 0
	currentPosition := "AAA"

	for i := 0; i < len(steps); i++ {
		if steps[i] == "L" {
			currentPosition = mapMap[currentPosition][0]
		} else {
			currentPosition = mapMap[currentPosition][1]
		}
		stepCount1++
		if currentPosition == "ZZZ" {
			break
		}
		if len(steps) == i+1 {
			i = -1
		}
	}
	currentPositions := make([]ghostStep, 0)
	for k := range mapMap {
		if k[2] == 'A' {
			currentPositions = append(currentPositions, ghostStep{k, k})
		}
	}

	stepCount2 := 0
	cycleMap := make(map[string]int)
	for i := 0; i < len(steps); i++ {
		// go until each start has been cycled through the end
		if len(cycleMap) == len(currentPositions) {
			break
		}
		for j, gS := range currentPositions {
			if steps[i] == "L" {
				currentPositions[j] = ghostStep{gS.start, mapMap[gS.position][0]}
			} else {
				currentPositions[j] = ghostStep{gS.start, mapMap[gS.position][1]}
			}
			if currentPositions[j].position[2] == 'Z' {
				if _, ok := cycleMap[gS.start]; !ok {
					cycleMap[gS.start] = stepCount2 + 1
					currentPositions = slices.Delete(currentPositions, j, j)
				}
			}
		}
		stepCount2++
		if len(steps) == i+1 {
			i = -1
		}
	}

	cycleList := make([]int, 0, len(cycleMap))
	for _, v := range cycleMap {
		cycleList = append(cycleList, v)
	}

	fmt.Printf("PART 1: Steps required to reach ZZZ: %v\nPART 2: Ghost steps required to reach **Z: %d\n", stepCount1, primes.GetLCM(&cycleList))
}
