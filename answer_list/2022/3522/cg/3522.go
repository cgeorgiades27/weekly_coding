/*
Weekly Coding Problem

Week 35, 2022

This problem was recently asked by Twitter:

A palindrome is a sequence of characters that reads the same backwards and forwards. Given a string, `s`, find the longest palindromic substring in `s`.

Example:
Input: "banana"
Output: "anana"

Input: "million"
Output: "illi"

class Solution:
    def longestPalindrome(self, s):
      # Fill this in.

# Test program
s = "tracecars"
print(str(Solution().longestPalindrome(s)))
# racecar
*/

package main

import (
	"fmt"
)

func main() {
	s1 := "million"
	s2 := "tracecars"
	fmt.Println(longestPalindrome2(s1))
	fmt.Println(longestPalindrome2(s2))
}

func longestPalindrome2(s string) string {

  longest := ""

  for i := range(s) {
    
    left, right := i, i  
    for left >= 0 && right < len(s) && s[left] == s[right] {
      if right - left > len(longest) {
        longest = s[left:right+1]
      }
      right++
      left--
    }

    left, right = i, i+1
    for left >= 0 && right < len(s) && s[left] == s[right] {
      if right - left > len(longest) {
        longest = s[left:right+1]
      }
      right++
      left--
    }
    
  }
  return longest
}

// O(n^3)
func longestPalindrome(s string) string {

	longest := ""

	for i := range s {
		
    subString := s[i:]
		temp := ""
    
    for j := 0; j < len(subString); j++ {
      if isPalindrome(subString[:j]) {
        temp = subString[:j]
      }
    }
    
		if len(temp) > len(longest) {
			longest = temp
		}
    
	}
	return longest
}

func isPalindrome(s string) bool {
  for i := 0; i < len(s); i++ {
    if s[i] != s[len(s)-i-1] { 
      return false
    }
  }  
  return true
}