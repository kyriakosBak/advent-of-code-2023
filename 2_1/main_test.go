package main

import (
	"testing"
)

func TestLineParsing(t *testing.T) {
	game := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"

	res, err := parseGame(game)
	if err != nil {
		t.Error(err)
	}

	if res.MaxRed != 4 {
		t.Errorf("Red: Expected 4 and got %d", res.MaxRed)
	}
	if res.MaxGreen != 2 {
		t.Errorf("Green: Expected 2 and got %d", res.MaxGreen)
	}
	if res.MaxBlue != 6 {
		t.Errorf("Blue: Expected 6 and got %d", res.MaxBlue)
	}
}
