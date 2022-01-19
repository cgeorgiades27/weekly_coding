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
	nodeF := &Node{val: 7, left: nil, right: nil}
	nodeE := &Node{val: 5, left: nil, right: nil}
	nodeD := &Node{val: 3, left: nil, right: nil}
	nodeC := &Node{val: 1, left: nil, right: nil}
	nodeB := &Node{val: 6, left: nodeE, right: nodeF}
	nodeA := &Node{val: 2, left: nodeC, right: nodeD}
	root := &Node{val: 4, left: nodeA, right: nodeB}
	isBST := checkBST(root)

	fmt.Println(isBST)
}

func checkBST(top *Node) bool {
	qu := []*Node{top}
	node := qu[0]
	bst := true
	for i := 0; node != nil; i, node = i+1, qu[i] {

		if node.left == nil && node.right == nil {
			break
		}

		if node.left.val > node.val {
			bst = false
			break
		}
		qu = append(qu, node.left)

		if node.right.val < node.val {
			bst = false
			break
		}
		qu = append(qu, node.right)
	}
	return bst
}
