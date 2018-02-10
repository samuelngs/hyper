// Package dataloader source: https://github.com/nicksrandall/dataloader
//
// dataloader is an implimentation of facebook's dataloader in go.
// See https://github.com/facebook/dataloader for more information
package dataloader

import (
	"sync"
)

// The Cache interface. If a custom cache is provided, it must implement this interface.
type Cache interface {
	Get(interface{}) (Thunk, bool)
	Set(interface{}, Thunk)
	Delete(interface{}) bool
	Clear()
}

// InMemoryCache is an in memory implementation of Cache interface.
// this simple implementation is well suited for
// a "per-request" dataloader (i.e. one that only lives
// for the life of an http request) but it not well suited
// for long lived cached items.
type InMemoryCache struct {
	items map[interface{}]Thunk
	mu    sync.RWMutex
}

// newCache constructs a new InMemoryCache
func newCache() *InMemoryCache {
	items := make(map[interface{}]Thunk)
	return &InMemoryCache{
		items: items,
	}
}

// Set sets the `value` at `key` in the cache
func (c *InMemoryCache) Set(key interface{}, value Thunk) {
	c.mu.Lock()
	c.items[key] = value
	c.mu.Unlock()
}

// Get gets the value at `key` if it exsits, returns value (or nil) and bool
// indicating of value was found
func (c *InMemoryCache) Get(key interface{}) (Thunk, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.items[key]
	if !found {
		return nil, false
	}

	return item, true
}

// Delete deletes item at `key` from cache
func (c *InMemoryCache) Delete(key interface{}) bool {
	if _, found := c.Get(key); found {
		c.mu.Lock()
		defer c.mu.Unlock()
		delete(c.items, key)
		return true
	}
	return false
}

// Clear clears the entire cache
func (c *InMemoryCache) Clear() {
	c.mu.Lock()
	c.items = map[interface{}]Thunk{}
	c.mu.Unlock()
}
