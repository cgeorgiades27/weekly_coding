package main

import (
	"fmt"
)

type Node struct {
	val     int
	left    *Node
	right   *Node
	visited bool
}

func getPaths(root *Node) [][]int {
	qu := []*Node{root}
	path := make([]int, 0)
	paths := make([][]int, 0)
	for ; len(qu) > 0; qu = qu[1:] {
		currNode := qu[0]

		if currNode.left != nil && currNode.right != nil {
			path = append(path, currNode.val)

			for currNode.left != nil && !currNode.left.visited {
				currNode = currNode.left
				currNode.visited = true
				qu = append(qu, currNode)
				path = append(path, currNode.val)
			}

			if len(path) > 1 {
				paths = append(paths, path)
				path = []int{1}
			}
			currNode = qu[0]

			for currNode.right != nil && !currNode.right.visited {
				currNode = currNode.right
				currNode.visited = true
				qu = append(qu, currNode)
				path = append(path, currNode.val)
			}

			if len(path) > 1 {
				paths = append(paths, path)
				path = []int{1}
			}
		}
	}
	return paths
}

func main() {
	leafD := &Node{val: 5, left: nil, right: nil, visited: false}
	leafC := &Node{val: 4, left: nil, right: nil, visited: false}
	leafB := &Node{val: 3, left: leafC, right: leafD, visited: false}
	leafA := &Node{val: 2, left: nil, right: nil, visited: false}
	root := &Node{val: 1, left: leafA, right: leafB, visited: false}

	fmt.Println(getPaths(root))
}
