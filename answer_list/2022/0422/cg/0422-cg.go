// ## Weekly Coding Problem
//
// ### Week 04, 2022 - (Hard)
//
// This problem was asked by **Nvidia**.
//
// Find the maximum of two numbers without using any if-else statements, branching, or direct comparisons.
package main

import (
	"fmt"
)

func main() {
	fmt.Println(MaxOfTwo(20, 3))
}

func MaxOfTwo(x, y int) int {

	return x & y
}
