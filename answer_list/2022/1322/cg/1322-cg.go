// Weekly Coding Problem

// Week 13, 2022 - (Medium)

// This problem was asked by **Google**.

// Given the sequence of keys visited by a postorder traversal of a binary search tree, reconstruct the tree.

// For example, given the sequence 2, 4, 3, 8, 7, 5, you should construct the following tree:

//     5
//    / \
//   3   7
//  / \   \
// 2   4   8

package main

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

func main() {

	postorder := []int{2, 4, 3, 8, 7, 5}
	var root *Node

	for i := len(postorder) - 1; i >= 0; i-- {
		node := Node{val: postorder[i]}
		if i == len(postorder)-1 {
			root = &node
		} else {
			AddNode(root, &node)
		}
	}
	depth := 0
	PrintPreorder(root, depth)
}

func AddNode(root, node *Node) {
	// base case
	if root == nil {
		return
	}

	if node.val > root.val {
		if root.right == nil {
			root.right = node
		} else {
			AddNode(root.right, node)
		}
	}

	if node.val < root.val {
		if root.left == nil {
			root.left = node
		} else {
			AddNode(root.left, node)
		}
	}
}

func PrintPreorder(root *Node, depth int) {
	// base case
	if root == nil {
		return
	}

	var str string
	for i := 0; i < depth; i++ {
		str += "| "
	}
	fmt.Printf("%s| -  %d\n", str, root.val)
	depth++

	PrintPreorder(root.left, depth)
	PrintPreorder(root.right, depth)
}
