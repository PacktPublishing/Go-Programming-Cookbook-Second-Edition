package main

import (
	"sync"
)

var items []string
var mu *sync.RWMutex

func init() {
	mu = &sync.RWMutex{}
}

// AddItem adds an item to our list
// in a thread-safe way
func AddItem(item string) {
	mu.Lock()
	items = append(items, item)
	mu.Unlock()
}

// ReadItems returns our list of items
// in a thread-safe way
func ReadItems() []string {
	mu.RLock()
	defer mu.RUnlock()
	return items
}
