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
