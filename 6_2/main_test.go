package main

import (
	"testing"
)

func TestGetDistanceTraveled(t *testing.T) {
	var tests = []struct {
		name          string
		inputMilis    int
		inputRaceTime int
		want          int
	}{
		{"press 0 should return 0", 0, 7, 0},
		{"press 2 should return 10", 2, 7, 10},
		{"press 5 should return 10", 5, 7, 10},
		{"press 7 should return 0", 7, 7, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getDistancedTraveled(tt.inputMilis, tt.inputRaceTime)
			if result != tt.want {
				t.Errorf("got %d want %d", result, tt.want)
			}
		})
	}
}
