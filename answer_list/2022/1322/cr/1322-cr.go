package main

import (
	"fmt"
)

func main() {
	nodes := []int{2, 4, 3, 8, 7, 5}
	root := buildBST(nodes)
	fmt.Printf("Here ya go, but last one...these BSTs aint good for ya kid!\nInventory (+1): Root Node %v", root)
}

type node struct {
	val   int
	left  *node
	right *node
}

func buildBST(input []int) *node {
	length := len(input) - 1
	root := node{val: input[length], left: nil, right: nil}
	currNode := &root
	last := &root
	for i := length - 1; i >= 0; i-- {
		v := input[i]
		if v > currNode.val {
			tempNode := &node{val: v, left: nil, right: nil}
			currNode.right = tempNode
			currNode = tempNode
		} else {
			tempNode := &node{val: v, left: nil, right: nil}
			currNode = last
			currNode.left = tempNode
			last = tempNode
			currNode = tempNode
		}
	}
	return &root
}
