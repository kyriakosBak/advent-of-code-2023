package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	filePath := "input.txt"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	// Get nodes and moves
	moves := ""
	var nextNode Node
	nodeMap := make(map[string]Node)
	for scanner.Scan() {
		line := scanner.Text()

		// if line empty skip
		if line == "" {
			continue
		}

		// if first line save the moves
		if !strings.Contains(line, " = ") {
			moves = line
			continue
		}

		node := getNode(line)
		nodeMap[node.Label] = node
		if node.Label == "AAA" {
			fmt.Println("First node", node)
			nextNode = node
		}
	}

	fmt.Println(moves)

	counter := 0
	for i := 0; i < 1000; i++ {
		for _, m := range moves {
			if string(m) == "R" {
				nextNode = nodeMap[nextNode.RightNode]
			} else if string(m) == "L" {
				nextNode = nodeMap[nextNode.LeftNode]
			}

			counter++
			fmt.Println(string(m), nextNode, counter)
			if nextNode.Label == "ZZZ" {
				fmt.Println(counter)
				return
			}
		}
	}
}

func getNode(line string) Node {
	label := line[:3]
	leftNode := line[7:10]
	rightNode := line[12:15]
	return Node{label, leftNode, rightNode}
}

type Node struct {
	Label     string
	LeftNode  string
	RightNode string
}
