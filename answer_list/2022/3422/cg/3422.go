/*
   Weekly Coding Problem

   Week 34, 2022

   You 2 integers n and m representing an n by m grid, determine the number of ways you can
   get from the top-left to the bottom-right of the matrix y going only right or down.

   Example:
   `n = 2, m = 2`

   This should return 2, since the only possible routes are:
   Right, down
   Down, right.

   Here's the signature:

   def num_ways(n, m):
     # Fill this in.

   print num_ways(2, 2)
   # 2

*/
package main

import "fmt"

type matrix [][]int

func main() {
	fmt.Println(num_ways(2, 2))
}

func num_ways(n, m int) int {

	mtx := make(matrix, n)
	for i := range mtx {
		mtx[i] = make([]int, m)
	}

	count := 0
	recurse_mtx(&mtx, 0, 0, &count)

	return count
}

func recurse_mtx(m *matrix, i, j int, count *int) {

	if i+1 == len(*m) &&
		j+1 == len((*m)[i]) {
		return
	}

	iN := i + 1
	jN := j + 1

	if iN < len(*m) {
		recurse_mtx(m, iN, j, count)
	}

	if jN < len((*m)[i]) {
		recurse_mtx(m, i, jN, count)
	}

	if i != 0 || j != 0 {
		*count++
	}

	return
}
