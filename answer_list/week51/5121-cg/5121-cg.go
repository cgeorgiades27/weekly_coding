// Weekly Coding Problem
// Week 51, 2021 - (Medium)
// This problem was asked by MongoDB.

// Given a list of elements ```lst```, find the majority element ```m```, which appears more than half the time. ```half = floor(len(lst) / 2))```

// You can assume that such element exists.

// For example, given ```[1, 2, 1, 1, 3, 4, 1]```, return ```1```.
package main

import (
	"fmt"
)

func main() {
	list := []int{4, 1, 3, 2, 3, 4, 4, 4, 4}
	fmt.Println(GetMajority(list))
}

func GetMajority(list []int) int {
	countSlc := make([]int, 100)
	for _, elem := range list {
		countSlc[elem]++
	}
	res := 0
	half := len(list) / 2
	for i, elem := range countSlc {
		if elem > half {
			res = i
		}
	}
	return res
}
