package main

import (
	"testing"
)

const (
	TestCaseEncode       = "AAAABBBCCDAA"
	ExpectedResultEncode = "4A3B2C1D2A"
	TestCaseDecode       = "4A3B2C1D2A"
	ExpectedResultDecode = "AAAABBBCCDAA"
)

func TestRunEncoding(t *testing.T) {
	res := encode(TestCaseEncode)
	if res != ExpectedResultEncode {
		t.Errorf("encode(%s) = %s; should be %s", TestCaseEncode, res, ExpectedResultEncode)
	}
}

func TestRunDecoding(t *testing.T) {
	res := decode(TestCaseDecode)
	if res != ExpectedResultDecode {
		t.Errorf("decode(%s) = %s; should be %s", TestCaseDecode, res, ExpectedResultDecode)
	}
}

func BenchmarkEncoding(b *testing.B) {
	for i := 0; i < b.N; i++ {
		encode(TestCaseEncode)
	}
}

func BenchmarkDecoding(b *testing.B) {
	for i := 0; i < b.N; i++ {
		encode(TestCaseDecode)
	}
}
