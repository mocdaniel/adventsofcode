package day9

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mocdaniel/adventsofcode/internal/lib/files"
)

func getPrevAndNext(sequence *[]int) (prev int, next int) {
	onlyZeroes := true
	for _, i := range *sequence {
		if i != 0 {
			onlyZeroes = false
			break
		}
	}

	if onlyZeroes {
		return 0, 0
	}

	nextSequence := make([]int, 0, len(*sequence)-1)
	for i := 0; i < len(*sequence)-1; i++ {
		nextSequence = append(nextSequence, int((*sequence)[i+1]-(*sequence)[i]))
	}

	lastNumber := (*sequence)[len(*sequence)-1]
	firstNumber := (*sequence)[0]

	nextPrev, nextNext := getPrevAndNext(&nextSequence)
	return firstNumber - nextPrev, lastNumber + nextNext
}

func Solve(f ...string) {
	var filePath string
	if len(f) > 0 && len(f[0]) > 0 {
		filePath = f[0]
	} else {
		filePath = "prompts/2023/day9.txt"
	}

	lines, err := files.GetLines(filePath)
	if err != nil {
		fmt.Printf("Error reading prompt: %v\n", err)
		return
	}

	sum1 := 0
	sum2 := 0
	for _, line := range lines {
		stringSequence := strings.Split(line, " ")
		sequence := make([]int, 0, len(stringSequence))
		for _, s := range stringSequence {
			i, _ := strconv.Atoi(s)
			sequence = append(sequence, i)
		}
		prev, next := getPrevAndNext(&sequence)
		sum1 += next
		sum2 += prev
	}

	fmt.Printf("PART 1: Sum of all next numbers is %d\nPART 2: Sum of all previous numbers is %d", sum1, sum2)
}
