package solutions

import (
	"fmt"
	"sync"
	"time"
)

// GoCacheItem represents an item in the cache with a creation time
type GoCacheItem struct {
	Value      interface{}
	CreateTime time.Time
}

// GoCache is a simple in-memory cache with TTL
type GoCache struct {
	data    map[string]GoCacheItem
	maxSize int
	ttl     time.Duration
	mu      sync.Mutex
}

// NewGoCache creates a new cache with the specified maxSize and TTL
func NewGoCache(maxSize int, ttl time.Duration) *GoCache {
	return &GoCache{
		data:    make(map[string]GoCacheItem),
		maxSize: maxSize,
		ttl:     ttl,
	}
}

// Set adds a key-value pair to the cache with TTL
func (c *GoCache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Check if the cache is full
	if len(c.data) >= c.maxSize {
		c.evict()
	}

	// Add the item to the cache with current time
	c.data[key] = GoCacheItem{
		Value:      value,
		CreateTime: time.Now(),
	}
}

// Get retrieves a value from the cache based on the key
func (c *GoCache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, exists := c.data[key]
	if exists && time.Since(item.CreateTime) <= c.ttl {
		return item.Value, true
	}

	// Item has expired or doesn't exist
	delete(c.data, key)
	return nil, false
}

// evict removes the least recently used item when the cache is full
func (c *GoCache) evict() {
	// In a real-world scenario, you might implement a more sophisticated eviction policy.
	// Here, we simply remove the first item in the cache.
	for k := range c.data {
		delete(c.data, k)
		break
	}
}

func init() {
	// Create a new cache with a maxSize of 2 and TTL of 5 seconds
	cache := NewGoCache(2, 5*time.Second)

	// Set key-value pairs in the cache
	cache.Set("key1", "value1")
	cache.Set("key2", "value2")

	// Retrieve values from the cache
	val1, found1 := cache.Get("key1")
	val2, found2 := cache.Get("key2")
	val3, found3 := cache.Get("key3")

	fmt.Printf("Found key1: %v, Value: %v\n", found1, val1)
	fmt.Printf("Found key2: %v, Value: %v\n", found2, val2)
	fmt.Printf("Found key3: %v, Value: %v\n", found3, val3)
}
