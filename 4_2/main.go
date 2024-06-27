package main

import (
	"bufio"
	"fmt"
	"log"
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
	copiesPerCard := make([]int, 1000)
	currentCardCounter := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		winningNumbers := getWinningNumbers(line)
		myNumbers := getMyNumbers(line)
		matchingNumbers := getNumberOfMatchingNumbers(winningNumbers, myNumbers)
		copiesPerCard[currentCardCounter]++
		for i := 1; i < matchingNumbers+1; i++ {
			copiesPerCard[currentCardCounter+i] += copiesPerCard[currentCardCounter]
		}
		fmt.Println("adding", copiesPerCard[currentCardCounter], "for card", currentCardCounter)
		sum += copiesPerCard[currentCardCounter]
		currentCardCounter++
	}
	fmt.Println(sum)
}

func getNumberOfMatchingNumbers(winningNumbers []string, myNumbers []string) int {
	// Slow iteration. Could be faster with a map
	result := 0
	for _, number := range myNumbers {
		if slices.Contains(winningNumbers, number) {
			result++
		}
	}
	return result
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
