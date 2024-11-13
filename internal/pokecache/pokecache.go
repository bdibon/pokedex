package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	store map[string]cacheEntry
	l     sync.RWMutex
}

func (c *Cache) Add(key string, val []byte) {
	c.l.Lock()
	defer c.l.Unlock()

	c.store[key] = cacheEntry{time.Now(), val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.l.RLock()
	defer c.l.RUnlock()

	ent, ok := c.store[key]
	return ent.val, ok
}

func (c *Cache) reaploop(interval time.Duration) {
	t := time.NewTicker(interval)
	for {
		select {
		case ct := <-t.C:
			c.l.Lock()
			for k, v := range c.store {
				lifetime := ct.Sub(v.createdAt)
				if lifetime > interval {
					delete(c.store, k)
				}
			}
			c.l.Unlock()
		}
	}
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		store: map[string]cacheEntry{},
	}
	go c.reaploop(interval)
	return &c
}
