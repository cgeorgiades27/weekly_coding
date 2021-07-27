package main

import "fmt"

type matrix [][]int

func main() {
	// Setup
	m0 := matrix{{0, 2, 1}, {3, 0, 5}, {1, 0, 3}, {1, 4, 1}}
	m1 := matrix{{1, 7, 8}, {3, 6, 9}, {1, 10, 2}, {7, 4, 1}}
	m2 := matrix{{1, 1, 0, 1}, {1, 1, 2, 5}, {2, 1, 6, 4}, {0, 1, 1, 3}}
	m3 := matrix{{1, 1, 0, 1}, {1, 1, 2, 6}, {2, 1, 6, 4}, {0, 1, 0, 3}}
	matrices := []matrix{m0, m1, m2, m3}
	chosenMatrixIndex := 1

	// Run
	max := 0
	fmt.Printf("Testing the following matrix: %d\n", matrices[chosenMatrixIndex])
	fmt.Printf("The max sum is: %d\n", maxMatrixSum(&matrices[chosenMatrixIndex], &max, 0, 0))
}

func maxMatrixSum(m *matrix, max *int, i int, j int) int {

	// basecase
	if i > len(*m)-1 ||
		j > len((*m)[i])-1 {
		return 0
	}

	// create incremented i, j
	iN := i + 1
	jN := j + 1

	// Recurse to find max of subtrees
	maxLeft := maxMatrixSum(m, max, i, jN)
	maxRight := maxMatrixSum(m, max, iN, j)

	// take the larger of the two
	maxOfAll := Max(maxLeft, maxRight) + (*m)[i][j]

	// reset max if needed
	if maxOfAll > *max {
		*max = maxOfAll
	}

	return maxOfAll
}

// Max util
func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
