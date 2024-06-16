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
		gamePower := getGamePower(ballPerGame)
		sum += gamePower
	}

	// Sum
	fmt.Println(sum)
}

func getGamePower(game BallsPerGame) int {
	return game.MaxRed * game.MaxGreen * game.MaxBlue
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
