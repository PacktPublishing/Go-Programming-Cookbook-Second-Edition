package atomic

import (
	"sync"
	"sync/atomic"
)

// Ordinal holds a global a value
// and can only be initialized once
type Ordinal struct {
	ordinal uint64
	once    *sync.Once
}

// NewOrdinal returns ordinal with once
// setup
func NewOrdinal() *Ordinal {
	return &Ordinal{once: &sync.Once{}}
}

// Init sets the ordinal value
// can only be done once
func (o *Ordinal) Init(val uint64) {
	o.once.Do(func() {
		atomic.StoreUint64(&o.ordinal, val)
	})
}

// GetOrdinal will return the current
// ordinal
func (o *Ordinal) GetOrdinal() uint64 {
	return atomic.LoadUint64(&o.ordinal)
}

// Increment will increment the current
// ordinal
func (o *Ordinal) Increment() {
	atomic.AddUint64(&o.ordinal, 1)
}
