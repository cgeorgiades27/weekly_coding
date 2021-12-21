package main

import (
	"fmt"
)

func findMajority(intSlice []int) int {
	var m int
	cMap := make(map[int]int)
	half := len(intSlice) / 2.0
	for _, val := range intSlice {
		_, ok := cMap[val]
		if !ok {
			cMap[val] = 1
		} else {
			cMap[val] += 1
		}
		if cMap[val] > half {
			m = val
			return m
		}
	}
	return m
}

func main() {
	lst := []int{1, 2, 1, 1, 3, 4, 1}
	fmt.Println(findMajority(lst))
}
