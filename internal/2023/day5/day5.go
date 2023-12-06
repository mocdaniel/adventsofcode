package day5

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getRange(line string) (dest int, source int, r int) {
	parts := strings.Split(line, " ")
	sourceStart, _ := strconv.Atoi(parts[0])
	destStart, _ := strconv.Atoi(parts[1])
	rangeLength, _ := strconv.Atoi(parts[2])

	return destStart, sourceStart, rangeLength
}

// id is for attribute identification
func mapRanges(dest int, source int, seeds *map[int][]int, r int, idx int) {
	for seed := range *seeds {
		if source != -1 {
			if (*seeds)[seed][idx] >= source && (*seeds)[seed][idx] < source+r {
				(*seeds)[seed][idx+1] = dest + (*seeds)[seed][idx] - source
			}
		} else if (*seeds)[seed][idx+1] == 0 {
			(*seeds)[seed][idx+1] = (*seeds)[seed][idx]
		}
		// if len(*source) != 0 { // test if we are doing default-mappings
		// 	sIdx := slices.Index(*source, seed[idx]) // test if seed is in source range
		// 	// value to be mapped
		// 	if sIdx != -1 {
		// 		seed[idx+1] = (*dest)[sIdx] // if seed is in source range, map to dest range
		// 	}
		// } else if seed[idx+1] == 0 {
		// 	seed[idx+1] = seed[idx] // if no destination, map to source range
		// }
	}
}

func Solve(files ...string) {
	var filePath string
	if len(files) > 0 && len(files[0]) > 0 {
		filePath = files[0]
	} else {
		filePath = "prompts/2023/day5.txt"
	}
	// Read file
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", file)
		return
	}

	maps := strings.Split(string(file), "\n\n")

	seeds := strings.Split(strings.TrimSpace(strings.Split(maps[0], ":")[1]), " ")
	seeSoi := strings.Split(strings.Split(maps[1], ":\n")[1], "\n")
	soiFer := strings.Split(strings.Split(maps[2], ":\n")[1], "\n")
	ferWat := strings.Split(strings.Split(maps[3], ":\n")[1], "\n")
	watLig := strings.Split(strings.Split(maps[4], ":\n")[1], "\n")
	ligTem := strings.Split(strings.Split(maps[5], ":\n")[1], "\n")
	temHum := strings.Split(strings.Split(maps[6], ":\n")[1], "\n")
	humLoc := strings.Split(strings.Split(maps[7], ":\n")[1], "\n")
	mapList := [][]string{seeSoi, soiFer, ferWat, watLig, ligTem, temHum, humLoc}

	// initialize seeds
	seedList1 := make(map[int][]int, len(seeds))
	for i := range seeds {
		value, _ := strconv.Atoi(seeds[i])
		seedList1[value] = []int{value, 0, 0, 0, 0, 0, 0, 0}
	}

	seedList2 := make(map[int][]int, 0)
	for i := 0; i < len(seeds); i += 2 {
		seedValue, _ := strconv.Atoi(seeds[i])

		rLength, _ := strconv.Atoi(seeds[i+1])
		for j := 1; j < rLength; j++ {
			if _, ok := seedList1[seedValue+j]; !ok {
				seedList2[seedValue+j] = []int{seedValue + j, 0, 0, 0, 0, 0, 0, 0}
			}
		}
	}

	// iterate over map lists
	for i := range mapList {
		fmt.Printf("Mapping %v\n", i)
		// iterate over range mappings per map List
		for _, line := range mapList[i] {
			dest, source, r := getRange(line)
			mapRanges(source, dest, &seedList1, r, i)
			mapRanges(source, dest, &seedList2, r, i)
		}
		// iterate once more for default-mappings
		mapRanges(-1, -1, &seedList1, -1, i)
		mapRanges(-1, -1, &seedList2, -1, i)
	}

	// find the lowest location number for seedList1
	answer1 := math.MaxInt64
	for _, seed := range seedList1 {
		if seed[7] < answer1 {
			answer1 = seed[7]
		}
	}

	// find the lowest location number for seedList2
	answer2 := math.MaxInt64
	for _, seed := range seedList2 {
		if seed[7] < answer2 {
			answer2 = seed[7]
		}
	}

	fmt.Printf("PART 1: The lowest corresponding location number is %v\nPART 2: The new lowest corresponding location number is %v\n", answer1, answer2)
}
