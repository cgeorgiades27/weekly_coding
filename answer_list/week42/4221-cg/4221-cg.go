// Given a 32-bit positive integer N,
// determine whether it is a power of four in faster than O(log N) time.
package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	var uint32Arr, randUint []uint32
	for i := 0; i < 32/2; i++ {
		uint32Arr = append(uint32Arr, 1<<(2*i))
	}
	for i := 0; i < 16; i++ {
		randUint = append(randUint, rand.Uint32())
	}
	for _, elem := range uint32Arr {
		fmt.Printf("%12d:%7t\n", elem, isPowerOfFour(elem))
	}
	for _, elem := range randUint {
		fmt.Printf("%12d:%7t\n", elem, isPowerOfFour(elem))
	}
}

func isPowerOfFour(num uint32) bool {
	res := math.Log(float64(num)) / math.Log(4)
	return (res-math.Floor(res) == 0)
}
