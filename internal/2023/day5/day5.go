package day5

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type seedRange struct {
	start  int
	length int
	mapped bool
}

type mapRange struct {
	start    int
	length   int
	newStart int
}

// maps all seeds in seedList to their new locations according to mapping
// and writes them to seedListTemp.
// This happens **in place**.
func applyMapping(seedList *[]seedRange, mapping *mapRange) {
	// if mapping is nil, we need to map all remaining, unmapped seeds
	if mapping == nil {
		for idx := range *seedList {
			(*seedList)[idx].mapped = false
		}
		return
	}

	offset := mapping.newStart - mapping.start

	for idx, seed := range *seedList {
		if seed.mapped {
			continue
		}
		// if seedRange is entirely inside of the mapping (incl. shared borders), map it
		if seed.start >= mapping.start && seed.start+seed.length <= mapping.start+mapping.length {
			(*seedList)[idx].start = seed.start + offset
			(*seedList)[idx].mapped = true
			// if seed range runs into the mapping from the left, map it and append a new seed range
		} else if seed.start < mapping.start && seed.start+seed.length > mapping.start {
			// prepare new seed range to be appended
			seedRangeRemainderLeft := seedRange{seed.start, mapping.start - seed.start, false}
			// map overlapping part
			(*seedList)[idx].start = mapping.start + offset
			(*seedList)[idx].length = seed.length - (mapping.start - seed.start)
			(*seedList)[idx].mapped = true
			*seedList = append(*seedList, seedRangeRemainderLeft)
			// if seed range starts in the mapping and runs out of it, map it and append a new seed range
		} else if seed.start >= mapping.start && seed.start < mapping.start+mapping.length && seed.start+seed.length > mapping.start+mapping.length {
			// prepare new seed range to be appended
			seedRangeRemainderRight := seedRange{mapping.start + mapping.length, seed.start + seed.length - mapping.start - mapping.length, false}
			// map overlapping part
			(*seedList)[idx].start = seed.start + offset
			(*seedList)[idx].length = mapping.start + mapping.length - seed.start
			(*seedList)[idx].mapped = true
			*seedList = append(*seedList, seedRangeRemainderRight)
			// if mapping range is entirely inside of the seed range (excl. shared borders), map it and append two new seed ranges
		} else if seed.start < mapping.start && seed.start+seed.length > mapping.start+mapping.length {
			// prepare new seed ranges to be appended
			seedRangeRemainderLeft := seedRange{seed.start, mapping.start - seed.start, false}
			seedRangeRemainderRight := seedRange{mapping.start + mapping.length, seed.length - (2*mapping.start + mapping.length - seed.start), false}
			// map overlapping part
			(*seedList)[idx].start = mapping.start + offset
			(*seedList)[idx].length = mapping.length
			(*seedList)[idx].mapped = true
			*seedList = append(*seedList, seedRangeRemainderLeft)
			*seedList = append(*seedList, seedRangeRemainderRight)
		}
		// otherwise there is no overlap
	}
}

func sortRanges(a, b seedRange) int {
	return cmp.Compare(a.start, b.start)
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

	seedList1 := make([]seedRange, len(seeds))
	seedList2 := make([]seedRange, len(seeds)/2)

	// Prepare seed ranges
	for idx, s := range seeds {
		seed, _ := strconv.Atoi(s)
		seedList1[idx] = seedRange{seed, 1, false}

		if idx/2 < len(seedList2) && idx%2 == 0 {
			length, _ := strconv.Atoi(seeds[idx+1])
			seedList2[idx/2] = seedRange{seed, length, false}
		}
	}

	// Prepare map ranges
	mappings := make([][]mapRange, len(mapList))
	for idx, m := range mapList {
		mapRanges := make([]mapRange, len(m))
		for jdx, mapping := range m {
			mappingSplit := strings.Split(mapping, " ")
			start, _ := strconv.Atoi(mappingSplit[1])
			length, _ := strconv.Atoi(mappingSplit[2])
			newStart, _ := strconv.Atoi(mappingSplit[0])
			mapRanges[jdx] = mapRange{start, length, newStart}
		}
		mappings[idx] = mapRanges
	}

	// iterate over each mapping stage
	for _, mappings := range mappings {
		// iterate over each mapping
		for _, mapping := range mappings {
			// map all seed ranges in place
			applyMapping(&seedList1, &mapping)
			applyMapping(&seedList2, &mapping)
		}
		applyMapping(&seedList1, nil)
		applyMapping(&seedList2, nil)
	}

	answer1 := slices.MinFunc(seedList1, sortRanges)
	answer2 := slices.MinFunc(seedList2, sortRanges)

	fmt.Printf("Len: %d\n", len(seedList2))
	fmt.Printf("PART 1: The lowest corresponding location number is %d\nPART 2: The new lowest corresponding location number is %d\n", answer1.start, answer2.start)
}
