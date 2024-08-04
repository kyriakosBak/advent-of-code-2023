package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	hands := getHands()
	sorted := sortHands(hands)

	fmt.Println(sorted)
	sum := 0
	for i, hand := range sorted {
		sum += (i + 1) * hand.bid
	}
	fmt.Println(sum)
}

func sortHands(hands []Hand) []Hand {
	sort.Slice(hands, func(i, j int) bool {
		return isHandStronger(hands[j], hands[i])
	})
	return hands
}

func getHands() []Hand {
	filePath := "input.txt"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	result := []Hand{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		bid, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, Hand{fields[0], bid})

	}
	return result
}

// Returns true if handA is stronger
func isHandStronger(handA Hand, handB Hand) bool {
	handAStrength := handA.getHandStength()
	handBStrength := handB.getHandStength()

	if handAStrength > handBStrength {
		return true
	} else if handAStrength < handBStrength {
		return false
	} else {
		for i := range handA.cards {
			if getCardValue(handA.cards[i]) > getCardValue(handB.cards[i]) {
				return true
			}
			if getCardValue(handA.cards[i]) < getCardValue(handB.cards[i]) {
				return false
			}
		}
	}
	return false
}

func getCardValue(card byte) int {
	// Numbers return their value
	if card >= 48 && card <= 57 {
		return int(card)
	}
	// For the charactes return bigger values

	// T
	if card == 84 {
		return 58
	}
	// J
	if card == 74 {
		return 46
	}
	// Q
	if card == 81 {
		return 60
	}
	// K
	if card == 75 {
		return 61
	}
	// A
	if card == 65 {
		return 62
	}
	panic("could not get card value")
}

type Hand struct {
	cards string
	bid   int
}

func (h *Hand) getHandStength() int {
	if h.isFiveOfAKind() {
		return 7
	} else if h.isFourOfAKind() {
		return 6
	} else if h.isFullHouse() {
		return 5
	} else if h.isThreeOfAKind() {
		return 4
	} else if h.isTwoPair() {
		return 3
	} else if h.isOnePair() {
		return 2
	} else {
		return 1
	}
}

func (h *Hand) isHighCard() bool {
	cardsMap := make(map[rune]int)
	for _, c := range h.cards {
		cardsMap[c]++
	}
	if len(cardsMap) == 5 {
		return true
	}
	return false
}

func (h *Hand) isOnePair() bool {
	cardsMap := make(map[rune]int)
	for _, c := range h.cards {
		cardsMap[c]++
	}
	JOccurences, exists := cardsMap[74]
	for _, c := range cardsMap {
		if c == 2 && len(cardsMap) == 4 {
			return true
		} else if c == 1 && len(cardsMap) == 5 && exists && JOccurences == 1 {
			return true
		}
	}
	return false
}

func (h *Hand) isTwoPair() bool {
	cardsMap := make(map[rune]int)
	for _, c := range h.cards {
		cardsMap[c]++
	}
	if len(cardsMap) == 3 {
		for _, c := range cardsMap {
			if c == 2 {
				return true
			}
		}
	}
	return false
}

func (h *Hand) isThreeOfAKind() bool {
	cardsMap := make(map[rune]int)
	for _, c := range h.cards {
		cardsMap[c]++
	}

	JOccurences, exists := cardsMap[74]
	for _, c := range cardsMap {
		if c == 3 && len(cardsMap) == 3 {
			return true
		} else if c == 2 && len(cardsMap) == 4 && exists && JOccurences == 1 {
			return true
		} else if c == 1 && len(cardsMap) == 4 && exists && JOccurences == 2 {
			return true
		}
	}
	return false
}

func (h *Hand) isFullHouse() bool {
	cardsMap := make(map[rune]int)
	for _, c := range h.cards {
		cardsMap[c]++
	}
	JOccurences, exists := cardsMap[74]
	for _, c := range cardsMap {
		if c == 3 && len(cardsMap) == 2 {
			return true
		} else if c == 2 && len(cardsMap) == 3 && exists && JOccurences == 1 {
			return true
		}
	}
	return false
}

func (h *Hand) isFourOfAKind() bool {
	cardsMap := make(map[rune]int)
	for _, c := range h.cards {
		cardsMap[c]++
	}
	JOccurences, exists := cardsMap[74]
	for _, c := range cardsMap {
		if c == 4 && !exists {
			return true
		} else if c == 3 && exists && JOccurences == 1 {
			return true
		} else if c == 2 && exists && JOccurences == 2 {
			return true
		} else if c == 1 && exists && JOccurences == 3 {
			return true
		}
	}
	return false
}

func (h *Hand) isFiveOfAKind() bool {
	cardsMap := make(map[rune]int)
	for _, c := range h.cards {
		cardsMap[c]++
	}
	if len(cardsMap) == 1 {
		return true
	} else if len(cardsMap) == 2 {
		_, exists := cardsMap[74]
		if exists {
			return true
		}
	}
	return false
}
