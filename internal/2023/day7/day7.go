package day7

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type ListItem struct {
	hand  string
	value int
	prize int
	next  *ListItem
}

var cardValues = map[string]int{
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"J": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}

func getHandValue(hand string, joker bool) int {
	cards := strings.Split(hand, "")
	cardOccurences := make(map[string]int)
	for _, card := range cards {
		cardOccurences[card]++
	}

	// save jokers for later
	jokerCount := 0

	if joker {
		jokerCount = cardOccurences["J"]
		delete(cardOccurences, "J")
	}

	cardValues := make([]int, 0)
	for _, occurence := range cardOccurences {
		cardValues = append(cardValues, occurence)
	}

	slices.Sort(cardValues)

	// different wincon with jokers
	if joker {
		if jokerCount == 5 {
			return 6
		}

		switch cardValues[len(cardValues)-1] {
		case 5:
			return 6
		case 4:
			return 5 + jokerCount
		case 3:
			if len(cardValues) == 3 {
				return 3
			} else if len(cardValues) == 1 {
				return 6 // 3 + two jokers
			} else if cardValues[len(cardValues)-2] == 2 {
				return 4 // full house
			} else {
				return 5 // 3 + 1 joker
			}
		case 2:
			switch len(cardValues) {
			case 1:
				return 6 // 2 + 3 jokers
			case 2:
				if cardValues[len(cardValues)-2] == 2 {
					return 4 // full house
				} else {
					return 5 // 2 + 2 joker
				}
			case 3:
				if jokerCount == 1 {
					return 3
				} else {
					return 2
				}
			case 4:
				return 1
			}
		case 1:
			switch jokerCount {
			case 0: // 5 different cards
				return 0
			case 1: // 4 different cards
				return 1
			case 2: // 3 different cards
				return 3
			case 3: // 2 different cards
				return 5
			case 4: // 1 different card
				return 6
			}
		}
		return 0
	}

	switch cardValues[len(cardValues)-1] {
	case 5:
		return 6
	case 4:
		return 5
	case 3:
		if cardValues[len(cardValues)-2] == 2 {
			return 4
		} else {
			return 3
		}
	case 2:
		if cardValues[len(cardValues)-2] == 2 {
			return 2
		} else {
			return 1
		}
	case 1:
		return 0
	}

	return 0
}

func (l *ListItem) isBigger(cmp *ListItem, joker bool) bool {
	if l.value != cmp.value {
		return l.value > cmp.value
	}

	lValueSplit := strings.Split(l.hand, "")
	cmpValueSplit := strings.Split(cmp.hand, "")

	for idx := range lValueSplit {
		if cardValues[lValueSplit[idx]] != cardValues[cmpValueSplit[idx]] {
			if joker && lValueSplit[idx] == "J" {
				return false
			} else if joker && cmpValueSplit[idx] == "J" {
				return true
			}
			return cardValues[lValueSplit[idx]] > cardValues[cmpValueSplit[idx]]
		}
	}
	return false
}

// inserts the new item into the list at its correct place, returns head
func insert(list *ListItem, item *ListItem, joker bool) *ListItem {
	if list == nil {
		return item
	}

	if !item.isBigger(list, joker) {
		item.next = list
		return item
	}

	if list.next == nil || !item.isBigger(list.next, joker) {
		tmp := list.next
		list.next = item
		item.next = tmp
		return list
	}

	list.next = insert(list.next, item, joker)
	return list
}

func Solve(files ...string) {
	var filePath string
	if len(files) > 0 && len(files[0]) > 0 {
		filePath = files[0]
	} else {
		filePath = "prompts/2023/day7.txt"
	}
	// Read file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	handList1 := (*ListItem)(nil)
	handList2 := (*ListItem)(nil)

	// build list of hands
	for scanner.Scan() {
		line := scanner.Text()
		handStr := strings.Split(line, " ")
		prize, _ := strconv.Atoi(handStr[1])
		hand1 := ListItem{hand: handStr[0], value: getHandValue(handStr[0], false), prize: prize, next: nil}
		handList1 = insert(handList1, &hand1, false)
		hand2 := ListItem{hand: handStr[0], value: getHandValue(handStr[0], true), prize: prize, next: nil}
		handList2 = insert(handList2, &hand2, true)
	}

	// calculate winnings
	winnings1 := 0
	rank := 1
	curHand := handList1
	for curHand != nil {
		winnings1 += curHand.prize * rank
		curHand = curHand.next
		rank++
	}

	winnings2 := 0
	rank = 1
	curHand = handList2
	for curHand != nil {
		winnings2 += curHand.prize * rank
		curHand = curHand.next
		rank++
	}

	fmt.Printf("PART 1: Total winnings are: %d\nPART 2: Total winnings with jokers are %d\n", winnings1, winnings2)

}
