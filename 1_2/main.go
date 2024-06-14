package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	filePath := "data.txt"

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
		sum += stringToNumber(scanner.Text())
	}

	// Sum
	fmt.Println(sum)
}

func stringToNumber(line string) int {
	var runeArray [2]string
	for _, rune := range line {
		if unicode.IsDigit(rune) {
			if runeArray[0] == "" {
				runeArray[0] = string(rune)
				runeArray[1] = string(rune)
				continue
			}
			runeArray[1] = string(rune)
		}
	}
	intResult, err := strconv.Atoi(runeArray[0] + runeArray[1])
	if err != nil {
		log.Fatal(err)
	}
	return intResult
}

func checkTextNumberExists(test string) (int, int) {

}

func textToNumberMap(text string) int {
	switch text {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	}
	return -1
}
