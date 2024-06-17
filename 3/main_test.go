package main

import (
	"testing"
)

func TestHasAdjacentSymbolReturnTrue(t *testing.T) {
	testMatrix := [][]rune{
		{46, 46, 25},
		{46, 48, 46},
		{46, 46, 46},
	}

	numberRow := 1
	numberCol := 1
	result := hasAdjacentSymbol(testMatrix, numberRow, numberCol)

	if !result {
		t.Fail()
	}
}

func TestHasAdjacentSymbolReturnFalse(t *testing.T) {
	testMatrix := [][]rune{
		{46, 46, 46},
		{46, 48, 46},
		{46, 46, 46},
	}

	numberRow := 1
	numberCol := 1
	result := hasAdjacentSymbol(testMatrix, numberRow, numberCol)

	if result {
		t.Fail()
	}
}
