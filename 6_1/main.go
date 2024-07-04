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
	var times, distances []int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if fields[0] == "Time:" {
			for _, s := range fields[1:] {
				time, err := strconv.Atoi(s)
				if err != nil {
					log.Fatal(err)
				}
				times = append(times, time)
			}
		} else {
			for _, s := range fields[1:] {
				distance, err := strconv.Atoi(s)
				if err != nil {
					log.Fatal(err)
				}
				distances = append(distances, distance)
			}
		}
	}

	races := []Race{}
	for i, _ := range times {
		races = append(races,
			Race{
				time: times[i], distance: distances[i],
			})
	}

	return races
}

type Race struct {
	time     int
	distance int
}
