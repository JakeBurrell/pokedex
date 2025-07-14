package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheData map[string]cacheEntry
	mu        *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(intival time.Duration) Cache {
	newCache := Cache{
		cacheData: make(map[string]cacheEntry),
		mu:        &sync.Mutex{},
	}
	go func() {
		for {
			time.Sleep(intival)
			newCache.reapLoop(intival)
		}

	}()
	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	newCacheEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheData[key] = newCacheEntry
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cacheEntry, exists := c.cacheData[key]
	return cacheEntry.val, exists
}

func (c *Cache) reapLoop(intival time.Duration) {
	expiryTime := time.Now().Add(-intival)
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, value := range c.cacheData {
		if value.createdAt.Before(expiryTime) {
			delete(c.cacheData, key)
		}
	}
}
