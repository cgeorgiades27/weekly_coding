package main

import (
	"testing"
)

var nodeF = &Node{val: 7, left: nil, right: nil}
var nodeE = &Node{val: 5, left: nil, right: nil}
var nodeD = &Node{val: 3, left: nil, right: nil}
var nodeC = &Node{val: 1, left: nil, right: nil}
var nodeB = &Node{val: 6, left: nodeE, right: nodeF}
var nodeA = &Node{val: 2, left: nodeC, right: nodeD}
var goodRoot = &Node{val: 4, left: nodeA, right: nodeB}
var badRoot = &Node{val: 4, left: nodeA, right: nodeC}
var tests = []struct {
	root Node
	exp  bool
}{
	{*goodRoot, true},
	{*badRoot, false},
}

func TestCheckBST(t *testing.T) {
	for _, test := range tests {
		result := checkBST(&test.root)
		if result != test.exp {
			t.Errorf("Test Failed: returned %t but expected %t", result, test.exp)
		}
	}
}

func BenchmarkCheckBST(b *testing.B) {
	test := tests[0]
	for i := 0; i < b.N; i++ {
		checkBST(&test.root)
	}
}
