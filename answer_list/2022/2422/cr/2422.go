package main

import (
	"fmt"
)

func main() {
	intArray := [5]int{1, 2, 3, 4, 5}
	arrSum := sum(intArray[:])
	answer := makeProdArr(arrSum, intArray[:])
	fmt.Println(answer)
}

func sum(arr []int) int {
	prod := 1
	for _, v := range arr {
		prod *= v
	}
	return prod
}

func makeProdArr(s int, arr []int) []int {
	prodArr := make([]int, 0)
	for _, v := range arr {
		div := s / v
		prodArr = append(prodArr, div)
	}
	return prodArr
}
