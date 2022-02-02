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
	var tests = []struct {
		x, y, expected int32
	}{
		{2, 3, 3},
		{3, 2, 3},
		{200, 199, 200},
		{1999, 1998, 1999},
	}
	fmt.Println(tests[1].x)
	fmt.Println(MaxOfTwo(20, 3))
}

// MaxOfTwo -  ***** I did NOT come up with this own my own. I gave in and sought help from the internet (╥_╥) ******
func MaxOfTwo(x, y int32) int32 {
	z := x - y         /* take the difference */
	k := (z >> 31) & 1 /* check the most significant bit, if 1, z is negative */
	return x - k*z     /* bigger number minus 0 or smaller plus difference bigger and smaller */
}
