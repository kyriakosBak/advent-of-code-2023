package main

import (
	"testing"
)

func TestGetSeedLocationReturnCorrectLocation(t *testing.T) {
	mapStructA := MapStruct{
		name: "seed-to-B",
		from: "seed",
		to:   "B",
		lines: []string{
			"50 97 5",
			"52 50 48",
		},
	}
	mapStructA.from = "seed"
	mapStructB := MapStruct{
		name:  mapStructA.from + "-to-location",
		from:  mapStructA.to,
		to:    "location",
		lines: []string{"3 51 4"},
	}

	mapStruct := make(map[string]MapStruct)
	mapStruct[mapStructA.name] = mapStructA
	mapStruct[mapStructB.name] = mapStructB
	seed := 99
	expectedDestination := 4
	result := getSeedLocation(seed, mapStruct)

	if result != expectedDestination {
		t.Errorf("Expected %d, got %d", expectedDestination, result)
	}
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
		name: "A-to-B",
		from: "A",
		to:   "B",
		lines: []string{
			"50 98 2",
			"52 50 48",
		},
	}
}
