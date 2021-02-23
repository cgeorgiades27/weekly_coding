/*
## Weekly Coding Problem

### CHRISmas 2021 - (Medium)

This problem was asked by **Mr. Claus**.

##Santa has been testing autonomous drones to help fulfill demand during labor shortages...and...well something mysterious has happened. So first thing...drones go to warehouses and pick up unwrapped gifts.

Each warehouse delivery is assigned a delivery ID, a positive integer. When one of Mr.Claus' drones takes off with a delivery, the delivery's ID is added to a list, ```delivery_id_confirmations```. When the drone comes back and lands, the ID is again added to the same list.

After breakfast this morning there was one less drone on the tarmac. One of the drones never made it back from a delivery. Mr.Claus suspects a secret agent from Bahumbug Inc. placed an order and stole one of the patented drones. To track them down, we need to find the delivery ID for the order they placed.

Given the list of IDs, which contains many duplicate integers and one unique integer, find the unique integer.

The IDs are not guaranteed to be sorted or sequential. Orders aren't always fulfilled in the order they were received, and some deliveries get cancelled before takeoff.

## Can you help save CHRISmas
*/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	delivery_id_confirmations := GenerateRandomList(10, 200000, true)
	rand.Seed(50)
	delivery_id_confirmations = append(delivery_id_confirmations, delivery_id_confirmations[rand.Intn(len(delivery_id_confirmations))])
	fmt.Println(delivery_id_confirmations)
	fmt.Println(findUnique(delivery_id_confirmations, 200000))
}

func findUnique(delivery_id_confirmations []int, max int) int {
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

func GenerateRandomList(size, max int, wDuplicates bool) []int {
	rand.Seed(50)
	var randomInts []int
	for i := 0; i < size; i++ {
		randNum := rand.Intn(max)
		count := rand.Intn(3) + 1
		if wDuplicates {
			for j := 0; j < count; j++ {
				randomInts = append(randomInts, randNum)
			}
		}
		randomInts = append(randomInts, randNum)
	}
	return randomInts
}
