/*
## Weekly Coding Problem

### Week 33, 2021

In order to win the prize for most cookies sold, my friend Alice and I are going to merge our Girl Scout Cookies orders and enter as one unit.
Each order is represented by an ```order_id``` (an integer).

We have our lists of orders sorted numerically already, in lists. Write a function to merge our lists of orders into one sorted list.

For example:

l1     = [3, 4, 6, 10, 11, 15]
l2 = [1, 5, 8, 12, 14, 19]

# Prints [1, 3, 4, 5, 6, 8, 10, 11, 12, 14, 15, 19]
print(merge_lists(l1, l2))

### Bonus
Write a function that takes as an input a list of sorted lists and outputs a single sorted list with all the items from each list.
*/

package main

import "fmt"

func main() {
	// my_list := []int{3, 4, 6, 10, 11, 15, 16, 17}
	// alices_list := []int{1, 5, 8}

	my_list := []int{3, 4, 6, 10, 11, 15}
	alices_list := []int{1, 5, 8, 12, 14, 19}

	fmt.Println(merge_lists(my_list, alices_list))
}

/**
** Merge sort approach **
**/
func merge_lists(l1 []int, l2 []int) []int {
	// temp pointer to hold appended value
	var q *int
	var final_list []int

	// pointers to the front of each list
	var p1 *int = &l1[0]
	var p2 *int = &l2[0]

	for (len(l1) > 0) && (len(l2) > 0) {
		// check first list elements
		if *p2 < *p1 {
			q = p2
			if len(l2) > 1 {
				p2 = &l2[1]
				l2 = l2[1:]
			} else {
				l2 = l2[:0]
			}
			final_list = append(final_list, *q)
		} else {
			q = p1
			if len(l1) > 1 {
				p1 = &l1[1]
				l1 = l1[1:]
			} else {
				l1 = l1[:len(l1)-1]
			}
			final_list = append(final_list, *q)
		}
	}
	// append what remains of the other sorted list
	if len(l1) > 0 {
		final_list = append(final_list, l1...)
	} else {
		final_list = append(final_list, l2...)
	}
	return final_list
}
