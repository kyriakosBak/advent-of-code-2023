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
	var currentNumber []rune
	for i, line := range matrix {
		hasSymbol := false
		for j, char := range line {
			if unicode.IsDigit(char) {
				currentNumber = append(currentNumber, char)
				if hasAdjacentSymbol(matrix, i, j) {
					hasSymbol = true
				}
				continue
			}
			// If not digit, check if it has symbol and add it to sum
			if hasSymbol {
				fmt.Println("adding number %d", getNumber(currentNumber))
				sum += getNumber(currentNumber)
			}
			currentNumber = nil
			hasSymbol = false
		}
		// Add number at end of line if it exists
		if hasSymbol {
			fmt.Println("adding number %d", getNumber(currentNumber))
			sum += getNumber(currentNumber)
		}
		currentNumber = nil
	}
	return sum
}

func getNumber(number []rune) int {
	if len(number) <= 0 {
		return 0
	}
	convertedNumber, err := strconv.Atoi(string(number))
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return convertedNumber
}

func hasAdjacentSymbol(matrix [][]rune, currentRow int, currentCol int) bool {
	directions := [][]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1}, // Up, Down, Left, Right
		{-1, -1}, {-1, 1}, {1, -1}, {1, 1}, // Diagonals: Top-Left, Top-Right, Bottom-Left, Bottom-Right
	}

	for _, direction := range directions {
		newRow := currentRow + direction[0]
		newCol := currentCol + direction[1]

		if newRow < 0 || newRow >= len(matrix) || newCol < 0 || newCol >= len(matrix[0]) {
			continue
		}

		if !unicode.IsDigit(matrix[newRow][newCol]) && matrix[newRow][newCol] != 46 {
			return true
		}
	}
	return false
}
