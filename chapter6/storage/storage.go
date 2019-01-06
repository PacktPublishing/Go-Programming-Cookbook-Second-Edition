package storage

import "context"

// Item represents an item at
// a shop
type Item struct {
	Name  string
	Price int64
}

// Storage is our storage interface
// We'll implement it with Mongo
// storage
type Storage interface {
	GetByName(context.Context, string) (*Item, error)
	Put(context.Context, *Item) error
}
