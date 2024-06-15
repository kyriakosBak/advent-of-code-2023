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
