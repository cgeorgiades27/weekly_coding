/*
Week 44, 2021 - (Easy)

This problem was asked by **Amazon**.

Given a `N` by `M` matrix of numbers, print out the matrix in a clockwise spiral.

For example, given the following matrix:

[[1,  2,  3,  4,  5],
 [6,  7,  8,  9,  10],
 [11, 12, 13, 14, 15],
 [16, 17, 18, 19, 20]]

You should print out the following:

1
2
3
4
5
10
15
20
19
18
17
16
11
6
7
8
9
14
13
12
*/

package main

import "fmt"

func main() {
	m := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
	}
	for len(m) > 0 {

		var front, back int

		for len(m[0]) > 0 {
			front, m[0] = m[0][0], m[0][1:]
			fmt.Println(front, m)
		}
		for i := 0; i < len(m); i++ {
			back, m[i] = m[i][len(m[i])-1], m[i][:len(m[i])]
			fmt.Println(back, m)
		}
		for i := len(m[len(m)-1]); i >= 0; i-- {
			back, m[len(m)-1] = m[len(m)-1][i], m[len(m)-1][:len(m)]
			fmt.Println(back, m)
		}
		for i := len(m) - 1; i >= 0; i-- {
			back, m[0] = m[0][i], m[0][:i]
		}
	}
}

func RotateL90(m [][]int) [][]int {
	max := len(m)
	if len(m[0]) > max {
		max = len(m[0])
	}
	rotated := make([][]int, max)
	for i := range rotated {
		rotated[i] = make([]int, max)
	}

	for i := range m {
		for j := range m[i] {
			rotated[i][j] = m[len(m)-1-j][i]
		}
	}
	return rotated
}
