package controllers

// Storage Interface Supports Get and Put
// of a single value
type Storage interface {
	Get() string
	Put(string)
}

// MemStorage implements Storage
type MemStorage struct {
	value string
}

// Get our in-memory value
func (m *MemStorage) Get() string {
	return m.value
}

// Put our in-memory value
func (m *MemStorage) Put(s string) {
	m.value = s
}
