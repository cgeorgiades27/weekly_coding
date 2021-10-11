// Given a complete binary tree, count the number of nodes in faster than **O(n)** time.
// Recall that a complete binary tree has every level filled except the last,
// and the nodes in the last level are filled starting from the left.
package main

import "fmt"

func main() {
	a := &Node{val: 1, left: nil, right: nil}

	a.left = &Node{val: 2, left: nil, right: nil}
	a.right = &Node{val: 3, left: nil, right: nil}

	a.left.left = &Node{val: 4, left: nil, right: nil}
	a.left.right = &Node{val: 5, left: nil, right: nil}
	a.right.left = &Node{val: 6, left: nil, right: nil}
	a.right.right = &Node{val: 7, left: nil, right: nil}

	a.left.left.left = &Node{val: 8, left: nil, right: nil}
	a.left.left.right = &Node{val: 9, left: nil, right: nil}
	a.left.right.left = &Node{val: 10, left: nil, right: nil}
	a.left.right.right = &Node{val: 11, left: nil, right: nil}
	a.right.left.left = &Node{val: 12, left: nil, right: nil}
	a.right.left.right = &Node{val: 13, left: nil, right: nil}
	a.right.right.left = &Node{val: 14, left: nil, right: nil}
	// a.right.right.right = &Node{val: 15, left: nil, right: nil}
	ClimbHelper(a)
}

type Node struct {
	val   int
	left  *Node
	right *Node
}

func ClimbTree(root *Node, nodes *[]*Node) {
	if root == nil {
		return
	}
	*nodes = append(*nodes, root)
	ClimbTree(root.left, nodes)
	ClimbTree(root.right, nodes)
}

func ClimbHelper(root *Node) {
	var nodes []*Node
	ClimbTree(root, &nodes)
	fmt.Printf("%-15s%5d\n%-15s%5d\n", "Total nodes:", len(nodes), "Last added:", nodes[len(nodes)-1].val)
}
