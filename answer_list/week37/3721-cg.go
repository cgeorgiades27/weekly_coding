/*
The **Sieve of Eratosthenes** is an algorithm used to generate all prime numbers smaller than N. The method is to take increasingly larger prime numbers, and mark their multiples as composite.

For example, to find all primes less than 100, we would first mark `[4, 6, 8, ...]` (multiples of two), then `[6, 9, 12, ...]` (multiples of three), and so on. Once we have done this for all primes less than N, the unmarked numbers that remain will be prime.

Implement this algorithm.

Bonus: Create a generator that produces primes indefinitely (that is, without taking N as an input).
*/

package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	var n, counter int
	var inf, q bool
	fmt.Printf("Enter the max prime (>2) or -1 for infinite: ")
	fmt.Scanln(&n)
	if n == -1 {
		n = 1000
		inf = true
		q = true
	} else {
		fmt.Printf("Do you want to print the primes? (1 for yes, 0 for no): ")
		fmt.Scanln(&q)
	}
	bitslc := make([]bool, n)
	start := time.Now()
	Sieve(n, &bitslc, &counter, inf, q)
	fmt.Printf("Primes found: %d, Time elapsed: %fs\n", counter, time.Now().Sub(start).Seconds())
}

func Sieve(n int, bitslc *[]bool, counter *int, inf bool, q bool) {
	max := n
	sqrtmax := math.Ceil(math.Sqrt(float64(max)))
	*counter = 0

	// Set all bits
	for i := range *bitslc {
		(*bitslc)[i] = true
	}
	// 0,1 are not prime
	(*bitslc)[0], (*bitslc)[1] = false, false
	// Unset all powers of 2
	for i := 4; i < max; i += 2 {
		(*bitslc)[i] = false
	}
	// Outer loop will run for sqrt(max)
	for i := 3; i < int(sqrtmax); i += 2 {
		if (*bitslc)[i] {
			for j := i * i; j < max; j += i {
				(*bitslc)[j] = false
			}
		}
	}
	// Only used for printing
	for i := range *bitslc {
		if (*bitslc)[i] {
			(*counter)++
			if q {
				fmt.Println(i)
			}
		}
	}
	// Infinite Sieve
	if inf {
		for {
			// Record old and reset max
			oldLen := len(*bitslc)
			max = len(*bitslc) * 2
			sqrtmax = math.Ceil(math.Sqrt(float64(max)))

			// resize the bit array
			new := make([]bool, (len(*bitslc)+1)*2)
			copy(new, *bitslc)
			bitslc = &new

			for i := oldLen - 1; i < max; i++ {
				(*bitslc)[i] = true
			}
			for i := oldLen; i < max; i += 2 {
				(*bitslc)[i] = false
			}
			for i := oldLen; i < int(sqrtmax); i += 2 {
				if (*bitslc)[i] {
					for j := i * i; j < max; j += 1 {
						(*bitslc)[j] = false
					}
				}
			}
			for i := oldLen; i < max; i++ {
				if (*bitslc)[i] {
					fmt.Println(i)
				}
			}
		}
	}
}
