// ### Week 01, 2022 - (Easy)

// This problem was asked by **Google**.

// Given two singly linked lists that intersect at some point, find the intersecting node. The lists are non-cyclical.

// For example, given A = 3 -> 7 -> 8 -> 10 and B = 99 -> 1 -> 8 -> 10, return the node with value 8.

// In this example, assume nodes with the same value are the exact same node objects.

// Do this in O(M + N) time (where M and N are the lengths of the lists) and constant space.
package main

import "fmt"

type Node struct {
	val_  int
	next_ *Node
}

func main() {
	intNode := &Node{val_: 8, next_: &Node{val_: 10, next_: nil}}
	a := &Node{val_: 3, next_: &Node{val_: 7, next_: intNode}}
	b := &Node{val_: 99, next_: &Node{val_: 1, next_: intNode}}
	fmt.Printf("Intersecting Node val: %d\n", FindIntersection(a, b))

}

func FindIntersection(a, b *Node) int {
	myPtr := &Node{}
	cura := a
	curb := b
	for {
		if cura.next_ == nil {
			break
		}
		temp := cura
		cura = cura.next_
		temp.next_ = myPtr
	}
	for {
		if curb.next_ == nil {
			break
		}
		if curb.next_ == myPtr {
			return curb.val_
		}
		temp := curb
		curb = curb.next_
		temp.next_ = myPtr
	}
	return -1
}
