// Week 48, 2021 - (Easy)
//
// This problem was asked by **Amazon**.
// Run-length encoding is a fast and simple method of encoding strings. The basic idea is to represent repeated successive
// characters as a single count and character. For example, the string ```"AAAABBBCCDAA"``` would be encoded as ```4A3B2C1D2A```.
// Implement run-length encoding and decoding. You can assume the string to be encoded have no digits and consists solely
// of alphabetic characters. You can assume the string to be decoded is valid.
//
// Run the test(s): go test -v
// Run the benchmark: go test -bench=.
package main

import (
	"strconv"
	"testing"
)

const (
	TestCase       = "AAAABBBCCDAA"
	ExpectedResult = "4A3B2C1D2A"
)

func Encode(str string) string {
	var bs string
	for len(str) > 0 {
		count := 0
		ctx := str[0]
		for str[0] == ctx {
			count++
			if len(str) > 1 {
				str = str[1:]
			} else {
				str = ""
				break
			}
		}
		bs = bs + strconv.Itoa(count)
		bs = bs + string(ctx)
	}
	return bs
}

func TestRunEncoding(t *testing.T) {
	res := Encode(TestCase)
	if res != ExpectedResult {
		t.Errorf("encode(%s) = %s; should be %s", TestCase, res, ExpectedResult)
	}
}

func BenchmarkEncoding(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Encode(TestCase)
	}
}
