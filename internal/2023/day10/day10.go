package day10

import (
	"fmt"
	"math"
	"slices"

	"github.com/mocdaniel/adventsofcode/internal/lib/files"
)

type tile struct {
	x    int
	y    int
	form string
}

type border struct {
	a *tile
	b *tile
}

func countContainedTiles(borderTiles *[]tile, borders *[]border, tileMap *[][]tile) int {
	totalCount := 0
	for _, row := range *tileMap {
		for _, t := range row {
			rayCount := 0
			if slices.Contains(*borderTiles, t) {
				continue
			}
			coveredBorders := make([]tile, 0)
			for _, border := range *borders {
				borderA := *border.a
				borderB := *border.b

				if !slices.Contains(coveredBorders, borderA) && borderA.x < t.x && borderA.y == t.y {
					coveredBorders = append(coveredBorders, borderA)
					rayCount++
				} else if !slices.Contains(coveredBorders, borderB) && borderB.x < t.x && borderB.y == t.y {
					coveredBorders = append(coveredBorders, borderB)
					rayCount++
				}
			}

			// deduplicate matches
			for idx, hitBorder := range coveredBorders {
				for idx < len(coveredBorders) {
					if hitBorder.y == coveredBorders[idx].y && (hitBorder.x == coveredBorders[idx].x+1 || hitBorder.x == coveredBorders[idx].x-1) {
						rayCount--
					}
					idx++
				}
			}
			if rayCount%2 == 1 {
				totalCount++
			}
		}
	}
	return totalCount
}

// gets the clockwise previous and next tile for each pipe type, return both
func (t *tile) getNextTile(tileMap *[][]tile) (*tile, *tile) {
	var prev *tile
	var next *tile
	switch t.form {
	case "|":
		if t.y > 0 {
			next = &(*tileMap)[t.y-1][t.x]
		}
		if t.y < len(*tileMap)-1 {
			prev = &(*tileMap)[t.y+1][t.x]
		}
	case "-":
		if t.x > 0 {
			prev = &(*tileMap)[t.y][t.x-1]
		}
		if t.x != len((*tileMap)[t.y])-1 {
			next = &(*tileMap)[t.y][t.x+1]
		}
	case "7":
		if t.x != 0 {
			prev = &(*tileMap)[t.y][t.x-1]
		}
		if t.y != len(*tileMap)-1 {
			next = &(*tileMap)[t.y+1][t.x]
		}
	case "J":
		if t.y != 0 {
			prev = &(*tileMap)[t.y-1][t.x]
		}
		if t.x != 0 {
			next = &(*tileMap)[t.y][t.x-1]
		}
	case "F":
		if t.y != len(*tileMap)-1 {
			prev = &(*tileMap)[t.y+1][t.x]
		}
		if t.x != len((*tileMap)[t.y])-1 {
			next = &(*tileMap)[t.y][t.x+1]
		}
	case "L":
		if t.x != len((*tileMap)[t.y])-1 {
			prev = &(*tileMap)[t.y][t.x+1]
		}
		if t.y != 0 {
			next = &(*tileMap)[t.y-1][t.x]
		}
	}

	return prev, next
}

// checks if tile t1 is clockwise connected to tile t2
func (t1 *tile) isConnected(t2 *tile) bool {
	// common cases can be short-circuited
	if t1.form == "|" && t2.form == "|" {
		return t1.x == t2.x && (t1.y-1 == t2.y || t1.y+1 == t2.y)
	}

	if t1.form == "-" && t2.form == "-" {
		return t1.y == t2.y && (t1.x-1 == t2.x || t1.x+1 == t2.x)
	}

	switch t1.form {
	case "|":
		// either 7 or F top, or L or J bottom
		if (t2.form == "7" || t2.form == "F") && (t1.x == t2.x && t1.y-1 == t2.y) {
			return true
		}
		if (t2.form == "L" || t2.form == "J") && (t1.x == t2.x && t1.y+1 == t2.y) {
			return true
		}
	case "-":
		// either F or L left, or 7 or J right
		if (t2.form == "F" || t2.form == "L") && (t1.x-1 == t2.x && t1.y == t2.y) {
			return true
		}
		if (t2.form == "7" || t2.form == "J") && (t1.x+1 == t2.x && t1.y == t2.y) {
			return true
		}
	case "7":
		// either F, L, or - left or J, L, or | bottom
		if (t2.form == "F" || t2.form == "L" || t2.form == "-") && (t1.x-1 == t2.x && t1.y == t2.y) {
			return true
		}
		if (t2.form == "J" || t2.form == "L" || t2.form == "|") && (t1.x == t2.x && t1.y+1 == t2.y) {
			return true
		}
	case "J":
		// either 7, F, or | top, or F, L, or - left
		if (t2.form == "7" || t2.form == "F" || t2.form == "|") && (t1.x == t2.x && t1.y-1 == t2.y) {
			return true
		}
		if (t2.form == "F" || t2.form == "L" || t2.form == "-") && (t1.x-1 == t2.x && t1.y == t2.y) {
			return true
		}
	case "F":
		// either 7, J, or - right, or J, L, or | bottom
		if (t2.form == "7" || t2.form == "J" || t2.form == "-") && (t1.x+1 == t2.x && t1.y == t2.y) {
			return true
		}
		if (t2.form == "J" || t2.form == "L" || t2.form == "|") && (t1.x == t2.x && t1.y+1 == t2.y) {
			return true
		}
	case "L":
		// either F, 7, or | top, or J, 7, or - right
		if (t2.form == "F" || t2.form == "7" || t2.form == "|") && (t1.x == t2.x && t1.y-1 == t2.y) {
			return true
		}
		if (t2.form == "J" || t2.form == "7" || t2.form == "-") && (t1.x+1 == t2.x && t1.y == t2.y) {
			return true
		}
	}
	return false
}

