package main

import (
	"testing"
)

var tests = []struct {
	list []int
	exp  int
}{
	{[]int{1, 2, 1, 1, 3, 4, 1}, 1},
	{[]int{4, 1, 3, 2, 3, 4, 4, 4, 4}, 4},
}

func TestGetMajority(t *testing.T) {
	for _, test := range tests {
		result := GetMajority(test.list)
		if result != test.exp {
			t.Errorf("Expected %d, got %d", test.exp, result)
		}
	}
}

func BenchmarkGetMajority(b *testing.B) {
	test := tests[0]
	for i := 0; i < b.N; i++ {
		GetMajority(test.list)
	}
}
