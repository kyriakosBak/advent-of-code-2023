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
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var fileMatrix [][]rune
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		fileMatrix = append(fileMatrix, row)
	}

	sum := getSum(fileMatrix)
	fmt.Println(sum)
}

func getSum(matrix [][]rune) int {
	// Go through each row and character
	sum := 0
	hasSymbol := false
	var currentNumber []rune
	for i, line := range matrix {
		for j, char := range line {
			if !unicode.IsDigit(char) {
				// if there is an adjacent symbol and we have recorded a number
				// add it to our sum
				if hasSymbol && len(currentNumber) > 0 {
					convertedNumber, err := strconv.Atoi(string(currentNumber))
					if err != nil {
						log.Fatal(err)
					}
					sum += convertedNumber
				}
				currentNumber = nil
				continue
			}
			// If it's a digit, check all adjacents matrix nodes for a symbol
			currentNumber = append(currentNumber, char)
			if hasAdjacentSymbol(matrix, i, j) {
				hasSymbol = true
			}
		}
	}
	return sum
}

func hasAdjacentSymbol(matrix [][]rune, currentRow int, currentCol int) bool {
	directions := [][]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1}, // Up, Down, Left, Right
		{-1, -1}, {-1, 1}, {1, -1}, {1, 1}, // Diagonals: Top-Left, Top-Right, Bottom-Left, Bottom-Right
	}

	for _, direction := range directions {
		newRow := currentRow + direction[0]
		newCol := currentCol + direction[1]

		if newRow < 0 || newRow >= len(matrix) || newCol < 0 || newCol > len(matrix[0]) {
			continue
		}

		if !unicode.IsDigit(matrix[newRow][newCol]) && matrix[newRow][newCol] != 46 {
			return true
		}
	}
	return false
}
