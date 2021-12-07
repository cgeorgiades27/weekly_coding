package main

import (
	"fmt"
	"strconv"
	"strings"
)

func encode(decoded_s string) string {
	var encoded strings.Builder
	l := len(decoded_s)
	i := 1
	total := 1

	for i < l {
		if decoded_s[i] != decoded_s[i-1] {
			encoded.WriteString(strconv.Itoa(total) + string(decoded_s[i-1]))
			total = 0
		}
		i++
		total++

	}
	encoded.WriteString(strconv.Itoa(total) + string(decoded_s[i-1]))

	return encoded.String()
}

func decode(encoded_s string) string {
	var decoded strings.Builder
	i := 0
	length := len(encoded_s)
	for ; i < length; i += 2 {
		multiple := int(encoded_s[i] - '0')
		for t := 0; t < multiple; t++ {
			decoded.WriteString(string(encoded_s[i+1]))
		}

	}

	return decoded.String()

}

func main() {
	s := "AAAABBBCCDAA"
	e := encode(s)
	d := decode(e)
	fmt.Printf("Encoded: %s, Decoded: %s", e, d)

}
