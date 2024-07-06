package main

import (
	"testing"
)

func TestSortHandsReturnCorrectOrder(t *testing.T) {
	unsortedHands := []Hand{
		{"A8342", 0},
		{"AAAAK", 0},
		{"AA234", 0},
		{"AAA34", 0},
	}
	sortedHands := sortHands(unsortedHands)

	expectedHands := []Hand{
		{"A8342", 0},
		{"AA234", 0},
		{"AAA34", 0},
		{"AAAAK", 0},
	}

	for i := range expectedHands {
		if sortedHands[i] != expectedHands[i] {
			t.Error("expected", expectedHands, "got", sortedHands)
		}
	}
}

func TestIsHighCardFalse(t *testing.T) {
	hand := Hand{"ABACB", 0}
	result := hand.isHighCard()

	if result {
		t.Fail()
	}
}

func TestIsHighCardTrue(t *testing.T) {
	hand := Hand{"ABCDE", 0}
	result := hand.isHighCard()

	if !result {
		t.Fail()
	}
}

func TestIsOnePairFalse(t *testing.T) {
	hand := Hand{"ABACB", 0}
	result := hand.isOnePair()

	if result {
		t.Fail()
	}
}

func TestIsOnePairTrue(t *testing.T) {
	hand := Hand{"ABBDC", 0}
	result := hand.isOnePair()

	if !result {
		t.Fail()
	}
}

func TestIsTwoPairFalse(t *testing.T) {
	hand := Hand{"ABAAB", 0}
	result := hand.isTwoPair()

	if result {
		t.Fail()
	}
}

func TestIsTwoPairTrue(t *testing.T) {
	hand := Hand{"ABBAC", 0}
	result := hand.isTwoPair()

	if !result {
		t.Fail()
	}
}

func TestIsThreeOfAKindFalse(t *testing.T) {
	hand := Hand{"ABAAB", 0}
	result := hand.isThreeOfAKind()

	if result {
		t.Fail()
	}
}

func TestIsThreeOfAKindTrue(t *testing.T) {
	hand := Hand{"ABAAC", 0}
	result := hand.isThreeOfAKind()

	if !result {
		t.Fail()
	}
}

func TestIsFullHouseFalse(t *testing.T) {
	hand := Hand{"ABAAA", 0}
	result := hand.isFullHouse()

	if result {
		t.Fail()
	}
}

func TestIsFullHouseTrue(t *testing.T) {
	hand := Hand{"ABAAB", 0}
	result := hand.isFullHouse()

	if !result {
		t.Fail()
	}
}

func TestIsFourOfAKindFalse(t *testing.T) {
	hand := Hand{"AAAAA", 0}
	result := hand.isFourOfAKind()

	if result {
		t.Fail()
	}
}

func TestIsFourOfAKindTrue(t *testing.T) {
	hand := Hand{"AAAAB", 0}
	result := hand.isFourOfAKind()

	if !result {
		t.Fail()
	}
}

func TestIsFiveOfAKindTrue(t *testing.T) {
	hand := Hand{"AAAAA", 0}
	result := hand.isFiveOfAKind()

	if !result {
		t.Fail()
	}
}

func TestIsFiveOfAKindFalse(t *testing.T) {
	hand := Hand{"AAAAB", 0}
	result := hand.isFiveOfAKind()

	if result {
		t.Fail()
	}
}
