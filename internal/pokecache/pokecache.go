package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	sync.RWMutex
	Cache map[string]CacheEntry
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache() Cache {
	return Cache{
		Cache: make(map[string]CacheEntry),
	}
}

func (c *Cache) Add(key string, val []byte) error {
	c.Lock()
	c.Cache[key] = CacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
	c.Unlock()

	fmt.Printf("%v ", key)
	fmt.Println()

	return nil
}

func (c *Cache) Get(key string) ([]byte, bool) {
	cacheEntry, ok := c.Cache[key]
	fmt.Printf("Cache entry (%s) exists? %v\n", key, ok)
	if !ok {
		return []byte{}, false
	}
	return cacheEntry.val, true
}

func (c *Cache) reapLoop() {}
