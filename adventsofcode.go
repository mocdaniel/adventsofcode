package main

import (
	"flag"

	"github.com/mocdaniel/adventsofcode/internal/2023/day1"
	"github.com/mocdaniel/adventsofcode/internal/2023/day10"
	"github.com/mocdaniel/adventsofcode/internal/2023/day11"
	"github.com/mocdaniel/adventsofcode/internal/2023/day2"
	"github.com/mocdaniel/adventsofcode/internal/2023/day3"
	"github.com/mocdaniel/adventsofcode/internal/2023/day4"
	"github.com/mocdaniel/adventsofcode/internal/2023/day5"
	"github.com/mocdaniel/adventsofcode/internal/2023/day6"
	"github.com/mocdaniel/adventsofcode/internal/2023/day7"
	"github.com/mocdaniel/adventsofcode/internal/2023/day8"
	"github.com/mocdaniel/adventsofcode/internal/2023/day9"
)

func main() {
	year := flag.Int("year", 2023, "The year")
	day := flag.Int("day", 4, "The day")
	filePath := flag.String("f", "", "The file path")

	flag.Parse()

	switch *year {
	case 2023:
		switch *day {
		case 1:
			day1.Solve(*filePath)
		case 2:
			day2.Solve(*filePath)
		case 3:
			day3.Solve(*filePath)
		case 4:
			day4.Solve(*filePath)
		case 5:
			day5.Solve(*filePath)
		case 6:
			day6.Solve()
		case 7:
			day7.Solve(*filePath)
		case 8:
			day8.Solve(*filePath)
		case 9:
			day9.Solve(*filePath)
		case 10:
			day10.Solve(*filePath)
		case 11:
			day11.Solve()
		}
	}
}
