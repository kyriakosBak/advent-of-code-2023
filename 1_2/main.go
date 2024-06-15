package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
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
		res, err := stringToNumber(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		sum += res
	}

	// Sum
	fmt.Println(sum)
}

func stringToNumber(line string) (int, error) {
	var runeArray [2]string
	for i, char := range line {
		if unicode.IsDigit(char) {
			if runeArray[0] == "" {
				runeArray[0] = string(char)
			}
			runeArray[1] = string(char)
		} else {
			number := checkTextNumberExistsAtTheStartOfText(line[i:])
			if number == -1 {
				continue
			}
			if runeArray[0] == "" {
				runeArray[0] = strconv.Itoa(number)
			}
			runeArray[1] = strconv.Itoa(number)
		}
	}
	intResult, err := strconv.Atoi(runeArray[0] + runeArray[1])
	if err != nil {
		return -1, err
	}
	return intResult, nil
}

// Returns -1 if it cannot find a number at index 0.
func checkTextNumberExistsAtTheStartOfText(text string) int {
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for _, num := range numbers {
		index := strings.Index(text, num)
		if index == 0 {
			return textToNumberMap(num)
		}
	}
	return -1
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
