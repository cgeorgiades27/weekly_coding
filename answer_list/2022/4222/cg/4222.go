/*
  Weekly Coding Problem

  Week 42, 2022

  This problem was recently asked by **Google**:

  A look-and-say sequence is defined as the integer sequence beginning with a single digit in
  which the next term is obtained by describing the previous term. An example is easier to understand:

  Each consecutive value describes the prior value.

  1      #
  11     # one 1's
  21     # two 1's
  1211   # one 2, and one 1.
  111221 # #one 1, one 2, and two 1's.

  Your task is, return the `nth` term of this sequence.
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("enter an upper limit (n) for look-and-say: ")
	var n int
	fmt.Scanln(&n)
	fmt.Println("final result is:", getLookAndSay(n))
}

func getLookAndSay(n int) string {
	count := 1
	s := lookAndSay("1") // basecase
	count++
	for count < n {
		s = lookAndSay(s)
		count++
	}
	return s
}

func lookAndSay(s string) string {

	if len(s) < 1 {
		fmt.Println(s)
		return "1"
	}

	result := ""
	count := 0
	curr := 0
	index := 0

	for index < len(s) {
		curr = index
		count = 0
		for index < len(s) && s[curr] == s[index] {
			count++
			index++
		}
		quant := strconv.Itoa(count)
		result += quant
		result += string(s[curr])
	}

	fmt.Println(s)
	return result
}
