package main

import (
	"testing"
)

var tests = []struct {
	input []int
	exp   int
}{
	{[]int{2, 5, 4, 8, 6, 3, 1, 4, 2, 3, 6, 5, 1}, 8},
	{[]int{4, 1, 3, 2, 3, 6, 4, 2, 1}, 6},
}

func TestGetDID(t *testing.T) {
	for _, test := range tests {
		result := getDID(test.input)
		if result != test.exp {
			t.Errorf("Result is %d, expected %d", result, test.exp)
		}
	}
}

func BenchmarkGetDID(b *testing.B) {
	test := tests[0]
	for i := 0; i < b.N; i++ {
		getDID(test.input)
	}
}
