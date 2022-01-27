// ## Weekly Coding Problem

// ### Week 03, 2022 - (Hard)

// This problem was asked by **Google**.

// Implement an LRU (Least Recently Used) cache. It should be able to be initialized with a cache size `n`, and contain the following methods:

// `set(key, value)`: sets `key` to `value`. If there are already `n` items in the cache and we are adding a new item, then it should also remove the least recently used item.

// `get(key)`: gets the value at key. If no such key exists, return null.

// Each operation should run in **O(1)** time.

// Example Data:

// ```
// n = 5

// input = ["playstation.us", "botnet.ru", "botnet.ru", "playstation.us", "playstation.us", "stonks.com", "playstation.us", "amscotcashadvance.com", "peloton.com", "cannedham.edu", "stonks.com", "amscotcashadvance.com", "playstation.us", "of.com", "of.com", "google.com", "cannedham.edu", "peloton.com", "amscotcashadvance.com", "stonks.com", "google.com", "of.com", "amscotcashadvance.com", "go.dev", "google.com"]
// ```
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input []string = []string{"playstation.us", "botnet.ru", "botnet.ru", "playstation.us", "playstation.us", "stonks.com", "playstation.us", "amscotcashadvance.com", "peloton.com", "cannedham.edu", "stonks.com", "amscotcashadvance.com", "playstation.us", "of.com", "of.com", "google.com", "cannedham.edu", "peloton.com", "amscotcashadvance.com", "stonks.com", "google.com", "of.com", "amscotcashadvance.com", "go.dev", "google.com"}
	c := NewCache(5)
	for i, in := range input {
		c.Set(in, strconv.Itoa(i))
	}
	c.Print()
}

type Cache struct {
	icount int
	hcount int
	mcount int
	size   uint32
	store  map[string]cache_entry
	age    map[int]string
}

type cache_entry struct {
	value string
	age   int
}

func (csh cache_entry) String() string {
	return fmt.Sprintf("value: %s, age: %d", csh.value, csh.age)
}

func NewCache(cacheSize uint32) Cache {
	return Cache{
		size:  cacheSize,
		age:   make(map[int]string, cacheSize),
		store: make(map[string]cache_entry, cacheSize)}
}

// get - private getter that trackers hit/misses
func (c *Cache) get(key string) cache_entry {
	val, exists := c.store[key]
	if !exists {
		c.mcount++
		return cache_entry{}
	}
	c.hcount++
	return val
}

// Get - public getter
func (c *Cache) Get(key string) string {
	return c.get(key).value
}

// Set - main cache functionality
func (c *Cache) Set(key, value string) {
	entry := c.get(key)
	if entry.value == "" {
		c.roll(1)
		_, exists := c.age[0]
		if exists {
			delete(c.store, c.age[0])
		}
		c.store[key] = cache_entry{value: value, age: 0}
		c.icount++
	} else {
		c.roll(c.store[key].age + 1)
		entry.age = 0
		entry.value = value
		c.store[key] = entry
	}
}

// roll - updates the ages of the bits by rolling them
func (c *Cache) roll(n int) {
	for i := 0; i < n; i++ {
		for key, entry := range c.store {
			entry.age = (entry.age + 1) % int(c.size)
			c.age[entry.age] = key
			c.store[key] = entry
		}
	}
}

func (c *Cache) Print() {
	fmt.Println()
	fmt.Println("****CACHE SUMMARY****")
	fmt.Printf("%-10s%5d\n%-10s%5d\n%-10s%5d\n\n", "inserts:", c.icount, "hits:", c.hcount, "misses", c.mcount)
	fmt.Println("Current State:")
	fmt.Printf("%-25s%15s%10s\n", "KEY", "AGE", "CONTENTS")
	fmt.Printf("%-25s%15s%10s\n", "___", "___", "________")
	for key, c_entry := range c.store {
		fmt.Printf("%-25s%15d%10s\n", key, c_entry.age, c_entry.value)
	}
}
