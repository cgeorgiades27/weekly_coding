package main

import (
	"testing"
)

func TestProdArr(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{120, 60, 40, 30, 24}},
		{[]int{3, 2, 1}, []int{2, 3, 6}},
	}
	for _, test := range tests {
		s := sum(test.input)
		res := makeProdArr(s, test.input)
		expected := testEq(res, test.want)
		if !expected {
			t.Errorf("expected: %d, got: %d", test.want, res)
		}
	}
}

func testEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
