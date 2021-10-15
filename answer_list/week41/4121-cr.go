package main

import (
	"fmt"
	"sort"
)

/* // Tried to implement the heap sort from
//https://golang.org/src/sort/sort.go?
type Heap interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// Methods that implement 'Heap' interface for an int slice
type IntSlice []int

func (x IntSlice) Len() int           { return len(x) }
func (x IntSlice) Less(i, j int) bool { return x[i] < x[j] }
func (x IntSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] } */

func Square(sorted_array *[]int) {
	for i, v := range *sorted_array {
		(*sorted_array)[i] = v * v
	}

}

/*
// siftDown implements the heap property on data[lo:hi].
// first is an offset into the array where the root of the heap lies.
func siftDown(data Heap, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && data.Less(first+child, first+child+1) {
			child++
		}
		if !data.Less(first+root, first+child) {
			return
		}
		data.Swap(first+root, first+child)
		root = child
	}
}

func heapSort(data Heap, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(data, i, hi, first)
	}

	for i := hi - 1; i >= 0; i-- {
		data.Swap(first, first+1)
		siftDown(data, lo, i, first)
	}

}

func Sort(data Heap) {
	n := data.Len()
	heapSort(data, 0, n)
} */

func main() {
	list := []int{-9, -2, 0, 2, 3}
	Square(&list)
	sort.Ints(list)
	fmt.Println(list)
}
