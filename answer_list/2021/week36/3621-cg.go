/*
The edit distance between two strings refers to the minimum number of character insertions, deletions, and substitutions required to change one string to the other. For example, the edit distance between “kitten” and “sitting” is three: substitute the “k” for “s”, substitute the “e” for “i”, and append a “g”.
*/

package main

import (
	"fmt"
	"os"
)

type Matrix [][]int

func main() {
	if len(os.Args) < 3 {
		fmt.Print("Too few args. Provide two strings to compare.")
		os.Exit(1)
	}

	// Sub cost
	const LevConstant = 1
	var lev int

	m := string(os.Args[1])
	n := string(os.Args[2])

	// Create the Matrix
	// Set up the initial vals
	mtrx := make(Matrix, len(m)+1)
	for i := range mtrx {
		mtrx[i] = make([]int, len(n)+1)
	}
	for i := 0; i <= len(m); i++ {
		mtrx[i][0] = i
	}
	for j := 0; j <= len(n); j++ {
		mtrx[0][j] = j
	}

	for i := 1; i < len(mtrx); i++ {
		for j := 1; j < len(mtrx[i]); j++ {
			// if matched, sub cost = 0, else Lev Const
			if m[i-1] == n[j-1] {
				lev = 0
			} else {
				lev = LevConstant
			}
			mtrx[i][j] = MinOfThree(
				mtrx[i-1][j-1]+lev,
				mtrx[i-1][j]+1,
				mtrx[i][j-1]+1)
		}
	}
	fmt.Printf("Sub cost = %d (with Levenshtein - substitution cost = %d)\n", mtrx[len(m)][len(n)], LevConstant)
}

func MinOfThree(x int, y int, z int) int {
	min := x
	if y < min {
		min = y
	}
	if z < min {
		min = z
	}
	return min
}
