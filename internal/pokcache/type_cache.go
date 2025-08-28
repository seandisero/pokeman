package pokcache

import (
	"sync"
	"time"
)

type Cache struct {
	sync.Mutex
	Entries map[string]CacheEntry
}

type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}
