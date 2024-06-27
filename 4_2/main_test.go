package main

import (
	"reflect"
	"testing"
)

func TestGetWinningNumbersReturnCorrectNumbers(t *testing.T) {
	line := "Card   1:  9 39 27  5 |  9 80 29 20 54 "
	result := getWinningNumbers(line)

	expected := []string{"9", "39", "27", "5"}

	if !reflect.DeepEqual(result, expected) {
		t.Fail()
	}
}

func TestGetMyNumbersReturnCorrectNumbers(t *testing.T) {
	line := "Card   1:  9 39 27  5 |  9 80 29 5"
	result := getMyNumbers(line)

	expected := []string{"9", "80", "29", "5"}

	if !reflect.DeepEqual(result, expected) {
		t.Fail()
	}
}

func TestGetNumberOfMatchingNumbersReturn4(t *testing.T) {
	line := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"

	result := getNumberOfMatchingNumbers(getWinningNumbers(line), getMyNumbers(line))

	if result != 4 {
		t.Errorf("Expected 4 got %d", result)
	}
}
