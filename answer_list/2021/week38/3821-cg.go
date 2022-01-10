/*
	Implement division of two positive integers without using the
	division, multiplication, or modulus operators.
	Return the quotient as an integer, ignoring the remainder...
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// test 20 random combinations
	for i := 0; i < 20; i++ {
		dividend := rand.Intn(1000)
		divisor := rand.Intn(dividend)
		fmt.Printf("%-3d/%3d = %3d\n", dividend, divisor, divide(dividend, divisor))
	}
}

// divide without /,*,%
func divide(dividend int, divisor int) int {
	accum, quotient := 0, 0
	for i := 0; i < dividend; i++ {
		if accum+divisor > dividend {
			return quotient
		}
		accum += divisor
		quotient++
	}
	return quotient
}
