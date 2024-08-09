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

	nodeStepsToZ := make(map[string]int)
	for _, node := range nodePaths {
		nodeStepsToZ[node.Label] = -1
	}

	counter := 0
	for counter < 100000 {
		for _, m := range moves {
			for i, node := range nodePaths {
				if string(m) == "R" {
					nodePaths[i] = nodeMap[node.RightNode]
				} else if string(m) == "L" {
					nodePaths[i] = nodeMap[node.LeftNode]
				}

				counter++

				if strings.HasSuffix(nodePaths[i].Label, "Z") {
					if nodeStepsToZ[nodePaths[i].Label] == -1 {
						nodeStepsToZ[nodePaths[i].Label] = counter
					}

					for _, steps := range nodeStepsToZ {
						if steps == -1 {
							continue
						}
					}

					steps := []int{}
					for _, step := range nodeStepsToZ {
						steps = append(steps, step)
					}
					lcm := LCM(steps[0], steps[1], steps[2:]...)
					fmt.Println(lcm)
					fmt.Println(nodeStepsToZ)
					fmt.Println(counter)
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

// Copied lcm from: https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
