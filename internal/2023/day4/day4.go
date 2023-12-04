package day4

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func prepareNumbers(line string) (idx int, winners string, numbers string) {
	parts := strings.Split(line, ":")
	idx, _ = strconv.Atoi(regexp.MustCompile(`\s+`).Split(parts[0], -1)[1])
	parts = strings.Split(parts[1], "|")

	return idx, strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
}

func countMatches(w string, numbers string) int {
	winners := regexp.MustCompile(`\s+`).Split(w, -1)

	count := 0
	for _, winner := range winners {
		for _, number := range regexp.MustCompile(`\s+`).Split(numbers, -1) {
			if winner == number {
				count++
			}
		}
	}

	return count
}

func collectCards(lines *[]string, line int) int {
	if line >= len(*lines) {
		return 0
	}

	sum := 0
	_, winners, numbers := prepareNumbers((*lines)[line])
	matches := countMatches(winners, numbers)
	sum += matches
	for i := 1; i <= matches; i++ {
		sum += collectCards(lines, line+i)
	}

	return sum
}

func Solve(files ...string) {
	var filePath string
	if len(files) > 0 && len(files[0]) > 0 {
		filePath = files[0]
	} else {
		filePath = "prompts/2023/day4.txt"
	}
	// Read file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum1 float64 = 0

	lines := []string{}
	cardWins := make(map[int]int)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		line := scanner.Text()
		idx, winners, numbers := prepareNumbers(line)
		matches := countMatches(winners, numbers)
		cardWins[idx] = matches
		if matches != 0 {
			sum1 += math.Pow(2, float64(matches-1))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error scanning file: %v\n", err)
		return
	}

	sum2 := 0
	cardsWon := make([]int, len(lines))

	for idx := range cardsWon {
		fmt.Printf("Line %v\n", idx)
		cardsWon[idx] += 1                   // we 'won' the card itself
		for i := 0; i < cardsWon[idx]; i++ { // for each copy of this card we won
			newWonCards := cardWins[idx+1]      // look up how many other cards this card wins us
			for i := 1; i <= newWonCards; i++ { // for each of those cards
				cardsWon[idx+i] += 1 // add the new cards to our total
			}
		}
	}

	for _, cards := range cardsWon {
		sum2 += cards
	}

	fmt.Printf("PART 1: Sum of all points: %v\nPART 2: Number of cards won: %v\n", int(sum1), sum2)
}
