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
		if strings.HasSuffix(node.Label, "A") {
			fmt.Println("Found starting node", node)
			nodePaths = append(nodePaths, node)
		}
	}

	nodeStepsToZ := make([]int, len(nodePaths))
	fmt.Println("Initial nodestepstoz", nodeStepsToZ)

	counter := 0
	for counter < 10000000 {
		for _, m := range moves {
			for i, node := range nodePaths {
				if string(m) == "R" {
					nodePaths[i] = nodeMap[node.RightNode]
				} else if string(m) == "L" {
					nodePaths[i] = nodeMap[node.LeftNode]
				}

				counter++

				label := nodePaths[i].Label
				if string(label[2]) == "Z" {
					if nodeStepsToZ[i] == 0 {
						fmt.Println("Updating", i, " ", nodeStepsToZ[i], " to", counter)
						nodeStepsToZ[i] = counter
					}

					allPathsPresent := true
					for _, step := range nodeStepsToZ {
						if step == 0 {
							allPathsPresent = false
							break
						}
					}

					if !allPathsPresent {
						continue
					}

					result := nodeStepsToZ[0]
					for step := range nodeStepsToZ[1:] {
						result = lcm(result, step)
					}
					fmt.Println(result)
					return
				}
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

// Partially copied lcm from: https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/

// greatest common divisor (GCD) via Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (lcm) via GCD
func lcm(a int, b int) int {
	return a * b / gcd(a, b)
}
