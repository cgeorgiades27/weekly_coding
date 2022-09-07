/*

  Weekly Coding Problem
  
  Week 36, 2022
  
  This problem was recently asked by **Twitter**:
  
  You are given the root of a binary tree. Invert the binary tree in place. That is, all left children should become right children, and all right children should become left children.
  
  Example:
  
      a
     / \
    b   c
   / \  /
  d   e f
  
  The inverted version of this tree is as follows:
  
    a
   / \
   c  b
   \  / \
    f e  d
  
  Here is the function signature:
  
  class Node:
    def __init__(self, value):
      self.left = None
      self.right = None
      self.value = value
    def preorder(self):
      print self.value,
      if self.left: self.left.preorder()
      if self.right: self.right.preorder()
  
  def invert(node):
    # Fill this in.
  
  root = Node('a')
  root.left = Node('b')
  root.right = Node('c')
  root.left.left = Node('d')
  root.left.right = Node('e')
  root.right.left = Node('f')
  
  root.preorder()
  # a b d e c f
  print "\n"
  invert(root)
  root.preorder()
  # a c f b e d
*/
package main

import "fmt"

type Node struct {
	left  *Node
	right *Node
	value byte
}

func main() {
  
	root := Node{value: 'a'}
	root.left = &Node{value: 'b'}
	root.right = &Node{value: 'c'}
	root.left.left = &Node{value: 'd'}
	root.left.right = &Node{value: 'e'}
	root.right.left = &Node{value: 'f'}

	fmt.Println("before:")
	preorder(&root)
	invert(&root)
	fmt.Println("after:")
	preorder(&root)
}

// preorder prints a binary tree in preorder
func preorder(root *Node) {

	// basecase
	if root == nil {
		return
	}

	fmt.Println(string(root.value))
	preorder(root.left)
	preorder(root.right)
}

// invert inverts a binary tree
func invert(root *Node) {

	// basecase
	if root == nil ||
		(root.left == nil && root.right == nil) {
		return
	}

	temp := root.left
	root.left = root.right
	root.right = temp

	invert(root.left)
	invert(root.right)
}
