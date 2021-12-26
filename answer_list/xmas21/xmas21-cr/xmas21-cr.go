package main

import (
	"fmt"
)

func main() {
	idList := []int{2, 5, 4, 8, 6, 3, 1, 4, 2, 3, 6, 5, 1}
	missingID := getDID(idList)
	fmt.Printf("The missing drone has ID %d", missingID)
}

// Optimal answer from the original example.
// Uses bitwise XOR to eliminate dups.
func getDID(lst []int) int {
	var result int
	for _, val := range lst {
		result ^= val
	}

	return result
}
