package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strings"
)

func main() {
	filePath := "input.txt"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		winningNumbers := getWinningNumbers(line)
		myNumbers := getMyNumbers(line)
		sum += getPoints(winningNumbers, myNumbers)
	}
	fmt.Println(sum)
}

func getPoints(winningNumbers []string, myNumbers []string) int {
	// Slow iteration. Could be faster with a map
	numbersInWinning := 0
	for _, number := range myNumbers {
		if slices.Contains(winningNumbers, number) {
			numbersInWinning++
		}
	}
	if numbersInWinning == 1 {
		return 1
	} else if numbersInWinning > 1 {
		return int(math.Pow(2, float64(numbersInWinning-1)))
	}
	return 0
}

func getMyNumbers(line string) []string {
	result := []string{}

	myNumbersString := strings.Split(strings.Split(line, ":")[1], "|")[1]
	for _, field := range strings.Split(myNumbersString, " ") {
		if field == "" || field == " " || field == "  " {
			continue
		}
		result = append(result, field)
	}

	return result
}

func getWinningNumbers(line string) []string {
	result := []string{}

	winningNumbersString := strings.Split(strings.Split(line, ":")[1], "|")[0]
	for _, field := range strings.Split(winningNumbersString, " ") {
		if field == "" || field == " " || field == "  " {
			continue
		}
		result = append(result, field)
	}

	return result
}
