/*
## Weekly Coding Problem

### Week 32, 2021

This problem was asked by Amazon:

Given a binary tree, return all values given a certain height h.

Here's a starting point:

``` python
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
```
*/

package main

import "math"

type Node struct {
	value byte
	left  *Node
	right *Node
}

func valuesAtHeight(root *Node, height byte) []byte {
	values := make([]byte, int(math.Pow(2, float64(height))))
	return values
}

func main() {
	a := &Node{value: 1, left: nil, right: nil}
	a.left = &Node{value: 2, left: nil, right: nil}
	a.right = &Node{value: 3, left: nil, right: nil}
	a.left.left = &Node{value: 4, left: nil, right: nil}
	a.left.left = &Node{value: 5, left: nil, right: nil}
	a.right.right = &Node{value: 7, left: nil, right: nil}
	println(valuesAtHeight(a, 3))
}
