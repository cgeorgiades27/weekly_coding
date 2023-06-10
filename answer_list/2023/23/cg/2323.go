package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func randomNumberFixedLength(k int) string {
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	chars := []rune("0123456789")
	firstChars := []rune("123456789")

	var b strings.Builder

	b.WriteRune(firstChars[randGen.Intn(len(firstChars))])
	for i := 1; i < k; i++ {
		b.WriteRune(chars[randGen.Intn(len(chars))])
	}

	return b.String()
}

func main() {
	fmt.Printf("enter a number: ")
	var length int
	fmt.Scanf("%d", &length)
	fmt.Printf("you entered: %d\n", length)
	if length < 1 {
		return
	}
	res := randomNumberFixedLength(length)
	fmt.Println("result is", res)
}
