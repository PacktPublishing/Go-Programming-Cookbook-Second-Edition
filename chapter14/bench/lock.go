package bench

import "sync"

// Counter uses a sync.RWMutex to safely
// modify a value
type Counter struct {
	value int64
	mu    *sync.RWMutex
}

// Add increments the counter
func (c *Counter) Add(amount int64) {
	c.mu.Lock()
	c.value += amount
	c.mu.Unlock()
}

// Read returns the current counter amount
func (c *Counter) Read() int64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}
