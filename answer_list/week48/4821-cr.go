package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Compares value at x[i] with previous value
// If they dont match writes count and value to encode
func encode(decoded_s string) string {
	var encoded strings.Builder
	length := len(decoded_s)
	i := 1
	total := 1

	for ; i < length; i++ {
		if decoded_s[i] != decoded_s[i-1] {
			encoded.WriteString(strconv.Itoa(total))
			encoded.WriteString(string(decoded_s[i-1]))
			total = 0
		}
		total++

	}
	encoded.WriteString(strconv.Itoa(total))
	encoded.WriteString(string(decoded_s[i-1]))

	return encoded.String()
}

// Gets int (x) from each even index until end
// writes value[i+1] to the decode x number of times
func decode(encoded_s string) string {
	var decoded strings.Builder
	i := 0
	length := len(encoded_s)
	for ; i < length; i += 2 {
		multiple := int(encoded_s[i] - '0') //int from ascii byte
		for t := 0; t < multiple; t++ {
			decoded.WriteString(string(encoded_s[i+1]))
		}

	}

	return decoded.String()

}

func main() {
	test_string := "AAAABBBCCDAA"
	encoded_string := encode(test_string)
	decoded_string := decode(encoded_string)
	fmt.Printf("Encoded: %s, Decoded: %s", encoded_string, decoded_string)

}
