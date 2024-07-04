package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	// "runtime"
	"runtime/pprof"
	"strconv"
	"strings"
	"sync"
)

func main() {
	f, err := os.Create("cpu.pprof")
	if err != nil {
		panic(err)
	}
	err = pprof.StartCPUProfile(f)
	if err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()

	seeds, dataMaps := getAllMapsAndSeeds()

	// cores := runtime.NumCPU()
	// semaphore := make(chan struct{}, cores)
	var wg sync.WaitGroup
	var mu sync.Mutex
	semaphore := make(chan struct{}, 4)

	lowestLocation := math.MaxInt
	for i, s := range seeds {
		if i%2 == 0 {
			continue
		}
		counter := 0
		for j := 0; j < seeds[i+1]; j++ {
			seed := s + j
			wg.Add(1)
			semaphore <- struct{}{}

			go func(seed int) {
				defer wg.Done()
				defer func() { <-semaphore }()

				seedLocation := getSeedLocation(seed, dataMaps)
				mu.Lock()
				if seedLocation < lowestLocation {
					lowestLocation = seedLocation
				}
				mu.Unlock()
			}(seed)
			counter++
			if counter > 1000000 {
				// break
			}
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

func getLineInfo(line string) LineInfo {
	splitted := strings.Split(line, " ")
	destinationStart, err := strconv.Atoi(splitted[0])
	if err != nil {
		log.Fatal(err)
	}
	lineSourceStart, err := strconv.Atoi(splitted[1])
	if err != nil {
		log.Fatal(err)
	}
	lineSourceOffset, err := strconv.Atoi(splitted[2])
	if err != nil {
		log.Fatal(err)
	}
	return LineInfo{
		DestinationStart: destinationStart,
		SourceStart:      lineSourceStart,
		SourceOffset:     lineSourceOffset,
	}
}

func getDestination(sourceVal int, currentDataMap MapStruct) int {
	result := sourceVal
	for _, line := range currentDataMap.lineInfos {
		lineSourceEnd := line.SourceStart + line.SourceOffset - 1
		if sourceVal >= line.SourceStart && sourceVal <= lineSourceEnd {
			result = line.DestinationStart + (sourceVal - line.DestinationStart)
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
				name:      currentMapName,
				from:      strings.Split(currentMapName, "-")[0],
				to:        strings.Split(currentMapName, "-")[2],
				lines:     append(dataMaps[currentMapName].lines, line),
				lineInfos: append(dataMaps[currentMapName].lineInfos, getLineInfo(line)),
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
	name      string
	lines     []string
	lineInfos []LineInfo
	from      string
	to        string
}

type LineInfo struct {
	DestinationStart int
	SourceStart      int
	SourceOffset     int
}
