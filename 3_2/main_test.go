package main

import (
	"testing"
)

func TestHasAdjacentGearReturnTrueAndCorrentPosition(t *testing.T) {
	testMatrix := [][]rune{
		{46, 46, 42},
		{46, 48, 46},
		{46, 46, 46},
	}

	numberRow := 1
	numberCol := 1
	ok, pos := hasAdjacentGear(testMatrix, numberRow, numberCol)

	if !ok {
		t.Fail()
	}

	if !(pos.x == 0 && pos.y == 2) {
		t.Fail()
	}
}

func TestHasAdjacentGearReturnFalse(t *testing.T) {
	testMatrix := [][]rune{
		{46, 46, 46},
		{46, 48, 46},
		{46, 46, 46},
	}

	numberRow := 1
	numberCol := 1
	ok, _ := hasAdjacentGear(testMatrix, numberRow, numberCol)

	if ok {
		t.Fail()
	}
}
