package day6

import "fmt"

func getWinningStrategies(duration int, record int) int {
	results := 0
	for i := 0; i <= duration/2; i++ {
		if i*(duration-i) > record {
			results += 2
		}
	}
	if duration%2 == 0 {
		results--
	}
	return results
}

func Solve() {

	inputs := [][]int{{40, 215}, {70, 1051}, {98, 2147}, {79, 1005}}
	// inputs := [][]int{{7, 9}, {15, 40}, {30, 200}}  // example

	answer1 := 1
	for _, input := range inputs {
		answer1 *= getWinningStrategies(input[0], input[1])
	}

	fmt.Printf("PART 1: The multiplied number of ways I can win the record is %d\nPART 2: The number of ways the long race can be won in is %d\n", answer1, getWinningStrategies(40709879, 215105121471005))
}
