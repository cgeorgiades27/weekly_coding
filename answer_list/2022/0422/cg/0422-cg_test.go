package main

import (
	"testing"
)

var tests = []struct {
	x, y, expected int32
}{
	{2, 3, 3},
	{3, 2, 3},
	{200, 199, 200},
	{1999, 1998, 1999},
}

func TestMaxOfTwo(t *testing.T) {
	for _, test := range tests {
		res := MaxOfTwo(test.x, test.y)
		if res != test.expected {
			t.Errorf("expected: %d, got: %d", test.expected, res)
		}
	}
}
