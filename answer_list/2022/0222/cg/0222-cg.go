// ## Weekly Coding Problem
// ### Week 02, 2022 - (Medium)
// This problem was asked by **LinkedIn**.
// Determine whether a tree is a valid binary search tree.
// A binary search tree is a tree with two children, `left` and `right`, and satisfies
// the constraint that the key in the `left` child
// must be less than or equal to the root and the key in the `right` child
// must be greater than or equal to the root.
//        4
//       / \
// 	   2     6
// 	  / \   / \
// 	 1   3 5   7
package main

import (
	"fmt"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

func main() {
	bst := &Node{
		val: 4,
		left: &Node{
			val: 2,
			left: &Node{
				val: 1,
			},
			right: &Node{
				val: 3,
			},
		},
		right: &Node{
			val: 6,
			left: &Node{
				val: 5,
			},
			right: &Node{
				val: 7,
			},
		},
	}
	fmt.Println(checkHelper(bst))
}

func checkHelper(root *Node) bool {
	if root == nil {
		return false
	}
	return checkTree(root.left) && checkTree(root.right)
}

func checkTree(root *Node) bool {
	if root.left == nil || root.right == nil {
		return true
	}
	if root.left.val > root.val {
		return false
	}
	if root.right.val < root.val {
		return false
	}
	checkTree(root.right)
	checkTree(root.left)
	return true
}
