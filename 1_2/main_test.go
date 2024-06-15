package main

import (
	"testing"
)

func TestLineParsing(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  int
	}{
		{"1as2 should be 12", "1as2", 12},
		{"sd2as3 should be 23", "sd2as3", 23},
		{"5sd2sdllas should be 52", "5sd2sdllas", 52},
		{"85sr2sd1s should be 81", "85sr2sd1s", 81},
		{"sr2sds should be 22", "sr2sds", 22},
		{"onetwo should be 12", "onetwo", 12},
		{"one2 should be 12", "one2", 12},
		{"3nine should be 39", "3nine", 39},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := stringToNumber(tt.input)
			if err != nil {
				t.Error(err)
			}
			if result != tt.want {
				t.Errorf("got %d want %d", result, tt.want)
			}
		})
	}
}

func TestTextNumberExists(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  int
	}{
		{"one should return 1", "one", 1},
		{"aone should return -1", "aone", -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := checkTextNumberExistsAtTheStartOfText(tt.input)
			if result != tt.want {
				t.Errorf("got %d want %d", result, tt.want)
			}
		})
	}
}
