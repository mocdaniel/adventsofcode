package day7

import "testing"

func TestGetHandValue(t *testing.T) {
	hand := "32T3K"
	value := getHandValue(hand, false)
	if value != 1 {
		t.Errorf("Expected 1, got %d", value)
	}

	hand = "T55J5"
	value = getHandValue(hand, false)
	if value != 3 {
		t.Errorf("Expected 3, got %d", value)
	}

	hand = "KK677"
	value = getHandValue(hand, false)
	if value != 2 {
		t.Errorf("Expected 2, got %d", value)
	}

	hand = "KTJJT"
	value = getHandValue(hand, false)
	if value != 2 {
		t.Errorf("Expected 2, got %d", value)
	}

	hand = "QQQJA"
	value = getHandValue(hand, false)
	if value != 3 {
		t.Errorf("Expected 3, got %d", value)
	}
}

func TestIsSmaller(t *testing.T) {
	// test different values
	lItem := ListItem{hand: "32431", value: 1, prize: 0}
	rItem := ListItem{hand: "T55J5", value: 3, prize: 0}
	if lItem.isBigger(&rItem, false) {
		t.Errorf("Expected %s to be smaller than %s", lItem.hand, rItem.hand)
	}

	// test same values
	lItem = ListItem{hand: "32431", value: 1, prize: 0}
	rItem = ListItem{hand: "T55J5", value: 1, prize: 0}
	if lItem.isBigger(&rItem, false) {
		t.Errorf("Expected %s to be smaller than %s", lItem.hand, rItem.hand)
	}

	lItem.hand = "T2431"
	if lItem.isBigger(&rItem, false) {
		t.Errorf("Expected %s to be smaller than %s", lItem.hand, rItem.hand)
	}

	lItem.hand = "T5431"
	if lItem.isBigger(&rItem, false) {
		t.Errorf("Expected %s to be smaller than %s", lItem.hand, rItem.hand)
	}

	lItem.hand = "T5531"
	if lItem.isBigger(&rItem, false) {
		t.Errorf("Expected %s to be smaller than %s", lItem.hand, rItem.hand)
	}

	lItem.hand = "T55J1"
	if lItem.isBigger(&rItem, false) {
		t.Errorf("Expected %s to be smaller than %s", lItem.hand, rItem.hand)
	}

	lItem.hand = "T55J5"
	if lItem.isBigger(&rItem, false) {
		t.Errorf("Expected %s to be equal to %s", lItem.hand, rItem.hand)
	}
}
