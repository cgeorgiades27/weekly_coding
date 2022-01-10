// This problem was asked by Google.
// Given a sorted list of integers, square the elements and give the output in sorted order.
// For example, given `[-9, -2, 0, 2, 3]`, return `[0, 4, 4, 9, 81]`.
package main

import "fmt"

func main() {
	arr := []int{-9, -2, 0, 2, 3}
	SquareNSort(&arr)
	fmt.Println(arr)
}

// SquareNSort square the elem and then sorts it
func SquareNSort(arr *[]int) {
	bubbleSort(square(arr))
}

func square(arr *[]int) *[]int {
	for i, elem := range *arr {
		(*arr)[i] = elem * elem
	}
	return arr
}

func bubbleSort(arr *[]int) {
	for {
		swaps := 0
		for i := 0; i < len(*arr)-1; i++ {
			if (*arr)[i] > (*arr)[i+1] {
				tmp := (*arr)[i]
				(*arr)[i] = (*arr)[i+1]
				(*arr)[i+1] = tmp
				swaps++
			}
		}
		if swaps == 0 {
			break
		}
	}
}
