/*
Weekly Coding Problem

## Week 41, 2022

This problem was recently asked by **Apple**:

You are given a binary tree representation of an arithmetic expression. In this tree, each leaf is an integer value,, and a non-leaf node is one of the four operations: `'+', '-', '*', or '/'`.

Write a function that takes this tree and evaluates the expression.

Example:
```
    *
   / \
  +    +
 / \  / \
3  2  4  5
```

This is a representation of the expression `(3 + 2) * (4 + 5)`, and should return `45`.

Here's a starting point:

``` python
class Node:
  def __init__(self, val, left=None, right=None):
    self.val = val
    self.left = left
    self.right = right

PLUS = "+"
MINUS = "-"
TIMES = "*"
DIVIDE = "/"

def evaluate(root):
  # Fill this in.

tree = Node(TIMES)
tree.left = Node(PLUS)
tree.left.left = Node(3)
tree.left.right = Node(2)
tree.right = Node(PLUS)
tree.right.left = Node(4)
tree.right.right = Node(5)
print evaluate(tree)
# 45
```
*/

package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	val   string
	left  *Node
	right *Node
}

const (
	PLUS   = "+"
	MINUS  = "-"
	TIMES  = "*"
	DIVIDE = "/"
)

func main() {
	tree := Node{val: TIMES}
	tree.left = &Node{val: PLUS}
	tree.left.left = &Node{val: "3"}
	tree.left.right = &Node{val: "2"}
	tree.right = &Node{val: PLUS}
	tree.right.left = &Node{val: "4"}
	tree.right.right = &Node{val: "5"}
	total := evaluate(&tree)
	fmt.Println(total)
}

func evaluate(root *Node) int {

	if root != nil {
		if v, err := strconv.Atoi(root.val); err == nil {
			return v
		}
	} else {
		return 0
	}

	l := evaluate(root.left)
	r := evaluate(root.right)
	return compute(l, r, root.val)
}

func compute(l, r int, op string) int {
	switch op {
	case PLUS:
		return l + r
	case MINUS:
		return l - r
	case TIMES:
		return l * r
	case DIVIDE:
		return l / r
	default:
		return 0
	}
}
