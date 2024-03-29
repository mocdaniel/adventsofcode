package day1

import (
	"fmt"
	"strings"

	"github.com/mocdaniel/adventsofcode/internal/lib/files"
)

func replaceDigitStrings(line string) string {
	replaced := strings.Replace(line, "one", "o1e", -1)
	replaced = strings.Replace(replaced, "two", "t2o", -1)
	replaced = strings.Replace(replaced, "three", "t3e", -1)
	replaced = strings.Replace(replaced, "four", "f4r", -1)
	replaced = strings.Replace(replaced, "five", "f5e", -1)
	replaced = strings.Replace(replaced, "six", "s6x", -1)
	replaced = strings.Replace(replaced, "seven", "s7n", -1)
	replaced = strings.Replace(replaced, "eight", "e8t", -1)
	replaced = strings.Replace(replaced, "nine", "n9e", -1)
	return replaced
}

func getFirstInt(chars *[]byte) int {
	for _, char := range *chars {
		if int(char) < 58 {
			return int(char - '0')
		}
	}
	return 0
}

func getLastInt(chars *[]byte) int {
	for idx := len(*chars) - 1; idx >= 0; idx-- {
		if int((*chars)[idx]) < 58 {
			return int((*chars)[idx] - '0')
		}
	}
	return 0
}

func buildTwoDigits(chars *[]byte) int {
	firstInt := getFirstInt(chars)
	lastInt := getLastInt(chars)
	return firstInt*10 + lastInt
}

func Solve(f ...string) {
	var filePath string
	if len(f) > 0 && len(f[0]) > 0 {
		filePath = f[0]
	} else {
		filePath = "prompts/2023/day1.txt"
	}

	lines, err := files.GetLines(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	sum1 := 0
	sum2 := 0
	for _, line := range lines {
		line1 := []byte(line)
		line2 := []byte(replaceDigitStrings(line))
		sum1 += buildTwoDigits(&line1)
		sum2 += buildTwoDigits(&line2)
	}

	fmt.Printf("PART 1: Sum of all calibration values: %v\nPART 2: Sum of all calibration values: %v\n", sum1, sum2)
}
