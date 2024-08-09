package main

import (
	"testing"
)

func TestGetNodeReturnsCorrectNode(t *testing.T) {
	line := "TMB = (KFQ, DMD) "

	node := getNode(line)

	if node.Label != "TMB" {
		t.Error("Node label incorrect. Expexted: TMB but got", node.Label)
	}
	if node.LeftNode != "KFQ" {
		t.Error("Node left node incorrect. Expexted: KFQ but got", node.LeftNode)
	}
	if node.RightNode != "DMD" {
		t.Error("Node right node incorrect. Expexted: DMD but got", node.RightNode)
	}
}

func TestGcd(t *testing.T) {
	numA := 15
	numB := 90
	expected := 15

	result := gcd(numA, numB)

	if result != expected {
		t.Error("Expected: ", expected, " but got: ", result)
	}
}

func TestLcm(t *testing.T) {
	numA := 12
	numB := 8
	expected := 24

	result := lcm(numA, numB)

	if result != expected {
		t.Error("Expected: ", expected, " but got: ", result)
	}
}

func TestLcmArray(t *testing.T) {
	nums := []int{12, 15, 75}
	expected := 300

	result := nums[0]
	for _, a := range nums[1:] {
		result = lcm(result, a)
	}

	if result != expected {
		t.Error("Expected: ", expected, " but got: ", result)
	}
}
