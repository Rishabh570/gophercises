package models

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	Stories    []Item
	Expiration time.Time
	CacheMutex sync.Mutex
}

func (c *Cache) GetStoriesFromCache() []Item {
	fmt.Println("Trying to fetch from cache:", len(c.Stories))
	if time.Since(c.Expiration) < 0 {
		return c.Stories
	}
	return nil
}

func (c *Cache) SetStoriesInCache(stories []Item, expiration time.Time) {
	fmt.Println("setting cache, stories: ", len(stories), "expiration:", expiration)
	c.Stories = stories
	c.Expiration = expiration
}
