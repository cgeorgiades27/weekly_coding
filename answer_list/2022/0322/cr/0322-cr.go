package main

import (
	"fmt"
	"time"
)

type CacheLRU struct {
	cap   int
	cache map[string]time.Time
	lruV  time.Time
	lruK  string
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
		fmt.Printf("Added %q to the cache at %v\n", v, time.Now())
	}
	for _, v := range input {
		x := lruCache.get(v)
		fmt.Printf("%q %v\n", v, x)
	}

}

func (c *CacheLRU) add(item string) {
	if len(c.cache) == 0 {
		c.cache = make(map[string]time.Time)
		c.lruV = time.Now()
	}
	if len(c.cache) < c.cap {
		c.cache[item] = time.Now()
		if c.cache[item].UnixMicro() <= c.lruV.UnixMicro() {
			c.lruV = c.cache[item]
			c.lruK = item
		}
	} else {
		delete(c.cache, c.lruK)
		c.cache[item] = time.Now()
		c.lruV = time.Now()
		c.lruK = item
	}
}

func (c *CacheLRU) get(key string) string {
	elem, ok := c.cache[key]
	if ok {
		return elem.String()
	} else {
		return "Null"
	}
}
