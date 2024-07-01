package main

import (
	"testing"
)

func TestGetSeedLocationReturnCorrectLocation(t *testing.T) {
	t.Fail()
}

func TestGetDestinationReturnsMappedValue(t *testing.T) {
	mapStruct := getDummyMapStruct()
	seed := 99
	expectedDestination := 51
	result := getDestination(seed, mapStruct)

	if result != expectedDestination {
		t.Errorf("Expected %d, got %d", expectedDestination, result)
	}
}

func TestGetDestinationReturnsDefaultValue(t *testing.T) {
	mapStruct := getDummyMapStruct()
	seed := 1001
	expectedDestination := 1001
	result := getDestination(seed, mapStruct)

	if result != expectedDestination {
		t.Errorf("Expected %d, got %d", expectedDestination, result)
	}
}

func getDummyMapStruct() MapStruct {
	return MapStruct{
		name: "from-to",
		from: "from",
		to:   "to",
		lines: []string{
			"50 98 2",
			"52 50 48",
		},
	}
}
