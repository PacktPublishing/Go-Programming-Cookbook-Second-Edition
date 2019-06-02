package bench

import "sync/atomic"

// AtomicCounter implements an atmoic lock
// using the atomic package
type AtomicCounter struct {
	value int64
}

// Add increments the counter
func (c *AtomicCounter) Add(amount int64) {
	atomic.AddInt64(&c.value, amount)
}

// Read returns the current counter amount
func (c *AtomicCounter) Read() int64 {
	var result int64
	result = atomic.LoadInt64(&c.value)
	return result
}
