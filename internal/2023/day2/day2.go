package day2

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type probe struct {
	blue  int
	red   int
	green int
}

var redRegex = regexp.MustCompile(`\d+ red`)
var blueRegex = regexp.MustCompile(`\d+ blue`)
var greenRegex = regexp.MustCompile(`\d+ green`)

func getMaxProbe(probes []probe) probe {
	maxProbe := probe{0, 0, 0}
	for _, probe := range probes {
		if probe.blue > maxProbe.blue {
			maxProbe.blue = probe.blue
		}
		if probe.red > maxProbe.red {
			maxProbe.red = probe.red
		}
		if probe.green > maxProbe.green {
			maxProbe.green = probe.green
		}
	}
	return maxProbe
}

func parseGameId(s string) int {
	gameIdStr := strings.Split(s, " ")[1]
	gameId, _ := strconv.Atoi(gameIdStr)
	return gameId
}

func parseProbe(s string) probe {
	red := 0
	redMatch := redRegex.Find([]byte(s))
	if redMatch != nil {
		redStr := strings.Split(string(redMatch), " ")[0]
		red, _ = strconv.Atoi(redStr)
	}

	blue := 0
	blueMatch := blueRegex.Find([]byte(s))
	if blueMatch != nil {
		blueStr := strings.Split(string(blueMatch), " ")[0]
		blue, _ = strconv.Atoi(blueStr)
	}

	green := 0
	greenMatch := greenRegex.Find([]byte(s))
	if greenMatch != nil {
		greenStr := strings.Split(string(greenMatch), " ")[0]
		green, _ = strconv.Atoi(greenStr)
	}

	return probe{blue, red, green}
}

func Solve(files ...string) {
	var filePath string
	if len(files) > 0 && len(files[0]) > 0 {
		filePath = files[0]
	} else {
		filePath = "prompts/2023/day2.txt"
	}
	// Read file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum1 := 0
	collection1 := probe{14, 12, 13}
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, ":")

		probes := []probe{}
		for _, probeStr := range strings.Split(splitLine[1], ";") {
			probes = append(probes, parseProbe(probeStr))
		}

		maxProbe := getMaxProbe(probes)

		if maxProbe.blue <= collection1.blue && maxProbe.red <= collection1.red && maxProbe.green <= collection1.green {
			sum1 += parseGameId(splitLine[0])
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error scanning file: %v\n", err)
		return
	}

	fmt.Printf("Sum1: %v\n", sum1)
}
