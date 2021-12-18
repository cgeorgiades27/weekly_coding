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

func get_paths(root *Node) [][]int {
	stack := []*Node{root}
	path := make([]int, 0)
	paths := make([][]int, 0)
	for len(stack) > 0 {
		cur_node := stack[0]

		if cur_node.left != nil && cur_node.right != nil {
			path = append(path, cur_node.val)

			for cur_node.left != nil && !cur_node.left.visited {
				cur_node = cur_node.left
				cur_node.visited = true
				stack = append(stack, cur_node)
				path = append(path, cur_node.val)
			}

			if len(path) > 1 {
				paths = append(paths, path)
				path = []int{1}
			}
			cur_node = stack[0]

			for cur_node.right != nil && !cur_node.right.visited {
				cur_node = cur_node.right
				cur_node.visited = true
				stack = append(stack, cur_node)
				path = append(path, cur_node.val)
			}

			if len(path) > 1 {
				paths = append(paths, path)
				path = []int{1}
			}
		}
		stack = stack[1:]
	}
	return paths
}

func main() {
	leaf_d := &Node{val: 5, left: nil, right: nil, visited: false}
	leaf_c := &Node{val: 4, left: nil, right: nil, visited: false}
	leaf_b := &Node{val: 3, left: leaf_c, right: leaf_d, visited: false}
	leaf_a := &Node{val: 2, left: nil, right: nil, visited: false}
	root := &Node{val: 1, left: leaf_a, right: leaf_b, visited: false}

	fmt.Println(get_paths(root))
}
