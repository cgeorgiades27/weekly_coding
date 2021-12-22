// Weekly Coding Problem
// Week 49, 2021 - (Medium)
// This problem was asked by Facebook.
// Given a binary tree, return all paths from the root to leaves.
// For example, given the tree:
//    1
//   / \
//  2   3
//     / \
//    4   5
// Return [[1, 2], [1, 3, 4], [1, 3, 5]].
package main

import (
	"fmt"
)

func main() {
	root := &Node{val: 1,
		left:  &Node{val: 2},
		right: &Node{val: 3, left: &Node{val: 4}, right: &Node{val: 5}}}
	var container [][]int
	GetPaths(root, []int{}, &container)
	fmt.Println(container)
}

// Node: tree elements
type Node struct {
	left  *Node
	right *Node
	val   int
}

// GetPaths records unique paths to reach binary tree leaves
func GetPaths(root *Node, slc []int, pathSlc *[][]int) {
	// base case
	if root == nil {
		return
	}
	slc = append(slc, root.val)
	GetPaths(root.left, slc, pathSlc)
	GetPaths(root.right, slc, pathSlc)

	// if a leaf is encountered store current route and pop
	if root.left == nil && root.right == nil {
		*pathSlc = append(*pathSlc, slc)
		slc = (slc)[:len(slc)-1]
	}
}
