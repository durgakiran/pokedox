package cache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	CratedAt time.Time
	Val      []byte
}

type Cache struct {
	mu       sync.RWMutex
	Entries  map[string]CacheEntry
	Duration time.Duration
}

func deleteEntry(cache *Cache, key string) {
	// fmt.Printf("Deleting entry from cache with key: %s", key)
	cache.mu.Lock()
	delete(cache.Entries, key)
	cache.mu.Unlock()
}

func (c *Cache) ReapLoop() {
	for {
		time.Sleep(30 * time.Second)
		for key, entry := range c.Entries {
			if time.Since(entry.CratedAt) > c.Duration {
				// fmt.Printf("Deleting entry from cache with key: %s", key)
				go deleteEntry(c, key)
			} else {
				// fmt.Printf("Entry still valid in cache with key: %s, duration: %s", key, c.Duration-time.Since(entry.CratedAt))
			}
		}
	}
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{Entries: make(map[string]CacheEntry), Duration: interval}
	go cache.ReapLoop()
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	// fmt.Printf("Adding new entry to cache with key: %s", key)
	c.mu.Lock()
	c.Entries[key] = CacheEntry{time.Now(), val}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	// fmt.Printf("Getting entry from cache with key: %s", key)
	c.mu.RLock()
	entry, ok := c.Entries[key]
	c.mu.RUnlock()
	if !ok {
		// fmt.Printf("Entry not found in cache with key: %s", key)
		return nil, false
	}
	return entry.Val, true
}
