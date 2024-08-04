package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strings"
)

func main() {
	filePath := "input.txt"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	f1, err := os.Create("cpu.pprof")
	if err != nil {
		panic(err)
	}
	pprof.StartCPUProfile(f1)
	defer pprof.StopCPUProfile()

	scanner := bufio.NewScanner(f)

	// Get nodes and moves
	moves := ""
	nodePaths := []Node{}
	nodeMap := make(map[int]Node)
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
		nodeMap[node.NodeKey] = node
		if strings.HasSuffix(node.Label, "A") {
			fmt.Println("Found starting node", node)
			nodePaths = append(nodePaths, node)
		}
	}

	counter := 0
	for counter < 100000000 {
		for _, m := range moves {
			for i, node := range nodePaths {
				if string(m) == "R" {
					nodePaths[i] = nodeMap[node.RightNodeKey]
				} else if string(m) == "L" {
					nodePaths[i] = nodeMap[node.LeftNodeKey]
				}
			}

			counter++

			// Check if we are at Z
			allNodesAreZ := true
			for _, node := range nodePaths {
				if !strings.HasSuffix(node.Label, "Z") {
					allNodesAreZ = false
					break
				}
			}

			if allNodesAreZ {
				fmt.Println(counter)
				fmt.Println(nodePaths)
				return
			}
		}
	}
}

func getNodeKey(label string) int {
	return (int(label[0]-'A')*26*26 +
		int(label[1]-'A')*26 +
		int(label[2]-'A'))
}

func getNode(line string) Node {
	label := line[:3]
	leftNode := line[7:10]
	rightNode := line[12:15]
	return Node{label, getNodeKey(label), getNodeKey(leftNode), getNodeKey(rightNode)}
}

type Node struct {
	Label        string
	NodeKey      int
	LeftNodeKey  int
	RightNodeKey int
}
