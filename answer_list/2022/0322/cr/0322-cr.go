package main

import "fmt"

type CacheLRU struct {
	cap   int
	cache map[string]int
	min   int
	lru   string
}

func main() {
	n := 5
	lruCache := new(CacheLRU)
	lruCache.cap = n
	input := [25]string{"playstation.us", "botnet.ru", "botnet.ru",
		"playstation.us", "playstation.us", "stonks.com",
		"playstation.us", "amscotcashadvance.com", "peloton.com",
		"cannedham.edu", "stonks.com", "amscotcashadvance.com",
		"playstation.us", "of.com", "of.com", "google.com",
		"cannedham.edu", "peloton.com", "amscotcashadvance.com", "stonks.com",
		"google.com", "of.com", "amscotcashadvance.com", "go.dev", "google.com"}

	for _, v := range input {
		lruCache.add(v)
		fmt.Printf("Added %q to the cache \n", v)
	}
	for _, v := range input {
		x, y := lruCache.get(v)
		fmt.Println(x, y)
	}

}

func (c *CacheLRU) add(item string) {
	if len(c.cache) == 0 {
		c.cache = make(map[string]int)
		c.min = 1
	}
	if len(c.cache) < c.cap {
		c.cache[item]++
		if c.cache[item] <= c.min {
			c.min = c.cache[item]
			c.lru = item
		}
	} else {
		delete(c.cache, c.lru)
		c.cache[item]++
		c.min = 1
		c.lru = item
	}
}

func (c *CacheLRU) get(key string) (int, bool) {
	elem, ok := c.cache[key]
	return elem, ok
}
