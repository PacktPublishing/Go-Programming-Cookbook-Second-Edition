package atomic

import (
	"errors"
	"sync"
)

// SafeMap uses a mutex to allow
// getting and setting in a thread-safe way
type SafeMap struct {
	m  map[string]string
	mu *sync.RWMutex
}

// NewSafeMap creates a SafeMap
func NewSafeMap() SafeMap {
	return SafeMap{m: make(map[string]string), mu: &sync.RWMutex{}}
}

// Set uses a write lock and sets the value given
// a key
func (t *SafeMap) Set(key, value string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.m[key] = value
}

// Get uses a RW lock and gets the value if it exists,
// otherwise an error is returned
func (t *SafeMap) Get(key string) (string, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	if v, ok := t.m[key]; ok {
		return v, nil
	}

	return "", errors.New("key not found")
}
