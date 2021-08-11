/*
Weekly Coding Problem

Week 32, 2021

This problem was asked by Amazon:

Given a binary tree, return all values given a certain height h.

Here's a starting point:

class Node():
  def __init__(self, value, left=None, right=None):
    self.value = value
    self.left = left
    self.right = right

def valuesAtHeight(root, height):
  # Fill this in.

#     1
#    / \
#   2   3
#  / \   \
# 4   5   7

a = Node(1)
a.left = Node(2)
a.right = Node(3)
a.left.left = Node(4)
a.left.right = Node(5)
a.right.right = Node(7)
print valuesAtHeight(a, 3)
# [4, 5, 7]
*/

package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func valuesAtHeight(root *Node, height int) []int {
	var values []int
	traverseTree(root, 1, height, &values)
	return values
}

func traverseTree(root *Node, count int, maxHeight int, arr *[]int) {
	// basecase
	if root == nil || count > maxHeight {
		return
	}
	// inorder traversal
	if count == maxHeight {
		*arr = append(*arr, root.value)
	}
	traverseTree(root.left, count+1, maxHeight, arr)
	traverseTree(root.right, count+1, maxHeight, arr)
}

func main() {
	a := &Node{value: 1, left: nil, right: nil}
	a.left = &Node{value: 2, left: nil, right: nil}
	a.right = &Node{value: 3, left: nil, right: nil}
	a.left.left = &Node{value: 4, left: nil, right: nil}
	a.left.right = &Node{value: 5, left: nil, right: nil}
	a.right.right = &Node{value: 7, left: nil, right: nil}
	fmt.Println(valuesAtHeight(a, 3))
}
