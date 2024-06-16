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
	filePath := "input.txt"

	// Scan file line by line
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		// Extract number and add to sum
		gameLine := scanner.Text()
		ballPerGame, err := parseGame(gameLine)
		if err != nil {
			log.Fatal(err)
		}
		if !isGameValid(ballPerGame) {
			continue
		}
		gameValue, err := extractGameValue(gameLine)
		if err != nil {
			log.Fatal(err)
		}
		sum += gameValue
	}

	// Sum
	fmt.Println(sum)
}

func extractGameValue(game string) (int, error) {
	//  "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	gamePart := strings.Split(game, ":")[0]
	//  "Game 1"
	gameValue := strings.Split(gamePart, " ")[1]
	val, err := strconv.Atoi(gameValue)
	if err != nil {
		return -1, err
	}
	return val, nil

}

func isGameValid(game BallsPerGame) bool {
	// only 12 red cubes, 13 green cubes, and 14 blue cubes
	if game.MaxRed > 12 {
		return false
	}
	if game.MaxGreen > 13 {
		return false
	}
	if game.MaxBlue > 14 {
		return false
	}
	return true
}

func parseGame(game string) (BallsPerGame, error) {
	result := BallsPerGame{}
	//  "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	stringWithoutGame := strings.Split(game, ":")[1]
	//  " 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	allRounds := strings.Split(stringWithoutGame, ";")

	for _, round := range allRounds {
		//  " 3 blue, 4 red"
		allBallsPerRound := strings.Split(round, ",")
		for _, ballColor := range allBallsPerRound {
			// " 3 blue"
			ball := strings.Split(ballColor, " ")
			ballNumber, err := strconv.Atoi(ball[1])
			if err != nil {
				return BallsPerGame{}, err
			}
			ballColor := ball[2]

			switch ballColor {
			case "red":
				if result.MaxRed < ballNumber {
					result.MaxRed = ballNumber
				}
			case "green":
				if result.MaxGreen < ballNumber {
					result.MaxGreen = ballNumber
				}
			case "blue":
				if result.MaxBlue < ballNumber {
					result.MaxBlue = ballNumber
				}
			}
		}
	}

	return result, nil
}

type BallsPerGame struct {
	MaxBlue  int
	MaxRed   int
	MaxGreen int
}
