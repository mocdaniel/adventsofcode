package main

import (
	"flag"

	"github.com/mocdaniel/adventsofcode/internal/2023/day1"
	"github.com/mocdaniel/adventsofcode/internal/2023/day2"
)

func main() {
	year := flag.Int("year", 2023, "The year")
	day := flag.String("day", "day1", "The day")
	filePath := flag.String("f", "", "The file path")

	flag.Parse()

	switch *year {
	case 2023:
		switch *day {
		case "day1":
			day1.Solve(*filePath)
		case "day2":
			day2.Solve(*filePath)
		}
	}
}