func Solve(f ...string) {
	var filePath string
	if len(f) > 0 && len(f[0]) > 0 {
		filePath = f[0]
	} else {
		filePath = "prompts/2023/day10.txt"
	}

	lines, err := files.GetLines(filePath)
	if err != nil {
		fmt.Printf("Error reading prompt: %v\n", err)
		return
	}

	// Create a map and save start tile
	tileMap := make([][]tile, 0, len(lines))
	var start *tile
	for y, line := range lines {
		row := make([]tile, 0, len(line))
		for x, c := range line {
			row = append(row, tile{x: x, y: y, form: string(c)})
			if string(c) == "S" {
				start = &row[x]
			}
		}
		tileMap = append(tileMap, row)
	}

	variations := []string{"|"}

	// Simulate all possible paths from S
	maxRoundtrip := 0
	var startPipe string
	var finalBorderTiles []tile
	var finalBorders []border
	finalFieldCount := 0
	for _, v := range variations {
		*start = tile{x: start.x, y: start.y, form: string(v)}
		lastTile := start
		var nextTile *tile
		borderTiles := []tile{*start}
		_, curTile := start.getNextTile(&tileMap) // we only need to go forward
		borders := make([]border, 0)
		openMap := make(map[int]int)
		fieldCount := 0
		if curTile != nil && lastTile.y != curTile.y {
			borders = append(borders, border{a: lastTile, b: curTile})
		}
		roundtrip := 1
		openMap[lastTile.y] = lastTile.x // populate with start field

		for curTile != start {
			// if there's no next tile, we can stop early
			if curTile == nil {
				break
			}

			// get the possibly connected tiles
			a, b := curTile.getNextTile(&tileMap)

			// check where we're coming from
			if a != nil && a != lastTile {
				nextTile = a
			} else if b != nil && b != lastTile {
				nextTile = b
			}

			if nextTile != nil && curTile.isConnected(nextTile) {
				roundtrip++

				// remember open fields

				// if we moved on a horizontal border
				if lastTile.y == curTile.y {
					if lastTile.x < curTile.x {
						// if we moved right before, and we go down, we can drop
						if nextTile.y > curTile.y {
							delete(openMap, curTile.y)
						} else {
							// if we move right before, and we go up, we need to set a new value
							openMap[curTile.y] = curTile.x
						}
					} else {
						// if we moved left before, and we go up, we can drop
						if nextTile.y < curTile.y {
							delete(openMap, lastTile.y)
						} else {
							// if we move left before, and we go down, we need to set a new value
							openMap[lastTile.y] = curTile.x
						}
					}
				} else {
					// if there's a value already, we can add all fields within this range
					if _, ok := openMap[curTile.y]; ok {
						fieldCount += int(math.Abs(float64(curTile.x-openMap[curTile.y]))) - 1
						delete(openMap, curTile.y)
					} else {
						// otherwise we start a new range
						openMap[curTile.y] = curTile.x
					}
				}

				// remember tile as border tile if it's not already
				if !slices.Contains(borderTiles, *curTile) {
					borderTiles = append(borderTiles, *curTile)
				}

				// remember border if it's vertical
				if curTile.y != nextTile.y {
					borders = append(borders, border{a: curTile, b: nextTile})
				}

				lastTile = curTile
				curTile = nextTile
			} else {
				break
			}
		}

		// if we broke out of the loop because we found a roundtrip
		if curTile == start {
			if roundtrip > maxRoundtrip {
				maxRoundtrip = roundtrip
				finalBorderTiles = borderTiles
				finalBorders = borders
				startPipe = v
				finalFieldCount = fieldCount
			}
		}
	}

	*start = tile{x: start.x, y: start.y, form: startPipe}

	containedTiles := countContainedTiles(&finalBorderTiles, &finalBorders, &tileMap)

	fmt.Printf("PART 1: Steps to the farthest point of the loop: %d\nPART 2: Tiles enclosed by the loop: %d %d\nVersion: %v", int(math.Ceil(float64(maxRoundtrip)/2)), finalFieldCount, containedTiles, startPipe)
}
