package main

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
}

func main() {
	nodeD := &Node{val: 10, next: nil}
	nodeC := &Node{val: 8, next: nodeD}
	nodeB := &Node{val: 7, next: nodeC}
	nodeF := &Node{val: 1, next: nodeC}
	rootB := &Node{val: 99, next: nodeF}
	rootA := &Node{val: 3, next: nodeB}
	interNode := getIntersectNode(rootA, rootB)
	fmt.Println(interNode.val)
}

// Traverses rootA until 'fast' reaches nil then
// Continues along rootB until slow == fast
func getIntersectNode(r1 *Node, r2 *Node) *Node {
	var ls *Node
	slow := r1
	fast := r1.next
	for ; slow != fast; ls, slow = slow, slow.next {
		if fast.next != nil {
			fast = fast.next.next
		} else {
			fast = r2.next
		}
	}
	return ls
}
