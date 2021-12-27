package main

import (
	"math/rand"
	"testing"
	"time"
)

const (
	seed      = 50
	listSize  = 800
	maxNumber = 1000000
	expected  = 123456
)

func TestFindUnique(t *testing.T) {
	list := MockDeliveryList(listSize, maxNumber, expected, seed)
	result := FindUnique(list, maxNumber)
	if result != expected {
		t.Errorf("Result should be: %d, got: %d", expected, result)
	}
}

func BenchmarkFindUnique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := MockDeliveryList(listSize, maxNumber, expected, time.Now().UnixNano())
		FindUnique(list, maxNumber)
	}
}

func FindUnique(delivery_id_confirmations []int, max int) int {
	countSlc := make([]int, max)
	for _, conf := range delivery_id_confirmations {
		countSlc[conf]++
	}
	for i, count := range countSlc {
		if count == 1 {
			return i
		}
	}
	return 0
}

// MockeDelivertList creates a fake delivery list and inserts
func MockDeliveryList(size, max, expected int, seed int64) []int {
	rand.Seed(seed)
	var randomInts []int
	for i := 0; i < size; i++ {
		randNum := rand.Intn(max)
		randomInts = append(randomInts, randNum)
		randomInts = append(randomInts, randNum)
	}
	// add final untracked package and shuffle
	randomInts = append(randomInts, expected)
	rand.Shuffle(len(randomInts), func(i, j int) { randomInts[i], randomInts[j] = randomInts[j], randomInts[i] })
	return randomInts
}
