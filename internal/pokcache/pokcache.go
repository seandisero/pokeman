package pokcache

import (
	"time"
)

func NewCache(interval time.Duration) *Cache {
	newCache := Cache{
		Entries: make(map[string]CacheEntry),
	}
	go newCache.reapLoop(interval)
	return &newCache
}

// Add -
func (cache *Cache) Add(key string, val []byte) {
	newEntry := CacheEntry{
		CreatedAt: time.Now(),
		Val:       val,
	}

	cache.Lock()
	defer cache.Unlock()

	cache.Entries[key] = newEntry
}

// Get -
func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.Lock()
	defer cache.Unlock()

	entry, exists := cache.Entries[key]
	return entry.Val, exists
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval) // you used to need to call defer ticker.Stop(), but now it is handled by garbage collection

	for range ticker.C {
		for key, entry := range cache.Entries {
			if time.Since(entry.CreatedAt) > interval {
				cache.Lock()
				delete(cache.Entries, key)
				cache.Unlock()
			}
		}
	}
}
