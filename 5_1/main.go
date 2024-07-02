package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	seeds, dataMaps := getAllMapsAndSeeds()

	lowestLocation := math.MaxInt
	for _, seed := range seeds {
		seedLocation := getSeedLocation(seed, dataMaps)
		if seedLocation < lowestLocation {
			lowestLocation = seedLocation
		}
	}

	fmt.Println(lowestLocation)
}

func getSeedLocation(seed int, dataMaps map[string]MapStruct) int {
	sourceVal := seed
	currentDataMap := MapStruct{to: "seed"}
	for {
		if len(dataMaps) == 0 {
			break
		}

		// Find correct map
		for _, val := range dataMaps {
			if val.from == currentDataMap.to {
				currentDataMap = val
				break
			}
		}

		sourceVal = getDestination(sourceVal, currentDataMap)

		if currentDataMap.to == "location" {
			break
		}
	}
	return sourceVal
}

func getDestination(sourceVal int, currentDataMap MapStruct) int {
	result := sourceVal
	for _, line := range currentDataMap.lines {
		splitted := strings.Split(line, " ")
		lineSourceStart, err := strconv.Atoi(splitted[1])
		if err != nil {
			log.Fatal(err)
		}
		lineSourceOffset, err := strconv.Atoi(splitted[2])
		if err != nil {
			log.Fatal(err)
		}
		lineSourceEnd := lineSourceStart + lineSourceOffset - 1
		if sourceVal >= lineSourceStart && sourceVal <= lineSourceEnd {
			lineDestStart, err := strconv.Atoi(splitted[0])
			if err != nil {
				log.Fatal(err)
			}
			result = lineDestStart + (sourceVal - lineSourceStart)
			break
		}

	}
	return result
}

func getAllMapsAndSeeds() ([]int, map[string]MapStruct) {
	filePath := "input.txt"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	seeds := []int{}
	currentMapName := ""
	dataMaps := make(map[string]MapStruct)
	for scanner.Scan() {
		line := scanner.Text()

		// Only for first line
		if len(seeds) == 0 {
			seedsStr := strings.Split(strings.TrimSpace(strings.Split(line, ":")[1]), " ")
			for _, s := range seedsStr {
				num, err := strconv.Atoi(s)
				if err != nil {
					log.Fatal(err)
				}
				seeds = append(seeds, num)
			}
			continue
		}

		// check if map start
		if strings.Contains(line, ":") {
			currentMapName = strings.Split(line, " ")[0]
			continue
		}

		// If we find data, fill current map
		if currentMapName != "" && line != "" {
			dataMaps[currentMapName] = MapStruct{
				name:  currentMapName,
				from:  strings.Split(currentMapName, "-")[0],
				to:    strings.Split(currentMapName, "-")[2],
				lines: append(dataMaps[currentMapName].lines, line),
			}
		}

		// If empty line, we say that the map finished
		if line == "" {
			currentMapName = ""
			continue
		}
	}

	return seeds, dataMaps
}

type MapStruct struct {
	name  string
	lines []string
	from  string
	to    string
}
