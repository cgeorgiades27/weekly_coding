package main

import (
	"testing"
)

const (
	ExpectedResultMajority = 1
)

func TestRunMajority(t *testing.T) {
	lst := []int{1, 2, 1, 1, 3, 4, 1}
	res := findMajority(lst)
	if res != ExpectedResultMajority {
		t.Errorf("Error: findMajority(%v) returned %d, Expected %d", lst, res, ExpectedResultMajority)
	}
}

func BenchmarkMajority(b *testing.B) {
	lst := []int{1, 2, 1, 1, 3, 4, 1}
	for i := 0; i < b.N; i++ {
		findMajority(lst)
	}
}
