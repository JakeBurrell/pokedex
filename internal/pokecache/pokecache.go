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
	return Cache{}
}

func (c *Cache) Add(key string, val []byte) {

}

func (c Cache) Get(key string) ([]byte, bool) {
	return []byte{}, false
}

func (c *Cache) reapLoop() {

}
