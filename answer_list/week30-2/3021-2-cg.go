/*
### Week 30, question 2, 2021

This problem was asked by Stripe.

Given an array of integers, find the first missing positive integer in linear time and constant space. In other words, find the lowest positive integer that does not exist in the array. The array can contain duplicates and negative numbers as well.

For example, the input `[3, 4, -1, 1]` should give `2`. The input `[1, 2, 0]` should give `3`.

You can modify the input array in-place.
*/

package main

import "fmt"

func main() {
	arr1 := []int{3, 4, -1, 1}
	arr2 := []int{1, 2, 0}
	arrs := [][]int{arr1, arr2}
	fmt.Println("Lowest is: ", lowestPositive(arrs[1]))
}

// NOT constant space
func lowestPositive(arr []int) int {
	var lowest, max int = 0, 0
	var coll = make([]int, 100, 100)

	// holding container
	for i := range arr {
		if arr[i] > 0 {
			coll[arr[i]] += 1
		}
	}
	// get the max
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		}
	}
	// find the missing val > 0
	for i := 0; i < max; i++ {
		if coll[i] == 0 {
			lowest = i
		}
	}
	// lowest is next if none are missing
	if lowest == 0 {
		lowest = max + 1
	}

	return (lowest)
}
