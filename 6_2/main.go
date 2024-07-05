package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	races := readData()
	prod := 1
	for _, r := range races {
		sum := 0
		for timePressed := 1; timePressed < r.time; timePressed++ {
			dist := getDistancedTraveled(timePressed, r.time)
			if dist > r.distance {
				sum++
			}
		}
		if sum > 0 {
			prod *= sum
		}
	}

	fmt.Println(prod)
}

func getDistancedTraveled(milisPressed int, raceTime int) int {
	if milisPressed <= 0 || milisPressed >= raceTime {
		return 0
	}

	return milisPressed * (raceTime - milisPressed)
}

func readData() []Race {
	filePath := "input.txt"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	var times, distances []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if fields[0] == "Time:" {
			for _, f := range fields[1:] {
				times = append(times, f)
			}
		} else {
			for _, s := range fields[1:] {
				distances = append(distances, s)
			}
		}
	}

	totalTimeString := strings.Join(times, "")
	totalTimeInt, err := strconv.Atoi(totalTimeString)
	if err != nil {
		log.Fatal(err)
	}

	totalDistanceString := strings.Join(distances, "")
	totalDistanceInt, err := strconv.Atoi(totalDistanceString)
	if err != nil {
		log.Fatal(err)
	}

	return []Race{{time: totalTimeInt, distance: totalDistanceInt}}
}

type Race struct {
	time     int
	distance int
}
