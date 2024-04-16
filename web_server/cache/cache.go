package cache

import (
	"fmt"
	"sync"
	"time"
)

// CacheItem struct to store individual cache items
type CacheItem struct {
	Value      interface{}
	ExpiryTime time.Time
}

// Cache struct to hold the cache data
type Cache struct {
	items map[string]*CacheItem
	mutex sync.Mutex
}

// NewCache creates a new cache object
func NewCache() *Cache {
	return &Cache{
		items: make(map[string]*CacheItem),
	}
}

// Set adds an item to the cache with an expiry duration
func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.items[key] = &CacheItem{
		Value:      value,
		ExpiryTime: time.Now().Add(duration),
	}
}

// Get retrieves an item from the cache
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	item, found := c.items[key]
	if !found || time.Now().After(item.ExpiryTime) {
		if found {
			delete(c.items, key) // Clean up expired item
		}
		return nil, false
	}
	return item.Value, true
}

func CacheData(cache *Cache, key string, value interface{}, duration time.Duration) {
	fmt.Println("Caching data: " + key)
	cache.Set(key, value, duration)
}

// Main function to demonstrate cache usage
func TestCache() {
	fmt.Println("Testing cache")
	cache := NewCache()
	cache.Set("hello", "world", 10*time.Second)

	if value, found := cache.Get("hello"); found {
		fmt.Println("Found in cache:", value)
	} else {
		fmt.Println("Not found in cache")
	}

	// Wait for the item to expire
	time.Sleep(11 * time.Second)
	if _, found := cache.Get("hello"); !found {
		fmt.Println("Item has expired and is no longer available.")
	}
}
