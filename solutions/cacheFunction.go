package solutions

import (
	"fmt"
	"sync"
)

// Cache is a simple in-memory cache
type Cache struct {
	data    map[string]interface{}
	maxSize int
	mu      sync.Mutex
}

// NewCache creates a new cache with the specified maxSize
func NewCache(maxSize int) *Cache {
	return &Cache{
		data:    make(map[string]interface{}),
		maxSize: maxSize,
	}
}

// Set adds a key-value pair to the cache
func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Check if the cache is full
	if len(c.data) >= c.maxSize {
		c.evict()
	}

	c.data[key] = value
}

// Get retrieves a value from the cache based on the key
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	value, exists := c.data[key]
	return value, exists
}

// evict removes the least recently used item when the cache is full
func (c *Cache) evict() {
	// In a real-world scenario, you might implement a more sophisticated eviction policy.
	// Here, we simply remove the first item in the cache.
	for k := range c.data {
		delete(c.data, k)
		break
	}
}

func init() {
	// Create a new cache with a maxSize of 2
	cache := NewCache(2)

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
